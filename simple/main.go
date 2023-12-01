package main

import (
	"context"
	"flag"
	"runtime"
	"time"

	"github.com/gocql/gocql"
	"github.com/golang/glog"
)

func main() {
	ctx := context.Background()
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	runtime.GOMAXPROCS(runtime.NumCPU())
	_ = flag.Set("logtostderr", "true")
	_ = flag.Set("v", "2")

	time.Local = time.UTC
	glog.V(2).Infof("set global timezone to UTC")

	session, err := gocql.NewCluster("127.0.0.1:9042").CreateSession()
	if err != nil {
		glog.Exitf("create session failed: %v", err)
	}
	defer session.Close()

	{
		// https://opensource.docs.scylladb.com/stable/cql/ddl.html#id2
		q := `
			CREATE KEYSPACE IF NOT EXISTS myks 
			WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1}
			`
		query := session.Query(q)
		if err := query.Exec(); err != nil {
			glog.Exitf("exec query failed: %v", err)
		}

	}

	{
		q := `DROP TABLE IF EXISTS myks.mytable`

		query := session.Query(q)
		if err := query.Exec(); err != nil {
			glog.Exitf("exec query failed: %v", err)
		}
	}

	{
		// https://opensource.docs.scylladb.com/stable/cql/ddl.html#create-table-statement
		q := `CREATE TABLE IF NOT EXISTS myks.mytable (
				k text PRIMARY KEY,
				v text,
			)`

		query := session.Query(q)
		if err := query.Exec(); err != nil {
			glog.Exitf("exec query failed: %v", err)
		}
	}

	{
		q := `INSERT INTO myks.mytable (k, v) VALUES (?, ?)`
		if err := session.Query(q, "a", "b").WithContext(ctx).Exec(); err != nil {
			glog.Exitf("exec query failed: %v", err)
		}
	}

	{
		q := `
			SELECT k, v FROM myks.mytable
			`
		scanner := session.Query(q).WithContext(ctx).Iter().Scanner()

		for scanner.Next() {
			var (
				key   string
				value string
			)

			if err := scanner.Scan(&key, &value); err != nil {
				glog.Exitf("scan failed: %v", err)
			}
			glog.V(2).Infof("key %s value %s", key, value)
		}
		if err := scanner.Err(); err != nil {
			glog.Exitf("close scanner failed: %v", err)
		}
	}

	glog.V(2).Info("done")
}
