package system

import (
	"strings"

	"zenithar.org/go/nikoniko/cmd/server/shared"

	"github.com/Sirupsen/logrus"
	"zenithar.org/go/common/cache"
	"zenithar.org/go/common/eventbus"
)

// Setup the application
func Setup(flags *shared.Flags) Application {
	logrus.Infoln("**********************************************************")
	// Initialize event bus
	logrus.Infoln("Initializing EvenBus : local mode")
	bus := eventbus.NewLocal()

	logrus.Infoln("**********************************************************")
	// Initialize cache manager
	var cacheStore cache.CacheStore
	if len(strings.TrimSpace(flags.MemcachedHosts)) > 0 {
		logrus.Infoln("Initializing CacheManager : memcached")
		cacheStore = cache.NewMemcachedStore(strings.Split(flags.MemcachedHosts, ","), cache.DEFAULT)
	} else if len(strings.TrimSpace(flags.RedisHost)) > 0 {
		logrus.Infoln("Initializing CacheManager : redis")
		cacheStore = cache.NewRedisCache(flags.RedisHost, "", cache.DEFAULT)
	} else {
		logrus.Infoln("Initializing CacheManager : inMemory")
		cacheStore = cache.NewInMemoryStore(cache.DEFAULT)
	}

	return &baseApplication{
		config:     shared.Config,
		bus:        bus,
		cacheStore: cacheStore,
	}
}
