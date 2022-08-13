package leveldb

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/zRich/zFusion/storage/spi"
)

const (
	pathPattern = "%s-%s"

	invalidTagName                  = `"%s" is an invalid tag name since it contains one or more ':' characters`
	invalidTagValue                 = `"%s" is an invalid tag value since it contains one or more ':' characters`
	tagMapKey                       = "TagMap"
	storeConfigKey                  = "StoreConfig"
	expressionTagNameOnlyLength     = 1
	expressionTagNameAndValueLength = 2
	invalidQueryExpressionFormat    = `"%s" is not in a valid expression format. ` +
		"it must be in the following format: TagName:TagValue"
)

type Provider struct {
	dbPath string
	dbs    map[string]*store
	lock   sync.RWMutex
}

func NewProvider(dbpath string) *Provider {
	return &Provider{dbs: make(map[string]*store), dbPath: dbpath}
}

type closer func(storename string)

type tagMapping map[string]map[string]struct{}

type dbEntry struct {
	Value []byte
	Tags  []spi.Tag
}

type store struct {
	db    *leveldb.DB
	name  string
	close closer
	lock  sync.RWMutex
}

func (p *Provider) OpenStore(name string) (spi.Store, error) {
	if name == "" {
		return nil, errors.New("store name cannot be blank")
	}

	name = strings.ToLower(name)
	store := p.getLevelDbStore(name)

	if store == nil {
		return p.newLevelDbStore(name)
	}

	return store, nil
}

func (p *Provider) newLevelDbStore(name string) (*store, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	db, err := leveldb.OpenFile(fmt.Sprintf(pathPattern, p.dbPath, name), nil)
	if err != nil {
		return nil, err
	}

	store := &store{db: db, name: name, close: p.removeStore}
	p.dbs[name] = store
	return store, nil
}

func (p *Provider) getLevelDbStore(name string) *store {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return p.dbs[name]
}

func (p *Provider) removeStore(name string) {
	p.lock.Lock()
	defer p.lock.Unlock()

	_, ok := p.dbs[name]

	if ok {
		delete(p.dbs, name)
	}
}

func (p *Provider) GetOpenStores() []spi.Store {
	p.lock.RLock()
	defer p.lock.RUnlock()

	openedStores := make([]spi.Store, len(p.dbs))
	var counter int

	for _, db := range p.dbs {
		openedStores[counter] = db
		counter++
	}

	return openedStores
}

func (p *Provider) Close() error {
	p.lock.RLock()

	openedStores := make([]*store, len(p.dbs))
	var counter int

	for _, openStore := range p.dbs {
		openedStores[counter] = openStore
		counter++
	}

	p.lock.RUnlock()

	for _, openStore := range openedStores {
		err := openStore.Close()
		if err != nil {
			return fmt.Errorf(`failed to close open store with name "%s": %w`, openStore.name, err)
		}
	}

	return nil
}

func (s *store) Put(key string, value []byte, tags ...spi.Tag) error {
	if key == "" {
		return errors.New("key cannot be blank")
	}

	if value == nil {
		return errors.New("value cannot be nil")
	}

	for _, tag := range tags {
		if strings.Contains(tag.Name, ":") {
			return fmt.Errorf("invalid tag name: %s", tag.Name)
		}

		if strings.Contains(tag.Value, ":") {
			return fmt.Errorf("invalid tag value: %s", tag.Value)
		}
	}

	var newDbEntry dbEntry

	newDbEntry.Value = value

	if len(tags) > 0 {
		newDbEntry.Tags = tags

		err := s.updateTagMap(key, tags)

		if err != nil {
			return fmt.Errorf("failed to update tag map: %w", err)
		}
	}

	entryBytes, err := json.Marshal(newDbEntry)
	if err != nil {
		return fmt.Errorf("failed to marshal new DB entry: %w", err)
	}

	return s.db.Put([]byte(key), entryBytes, nil)
}

func (s *store) Get(key string) ([]byte, error) {
	retriveEntry, err := s.getDbEntry(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get DB entry: %w", err)
	}
	return retriveEntry.Value, nil
}

func (s *store) GetTags(key string) ([]spi.Tag, error) {
	retriveEntry, err := s.getDbEntry(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get DB entry: %w", err)
	}
	return retriveEntry.Tags, nil
}

