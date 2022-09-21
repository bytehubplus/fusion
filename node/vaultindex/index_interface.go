package vaultindex

type VaultIndex interface {
	VaultExits(did string) bool
	// register a vault, return vault ID
	RegisterVault(id string) (string, error)
	// unregiste a vault
	UnregisterVault(id string) error
}
