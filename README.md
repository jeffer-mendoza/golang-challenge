
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


## Pprof

* https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/
* https://golang.org/pkg/net/http/pprof/
* https://github.com/google/pprof

(![Pprof Diagram](https://raw.githubusercontent.com/jeffer-mendoza/golang-challenge/pprof/pprof.png))

1. import _ "net/http/pprof" in the main file


2. Generate profile file 

Each Profile has a unique name. A few profiles are predefined:

* goroutine    - stack traces of all current goroutines
* heap         - a sampling of all heap allocations
* threadcreate - stack traces that led to the creation of new OS threads
* block        - stack traces that led to blocking on synchronization primitives
* mutex        - stack traces of holders of contended mutexes


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

text report
```
(pprof)top
```
image report -> if you want generate images files sudo apt install python-pydot python-pydot-ng graphviz

```
(pprof)png
```

html report
```
(pprof)web
```




