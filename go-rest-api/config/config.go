package config
 
type Config struct {
	DB *DBConfig
}
 
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	DatabaseName     string
	Charset  string
}
 
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "",
			DatabaseName: "go_rest_api_database",
			Charset: "utf8",
		},
	}
}