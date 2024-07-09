package model

type Config struct {
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Network  string `yaml:"network"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
	Server struct {
		Host  string `yaml:"host"`
		Port  string `yaml:"port"`
		Https bool   `yaml:"https"`
	} `yaml:"server"`
}
