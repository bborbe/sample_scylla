module github.com/bborbe/sample_scylla

go 1.21.3

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.11.1

require github.com/gocql/gocql v1.11.1

require (
	github.com/golang/glog v1.1.2 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)
