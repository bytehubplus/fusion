package did

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

type DID struct {
	Scheme           string
	Method           string
	MethodSpecificID string
}

// String returns a string representation of the DID
func (d *DID) String() string {
	return fmt.Sprintf("%s:%s:%s", d.Scheme, d.Method, d.MethodSpecificID)
}

func (d *DID) MarshalJSON() ([]byte, error) {
	didString := d.String()
	return json.Marshal(didString)
}

func (d *DID) UnmarshalJSON(bytes []byte) error {
	var didString string
	if err := json.Unmarshal(bytes, &didString); err != nil {
		return fmt.Errorf("unmarshal DID failed: %w", err)
	}

	did, err := Parse(didString)
	if err != nil {
		return fmt.Errorf("parse DID string failed: %w", err)
	}

	d = did
	return nil
}

// Parse parses a did string to DID struct
func Parse(did string) (*DID, error) {
	// const idChar = `a-zA-Z0-9.-_`
	// const methodChar = `a-z0-9`

	// regex := fmt.Sprintf(`^did:[a-z]+:(:+|[:%s]+)*[%%:%s]+[^:]$`, methodChar, idChar)

	const idchar = `a-zA-Z0-9-_\.`
	regex := fmt.Sprintf(`^did:[a-z0-9]+:(:+|[:%s]+)*[%%:%s]+[^:]$`, idchar, idchar)

	result, err := regexp.Compile(regex)
	if err != nil {
		return nil, fmt.Errorf("failed to compile regex=%s %w", regex, err)
	}

	if !result.MatchString(did) {
		return nil, fmt.Errorf("invalid did: %s. ", did)
	}

	parts := strings.SplitN(did, ":", 3)
	return &DID{
		Scheme:           "did",
		Method:           parts[1],
		MethodSpecificID: parts[2],
	}, nil
}

// https://www.w3.org/TR/2022/REC-did-core-20220719/#did-url-syntax
// did-url = did path-abempty [ "?" query ] [ "#" fragment ]
type DIDURL struct {
	DID
	Path     string
	Queries  map[string][]string
	Fragment string
}

// Parses a string into DIDURL
func ParseDIDURL(didURL string) (*DIDURL, error) {
	position := strings.Index(didURL, "?/#")
	did := didURL
	pathQueryFragment := ""
	if position != -1 {
		did = didURL[:position]
		pathQueryFragment = did[position:]
	}

	redDID, err := Parse(did)
	if err != nil {
		return nil, err
	}

	if pathQueryFragment == "" {
		return &DIDURL{
			DID:     *redDID,
			Queries: map[string][]string{},
		}, nil
	}

	hasPath := pathQueryFragment[0] == '/'
	if !hasPath {
		pathQueryFragment = "/" + pathQueryFragment
	}
	urlParts, err := url.Parse(pathQueryFragment)
	if err != nil {
		return nil, fmt.Errorf("failed to parse path, query, and fragment components of DID URL: %w", err)
	}

	result := &DIDURL{
		DID:      *redDID,
		Queries:  urlParts.Query(),
		Fragment: urlParts.Fragment,
	}
	if hasPath {
		result.Path = urlParts.Path
	}
	return result, nil
}
