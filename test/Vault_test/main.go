package main

import (
	"log"

	vault "github.com/bytehubplus/fusion/node/vaultindex"
)

func main() {
	TestRegisterVault()
}

func TestRegisterVault() {
	conf := vault.Config{
		Scheme: "did",
		Method: "rich",
		DBPath: "./data/vaultindex",
	}

	provider, _ := vault.NewProvider(conf)
	//defer provider.CloseDB()
	vaultID, err := provider.RegisterVault("abcdef1234567")
	if err != nil {
		log.Printf("register failed: %s", err)
	}

	log.Printf("vauld %s registered", vaultID)
}

func TestVaultExist() {
	conf := vault.Config{
		Scheme: "did",
		Method: "rich",
		DBPath: "./data/vaultindex",
	}

	provider, _ := vault.NewProvider(conf)
	//defer provider.CloseDB()
	vaultID := provider.GenerateVaultID("abcdef1234567")
	exist := provider.VaultExits(vaultID)
	if exist {
		log.Printf("Vault %s already registered", vaultID)
	} else {

		log.Printf("vault %s not registered", vaultID)
	}
}

func TestUnregisterVault() {
	conf := vault.Config{
		Scheme: "did",
		Method: "rich",
		DBPath: "./data/vaultindex",
	}

	provider, _ := vault.NewProvider(conf)
	//defer provider.CloseDB()
	vaultID := provider.GenerateVaultID("abcdef1234567")
	if err := provider.UnregisterVault(vaultID); err != nil {
		log.Printf("unregister vault %s failed", err)
	} else {
		log.Printf("vault %s unregistered", vaultID)
	}
}
