package config

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = user
	return writeToFile(*cfg)
}