func (s *store) GetBulk(keys ...string) ([][]byte, error) {
	if len(keys) == 0 {
		return nil, errors.New("keys slice must contain at least one key")

	}

	values := make([][]byte, len(keys))
	for i, key := range keys {
		var err error
		values[i], err = s.Get(key)

		if err != nil {
			if errors.Is(err, spi.ErrDataNotFound) {
				continue
			}
			return nil, fmt.Errorf("unexpected failure while retrieving the value stored under %s: %w", key, err)
		}
	}
	return values, nil
}

func (s *store) Query(expression string, options ...spi.QueryOption) (spi.Iterator, error) {
	err := checkForUnsupportedQueryOptions(options)
	if err != nil {
		return nil, err
	}
	if expression == "" {
		return nil, fmt.Errorf(invalidQueryExpressionFormat, expression)
	}
	expressionSlit := strings.Split(expression, ":")
	var expressionTagName string
	var expressionTagValue string

	switch len(expressionSlit) {
	case expressionTagNameOnlyLength:
		expressionTagName = expressionSlit[0]
	case expressionTagNameAndValueLength:
		expressionTagName = expressionSlit[0]
		expressionTagValue = expressionSlit[1]
	default:
		return nil, fmt.Errorf(invalidQueryExpressionFormat, expression)
	}

	matchingDatabaseKeys, err := s.getDatabaseKeyMatchingQuery(expressionTagName, expressionTagValue)
	if err != nil {
		return nil, fmt.Errorf("failed to get database keys matching query: %w", err)
	}
	return &iterator{keys: matchingDatabaseKeys, store: s}, nil
}

func (s *store) Delete(key string) error {
	if key == "" {
		return errors.New("key cannot be blank")
	}

	err := s.db.Delete([]byte(key), nil)
	if err != nil {
		return fmt.Errorf("failed to delete from underlying database")
	}

	err = s.removeFromTagMap(key)

	if err != nil {
		return fmt.Errorf("failed to remove key from tag map: %w", err)
	}

	return nil
}

func (s *store) Batch(operations []spi.Operation) error {
	if len(operations) == 0 {
		return errors.New("batch requires at least one operation")
	}
	for _, operation := range operations {
		if operation.Value == nil {
			err := s.Delete(operation.Key)
			if err != nil {
				return fmt.Errorf("failed to delete value: %w", err)
			}
		} else {
			err := s.Put(operation.Key, operation.Value, operation.Tags...)
			if err != nil {
				return fmt.Errorf("failed to put value: %w", err)
			}
		}
	}
	return nil
}

func (s *store) Flush() error {
	return nil
}

func (s *store) Close() error {
	s.close(s.name)
	err := s.db.Close()
	if err != nil {
		if err.Error() != "leveldb: closed" {
			return err
		}
	}
	return nil
}

func (s *store) updateTagMap(key string, tags []spi.Tag) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	tagMap, err := s.getTagMap(true)
	if err != nil {
		return fmt.Errorf("failed to get tag map: %w", err)
	}

	for _, tag := range tags {
		if tagMap[tag.Name] == nil {
			tagMap[tag.Name] = make(map[string]struct{})
		}
		tagMap[tag.Name][key] = struct{}{}
	}

	tagMapBytes, err := json.Marshal(tagMap)
	if err != nil {
		return fmt.Errorf("failed to marshal updated tag map: %w", err)
	}

	err = s.Put(tagMapKey, tagMapBytes)
	if err != nil {
		return fmt.Errorf("failed to put updated tag map into the store: %w", err)
	}
	return nil
}

func (s *store) getDbEntry(key string) (dbEntry, error) {
	if key == "" {
		return dbEntry{}, errors.New("key cannot be blank")

	}

	retrievedEntryBytes, err := s.db.Get([]byte(key), nil)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {

			return dbEntry{}, spi.ErrDataNotFound
		}
		return dbEntry{}, err
	}
	var retriveEntry dbEntry

	err = json.Unmarshal(retrievedEntryBytes, &retriveEntry)
	if err != nil {
		return dbEntry{}, fmt.Errorf("failed to unmarshal retrieved DB entry: %w", err)
	}
	return retriveEntry, nil
}

func (s *store) removeFromTagMap(key string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	tagMap, err := s.getTagMap(false)
	if err != nil {
		if errors.Is(err, spi.ErrDataNotFound) {
			return nil
		}
		return fmt.Errorf("failed to get tag map: %w", err)
	}

	for _, tagNameToKeys := range tagMap {
		delete(tagNameToKeys, key)
	}

	tagMapBytes, err := json.Marshal(tagMap)
	if err != nil {
		return fmt.Errorf("failed to marshal updated tag map: %w", err)
	}

	err = s.Put(tagMapKey, tagMapBytes)

	if err != nil {
		return fmt.Errorf("failed to put updated tag map back into the store: %w", err)
	}

	return nil
}

