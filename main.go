package main

import (
	"time"

	"github.com/adjust/rmq"
	"github.com/prometheus/client_golang/prometheus"
)

type queueStatsCounters struct {
	readyCount      prometheus.Counter
	rejectedCount   prometheus.Counter
	connectionCount prometheus.Counter
	consumerCount   prometheus.Counter
	unackedCount    prometheus.Counter
}

func RecordRmqMetrics(connection rmq.Connection) {
	counters := registerCounters(connection)

	go func() {
		for {
			stats := connection.CollectStats(connection.GetOpenQueues())
			for queue, queueStats := range stats.QueueStats {
				if counter, ok := counters[queue]; ok {
					counter.readyCount.Add(float64(queueStats.ReadyCount))
					counter.rejectedCount.Add(float64(queueStats.RejectedCount))
					counter.connectionCount.Add(float64(queueStats.ConnectionCount()))
					counter.consumerCount.Add(float64(queueStats.ConsumerCount()))
					counter.unackedCount.Add(float64(queueStats.UnackedCount()))
				}
			}

			time.Sleep(1 * time.Second)
		}
	}()
}

func registerCounters(connection rmq.Connection) map[string]queueStatsCounters {
	counters := map[string]queueStatsCounters{}

	for _, queue := range connection.GetOpenQueues() {
		counters[queue] = queueStatsCounters{
			readyCount: prometheus.NewCounter(prometheus.CounterOpts{
				Namespace:   "rmq",
				Name:        "rmq_ready_count",
				Help:        "Number of ready messages on queue",
				ConstLabels: prometheus.Labels{"queue": queue},
			}),
			rejectedCount: prometheus.NewCounter(prometheus.CounterOpts{
				Namespace:   "rmq",
				Name:        "rmq_rejected_count",
				Help:        "Number of rejected messages on queue",
				ConstLabels: prometheus.Labels{"queue": queue},
			}),
			connectionCount: prometheus.NewCounter(prometheus.CounterOpts{
				Namespace:   "rmq",
				Name:        "rmq_connection_count",
				Help:        "Number of connections consuming a queue",
				ConstLabels: prometheus.Labels{"queue": queue},
			}),
			consumerCount: prometheus.NewCounter(prometheus.CounterOpts{
				Namespace:   "rmq",
				Name:        "rmq_consumer_count",
				Help:        "Number of consumers consuming messages for a queue",
				ConstLabels: prometheus.Labels{"queue": queue},
			}),
			unackedCount: prometheus.NewCounter(prometheus.CounterOpts{
				Namespace:   "rmq",
				Name:        "rmq_unacked_count",
				Help:        "Number of unacked messages on a consumer",
				ConstLabels: prometheus.Labels{"queue": queue},
			}),
		}
	}

	return counters
}
