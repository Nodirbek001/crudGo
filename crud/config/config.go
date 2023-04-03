package config
type Config struct{
	HTTPPort string
	Timeout int
}
func Load() Config{
	config :=Config{":8083",7}
	return config
}