func (s *store) getDatabaseKeyMatchingQuery(expressionTagName, expressionTagValue string) ([]string, error) {
	tagMag, err := s.getTagMap(false)
	if err != nil {
		if errors.Is(err, spi.ErrDataNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to get tag map: %w", err)
	}
	if expressionTagValue == "" {
		return getDatabaseKeyMatchingTagName(tagMag, expressionTagName), nil
	}

	matchingDatabaseKeys, err := s.getDatabaseKeyMatchingTagNameAndValue(tagMag, expressionTagName, expressionTagValue)
	if err != nil {
		return nil, fmt.Errorf("failed to get database keys matching tag name and value: %w", err)
	}
	return matchingDatabaseKeys, nil
}

func (s *store) getDatabaseKeyMatchingTagNameAndValue(tagMap tagMapping,
	expressionTagName, expressionTagValue string) ([]string, error) {

	var matchingDatabaseKeys []string
	for tagName, databaseKeysSet := range tagMap {
		if tagName == expressionTagName {
			for databaseKey := range databaseKeysSet {
				tags, err := s.GetTags(databaseKey)
				if err != nil {
					return nil, fmt.Errorf("failed to get tags: %w", err)
				}

				for _, tag := range tags {
					if tag.Name == expressionTagName && tag.Value == expressionTagValue {
						matchingDatabaseKeys = append(matchingDatabaseKeys, databaseKey)
						break
					}
				}
			}
			break
		}
	}
	return matchingDatabaseKeys, nil
}

func (s *store) getTagMap(createIfDoesNotExist bool) (tagMapping, error) {
	tagMapBytes, err := s.Get(tagMapKey)
	if err != nil {
		if createIfDoesNotExist && errors.Is(err, spi.ErrDataNotFound) {
			err = s.Put(tagMapKey, []byte("{}"))
			if err != nil {
				return nil, fmt.Errorf(`failed to create tag map for "%s": %w`, s.name, err)
			}

			tagMapBytes = []byte("{}")

		} else {
			return nil, fmt.Errorf("failed to unmarshal tag map bytes: %w", err)
		}
	}
	var tagMap tagMapping
	err = json.Unmarshal(tagMapBytes, &tagMap)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal tag map bytes: %w", err)
	}
	return tagMap, nil
}

type iterator struct {
	keys         []string
	currentIndex int
	currentKey   string
	store        *store
}

func (i *iterator) Next() (bool, error) {
	if len(i.keys) == i.currentIndex || len(i.keys) == 0 {
		if len(i.keys) == i.currentIndex || len(i.keys) == 0 {
			return false, nil
		}
	}

	i.currentKey = i.keys[i.currentIndex]
	i.currentIndex++
	return true, nil
}

func (i *iterator) Key() (string, error) {

	return i.currentKey, nil
}

func (i *iterator) Value() ([]byte, error) {
	value, err := i.store.Get(i.currentKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get value from store: %w", err)
	}
	return value, nil
}

func (i *iterator) Tags() ([]spi.Tag, error) {
	tags, err := i.store.GetTags(i.currentKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get tags from store: %w", err)
	}

	return tags, nil
}
func (i *iterator) TotalItems() (int, error) {
	return len(i.keys), nil
}

func (i *iterator) Close() error {
	return nil
}
func getQueryOptions(ops []spi.QueryOption) spi.QueryOptions {

	var queryOptions spi.QueryOptions

	for _, option := range ops {
		option(&queryOptions)
	}
	return queryOptions
}

func checkForUnsupportedQueryOptions(ops []spi.QueryOption) error {
	queySettings := getQueryOptions(ops)
	if queySettings.InitialPageNum != 0 {
		return errors.New("levelDB provider doesn not currently support " +
			"setting the initial page number of query results")
	}
	if queySettings.SortOptions != nil {
		return errors.New("levelDB provider does not currently support custom sort options for query results")
	}
	return nil
}

func getDatabaseKeyMatchingTagName(tagMap tagMapping, expressionTagName string) []string {
	var matchingDatabaseKeys []string
	for tagName, databaseKeysSet := range tagMap {
		if tagName == expressionTagName {
			for databaseKey := range databaseKeysSet {
				matchingDatabaseKeys = append(matchingDatabaseKeys, databaseKey)
			}
			break
		}
	}
	return matchingDatabaseKeys
}
