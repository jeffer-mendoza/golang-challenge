
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