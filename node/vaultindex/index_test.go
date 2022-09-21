package vaultindex

import "testing"

func TestRegisterVault(t *testing.T) {
	conf := Config{
		Scheme: "did",
		Method: "rich",
		DBPath: "./data/vaultindex",
	}

	provider, _ := NewProvider(conf)

	vaultID, err := provider.RegisterVault("abcdef1234567")
	if err != nil {
		t.Logf("register failed: %s", err)
	}

	t.Logf("vauld %s registered", vaultID)
}

func TestVaultExist(t *testing.T) {
	conf := Config{
		Scheme: "did",
		Method: "rich",
		DBPath: "./data/vaultindex",
	}

	provider, _ := NewProvider(conf)
	vaultID := provider.generateVaultID("abcdef1234567")
	exist := provider.VaultExits(vaultID)
	if exist {
		t.Logf("Vault %s already registered", vaultID)
	} else {

		t.Logf("vault %s not registered", vaultID)
	}
}

func TestUnregisterVault(t *testing.T) {
	conf := Config{
		Scheme: "did",
		Method: "rich",
		DBPath: "./data/vaultindex",
	}

	provider, _ := NewProvider(conf)
	vaultID := provider.generateVaultID("abcdef1234567")
	if err := provider.UnregisterVault(vaultID); err != nil {
		t.Logf("unregister vault %s failed", err)
	} else {
		t.Logf("vault %s unregistered", vaultID)
	}

}
