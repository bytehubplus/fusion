package config

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	confGot := NewConfig("config", "yaml", "../../conf")
	confWant := Config{FileName: "config",
		FileType: "yaml",
		FilePath: "../../conf"}

	if confGot != confWant {
		t.Errorf(`NewConfig("config", "yaml","../../conf") got %q, want %q`, confGot, confWant)
	}
}

func TestGet(t *testing.T) {

	//test content
	nodesWant := map[string]string{"node1": "192.168.0.2:4000", "node2": "192.168.0.3:4000"}
	vaultIndexWant := map[string]string{"path": "./data"}
	logFileWant := map[string]string{"path": "./logfile"}
	pluginWant := map[string]string{"path": "./plugin"}

	conf1 := Config{FileName: "config",
		FileType: "yaml",
		FilePath: "../../conf"}

	nodesGot, _ := conf1.GetNodes()
	vaultIndexGot, _ := conf1.GetVaultIndex()
	logFileGot, _ := conf1.GetLogFile()
	pluginGot, _ := conf1.GetPlugin()

	var eq bool
	eq = reflect.DeepEqual(nodesWant, nodesGot)
	if eq == false {
		t.Errorf("GetNodes() got %q, want %q", nodesGot, nodesWant)
	}
	eq = reflect.DeepEqual(vaultIndexWant, vaultIndexGot)
	if eq == false {
		t.Errorf("GetVaultIndex() got %q, want %q", vaultIndexGot, vaultIndexWant)
	}
	eq = reflect.DeepEqual(logFileWant, logFileGot)
	if eq == false {
		t.Errorf("GetLogFile() got %q, want %q", logFileGot, logFileWant)
	}
	eq = reflect.DeepEqual(pluginWant, pluginGot)
	if eq == false {
		t.Errorf("GetPlugin() got %q, want %q", pluginGot, pluginWant)
	}

	//test err
	conf2 := Config{FileName: "",
		FileType: "",
		FilePath: ""}

	_, err1 := conf2.GetNodes()
	if err1 == nil {
		t.Error("conf2.GetNodes() wants error")
	}
	_, err2 := conf2.GetLogFile()
	if err2 == nil {
		t.Error("conf2.GetLogFile() wants error")
	}
	_, err3 := conf2.GetVaultIndex()
	if err3 == nil {
		t.Error("conf2.GetVaultIndex() wants error")
	}
	_, err4 := conf2.GetPlugin()
	if err4 == nil {
		t.Error("conf2.GetPlugin() wants error")
	}
}
