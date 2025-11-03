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
	R2Storage struct {
		Endpoint   string `yaml:"endpoint"`
		BucketName string `yaml:"bucketname"`
		Region     string `yaml:"region"`
		AccessKey  string `yaml:"accesskey"`
		SecretKey  string `yaml:"secretkey"`
	} `yaml:"r2storage"`
}
