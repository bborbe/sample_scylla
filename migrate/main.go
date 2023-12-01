package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/gocql/gocql"
	"github.com/golang/glog"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/migrate"
)

//go:embed *.cql
var Files embed.FS

func main() {
	ctx := context.Background()
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	runtime.GOMAXPROCS(runtime.NumCPU())
	_ = flag.Set("logtostderr", "true")
	_ = flag.Set("v", "2")

	time.Local = time.UTC
	glog.V(2).Infof("set global timezone to UTC")

	// Create gocqlxSession using the keyspace
	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = "myks"

	session, err := cluster.CreateSession()
	if err != nil {
		glog.Exitf("create session failed: %v", err)
	}
	gocqlxSession := gocqlx.NewSession(session)
	defer gocqlxSession.Close()

	glog.V(2).Infof("gocqlxSession created")

	migrate.Callback = func(ctx context.Context, session gocqlx.Session, ev migrate.CallbackEvent, name string) error {
		glog.V(2).Infof("name: %v", name)
		return nil
	}

	if err := migrate.FromFS(ctx, gocqlxSession, Files); err != nil {
		glog.Exitf("migrate failed: %v", err)
	}

	list, err := migrate.List(ctx, gocqlxSession)
	if err != nil {
		glog.Exitf("list migrate failed: %v", err)
	}

	for _, l := range list {
		fmt.Printf("l: %v\n", l)
	}

	{
		q := `
			SELECT id FROM myks.bar
			`
		scanner := session.Query(q).WithContext(ctx).Iter().Scanner()

		for scanner.Next() {
			var (
				id string
			)

			if err := scanner.Scan(&id); err != nil {
				glog.Exitf("scan failed: %v", err)
			}
			glog.V(2).Infof("id %s", id)
		}
		if err := scanner.Err(); err != nil {
			glog.Exitf("close scanner failed: %v", err)
		}
	}

	glog.V(2).Info("done")
}
