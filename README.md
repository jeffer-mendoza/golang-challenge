
# Install 

```
git clone https://github.com/jeffer-mendoza/golang-challenge.git

go get github.com/goenning/go-cache-demo/cache
go get github.com/goenning/go-cache-demo/cache/memory
go get github.com/goenning/go-cache-demo/cache/redis
go get github.com/prometheus/client_golang/prometheus/promhttp
```

# Config

Create a file in project -> conf.yml

```
host: "ip"
user: "database username"
port: "port number"
name: "database name"
password: "user password"
cachetime: "cache time"
serveport: "serve port"
```

# Resources

[**Server side cache with Go**](https://goenning.net/2017/03/18/server-side-cache-go/)


# Resource Monitoring Tool

## Prometheus

1. Download  [The Prometheus monitoring system](https://prometheus.io/download/)
2. Unzip the file downloaded
3. Create the config.yml configuration file and save it in Prometheus project's root 

``` 
    # my global config
    global:
    scrape_interval:     15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.
    evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.

    scrape_configs:
    - job_name: 'golang'
        metrics_path: /metrics
        static_configs:
        - targets: ['localhost:9005']
```
4. Execute Prometheus
``` 
    ./prometheus --config.file=config.yml
``` 
5. Open Prometheus web application in the web browser at the address: [http://localhost:9090](http://localhost:9090)

## Pprof
https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/
https://golang.org/pkg/net/http/pprof/
https://github.com/google/pprof


1. import _ "net/http/pprof" in the main file


2. Generate profile file 

Each Profile has a unique name. A few profiles are predefined:

goroutine    - stack traces of all current goroutines
heap         - a sampling of all heap allocations
threadcreate - stack traces that led to the creation of new OS threads
block        - stack traces that led to blocking on synchronization primitives
mutex        - stack traces of holders of contended mutexes


goroutine profile
``` 
go tool pprof http://localhost:9006/debug/pprof/goroutine
```

heap profile
``` 
go tool pprof http://localhost:9006/debug/pprof/heap
```

CPU profile
``` 
go tool pprof http://localhost:9006/debug/pprof/profile
```

goroutine blocking profile
``` 
go tool pprof http://localhost:9006/debug/pprof/block
```

this command open the Profile Interactive Mode, here you can generate the
reports:

3. Generate reports(text, images):

```
(pprof)top
```
```
(pprof)png
```



if you want generate images files

```
 sudo apt install python-pydot python-pydot-ng graphviz
```