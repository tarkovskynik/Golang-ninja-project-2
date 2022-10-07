package mongodb

type Config struct {
	uri  string // "MONGO_URI"
	user string // "MONGO_USER"
	pass string // "MONGO_PASSWORD"
}

func NewConfig(uri, user, pass string) *Config {
	return &Config{
		uri:  uri,
		user: user,
		pass: pass,
	}
}
