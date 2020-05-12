module github.com/pffreitas/rmqprom

go 1.13

require (
	github.com/adjust/rmq/v2 v2.0.0
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/onsi/gomega v1.10.0 // indirect
	github.com/prometheus/client_golang v1.6.0
	gopkg.in/bsm/ratelimit.v1 v1.0.0-20160220154919-db14e161995a // indirect
	gopkg.in/redis.v3 v3.6.4 // indirect
)

replace github.com/adjust/rmq/v2 => github.com/pffreitas/rmq/v2 v2.0.0
