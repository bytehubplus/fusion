package common

type Config struct {
	FileName string
	FileType string
	FilePath string
}

func NewConfig(fileName string, fileType string, filePath string) Config {
  conf := Config{FileName: fileName,
		FileType: fileType,
		FilePath: filePath}

   viper. 
}

func (conf *Config) GetNodes() (map[string]string, error) {
	viper.GetStringMapString("nodes")
}
