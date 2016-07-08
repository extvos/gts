package gts_config

type LoggingConfig struct {
	Filename string
	Level string
	Format string
}

type ServerConfig struct {
	Listen string
	Group string
	User string
	Logging *LoggingConfig
}

type CacheConfig struct {
	Id string
	Size Size
	Expires Duration
	Storage string
}

type BackendServerConfig struct {
	Server string
	Weight uint
	Check bool
	Backup bool
	Interval Duration
}

type BackendConfig struct {
	Id string
	Servers []*BackendServerConfig
}

type PatternConfig struct {
	Pattern Regex
	Backends string
	Caches string
	Cache_key string
	Expires Duration
}

type FrontendConfig struct {
	Id string
	Listen string
}

type GTS_Config struct {
	Server *ServerConfig
	Backends []*BackendConfig
	Caches []*CacheConfig
	Frontends []*FrontendConfig
}
