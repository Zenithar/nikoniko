package main

import (
	"flag"
	"os"
	"strings"
	"time"

	"esec.sogeti.com/crbr/shared"
	"esec.sogeti.com/crbr/version"

	"github.com/Sirupsen/logrus"
	raven "github.com/getsentry/raven-go"
	"zenithar.org/go/common/logging/logrus/hooks"
)

var (
	// General flags
	bindAddress      = flag.String("bind", ":5000", "Network address used to bind")
	logFormatterType = flag.String("log", "text", "Log formatter type. Either \"json\" or \"text\"")
	logLevel         = flag.String("log_level", "info", "Defines the log level (panic, fatal, error, warn, info, debug)")
	forceColors      = flag.Bool("force_colors", false, "Force colored prompt?")
	ravenDSN         = flag.String("raven_dsn", "", "Defines the sentry endpoint dsn")

	// Database
	databaseDriver    = flag.String("db_driver", "rethinkdb", "Specify the database to use (mongodb, rethinkdb)")
	databaseHost      = flag.String("db_host", "localhost:28015", "Database hosts, split by ',' to add several hosts")
	databaseNamespace = flag.String("db_namespace", "portal", "Select the database")
	databaseUser      = flag.String("db_user", "", "Database user")
	databasePassword  = flag.String("db_password", "", "Database user password")

	// Developer mode
	devMode = flag.Bool("dev", false, "Enable developer mode")

	// Public URL
	publicURL = flag.String("public_url", "https://login.esec.fr.capgemini.com", "Defines the public url")

	// Cookie
	cookieSecure     = flag.Bool("cookie_secure", true, "Secure flag for cookies")
	cookieKey        = flag.String("cookie_key", ":|c}eico/t/H?Tj{@$UJU6@kGo^W@vfs", "Cookie key for signing (32 bytes long)")
	cookieExpiration = flag.Int("cookie_expiration", 60, "Cookie expiration in minutes")

	// Cache
	memcachedHosts = flag.String("memcached_hosts", "", "Memcached servers for cache (ex: 127.0.0.1:11211)")
	redisHost      = flag.String("redis_host", "", "Redis server for cache")
)

func init() {
	flag.Parse()

	// Set localtime to UTC
	time.Local = time.UTC

	logrus.Infoln("**********************************************************")
	logrus.Infoln("NikoNiko server starting ...")
	logrus.Infof("Version : %s (%s-%s)", version.Version, version.Revision, version.Branch)

	// Set the formatter depending on the passed flag's value
	if *logFormatterType == "text" {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors: *forceColors,
		})
	} else if *logFormatterType == "json" {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors: *forceColors,
		})
	}

	// Defines the log level
	level, err := logrus.ParseLevel(*logLevel)
	if err != nil {
		logrus.Fatalln("Invalid log level ! (panic, fatal, error, warn, info, debug) ")
	}
	logrus.SetLevel(level)

	// Connect to raven
	var rc *raven.Client
	if len(strings.TrimSpace(*ravenDSN)) > 0 {
		logrus.Infoln("**********************************************************")
		logrus.Infoln("Initializing Sentry client")
		logrus.Infof(" DSN : %s", *ravenDSN)
		h, err := os.Hostname()
		if err != nil {
			logrus.Fatal(err)
		}
		rc, err = raven.NewClient(*ravenDSN, map[string]string{
			"hostname": h,
			"app":      "nikoniko",
			"version":  version.Version,
			"revision": version.Revision,
			"branch":   version.Branch,
		})
		if err != nil {
			logrus.Fatal(err)
		}

		// Sentry hook
		sentryHook, err := hooks.NewWithClientSentryHook(rc, []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		})
		if err == nil {
			logrus.AddHook(sentryHook)
		}

		shared.Raven = rc
	}
}

func main() {

}
