package vaultindex

type VaultIndex interface {
	VaultExits(did string) bool
	// register a vault, return vault ID
	RegisterVault(id string) (string, error)
	// unregiste a vault
	UnregisterVault(id string) error
	// Put data into entryID if entryID exist. if entryID is nil, create a entryID and return it.
	Put(entryID string, data string) (string, error)
	// Get data from entryID and return data
	Get(entryID string) (string, error)
	// Delete data according to entryID
	Delete(entryID string) error
}
