package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	FileName string // "config"
	FileType string // "yaml"
	FilePath string // "./conf"
}

// viper instance
var vp *viper.Viper

func NewConfig(fileName, fileType, filePath string) Config {
	//init Config
	conf := Config{FileName: fileName,
		FileType: fileType,
		FilePath: filePath}
	return conf
}

// set config file and get nodes
func (conf *Config) GetNodes() (map[string]string, error) {

	//viper set config file
	vp = viper.New()
	vp.SetConfigName(conf.FileName)
	vp.SetConfigType(conf.FileType)
	vp.AddConfigPath(conf.FilePath)

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	//get nodes
	nodes := vp.GetStringMapString("nodes")
	return nodes, nil

}

// set config file and get vaultIndex
func (conf *Config) GetVaultIndex() (map[string]string, error) {

	//viper set config file
	vp = viper.New()
	vp.SetConfigName(conf.FileName)
	vp.SetConfigType(conf.FileType)
	vp.AddConfigPath(conf.FilePath)

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	//get vault
	vaultIndex := vp.GetStringMapString("vaultIndex")
	return vaultIndex, nil

}

// set config file and get logFile
func (conf *Config) GetLogFile() (map[string]string, error) {

	//viper set config file
	vp = viper.New()
	vp.SetConfigName(conf.FileName)
	vp.SetConfigType(conf.FileType)
	vp.AddConfigPath(conf.FilePath)

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	//get logFile
	logFile := vp.GetStringMapString("logFile")
	return logFile, nil

}

// set config file and get plugin
func (conf *Config) GetPlugin() (map[string]string, error) {

	//viper set config file
	vp = viper.New()
	vp.SetConfigName(conf.FileName)
	vp.SetConfigType(conf.FileType)
	vp.AddConfigPath(conf.FilePath)

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	//get plugin
	plugin := vp.GetStringMapString("plugin")
	return plugin, nil

}
