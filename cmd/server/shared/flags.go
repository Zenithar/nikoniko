package shared

// Flags is the command-line parameters holder
type Flags struct {
	BindAddress      string
	RavenDSN         string
	LogFormatterType string
	ForceColors      bool

	DatabaseDriver    string
	DatabaseHost      string
	DatabaseNamespace string
	DatabaseUser      string
	DatabasePassword  string

	DevMode   bool
	PublicURL string

	CookieSecure     bool
	CookieKey        string
	CookieExpiration int

	MemcachedHosts string
	RedisHost      string
}
