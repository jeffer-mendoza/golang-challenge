package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

    "github.com/goenning/go-cache-demo/cache"
	"github.com/goenning/go-cache-demo/cache/memory"
	"github.com/goenning/go-cache-demo/cache/redis"
)

var storage cache.Storage
var config conf

func init() {
	strategy := flag.String("s", "memory", "Cache strategy (memory or redis)")
	flag.Parse()

	if *strategy == "memory" {
		storage = memory.NewStorage()
	} else if *strategy == "redis" {
		var err error
		if storage, err = redis.NewStorage(os.Getenv("REDIS_URL")); err != nil {
			panic(err)
		}
	} else {
		panic(fmt.Sprintf("Invalid cache strategy %s.", *strategy))
	}
}


func main() {
    config.getConf()
	loadRoutes()
	http.ListenAndServe(config.Serveport, nil)
}
