# RMQ Prometheus Metrics

Collects [RMQ](https://github.com/adjust/rmq) stats and exposes them as Prometheus Metrics.

 ## How to use it 
 
 ```go
package main

import (
        "net/http"
        "time"

        "github.com/adjust/rmq"
        "github.com/pffreitas/rmqprom"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
        var connection rmq.Connection // get a RMQ Connection
        rmqprom.RecordRmqMetrics(connection)

        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe(":2112", nil)
}
```

## How does it look like

For each queue we expose (all as Gauge):

- Connection count
- Consumer count
- Ready messages count
- Rejected messages count
- Unacked messages count


```shell script
# HELP rmq_connection_count Number of connections consuming a queue
# TYPE rmq_connection_count counter
rmq_connection_count{queue="queue-name"} 0
# HELP rmq_consumer_count Number of consumers consuming messages for a queue
# TYPE rmq_consumer_count counter
rmq_consumer_count{queue="queue-name"} 0
# HELP rmq_ready_count Number of ready messages on queue
# TYPE rmq_ready_count counter
rmq_ready_count{queue="queue-name"} 0
# HELP rmq_rejected_count Number of rejected messages on queue
# TYPE rmq_rejected_count counter
rmq_rejected_count{queue="queue-name"} 0
# HELP rmq_unacked_count Number of unacked messages on a consumer
# TYPE rmq_unacked_count counter
rmq_unacked_count{queue="queue-name"} 0
```