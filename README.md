# Scylla DB

https://hub.docker.com/r/scylladb/scylla/


```bash
docker run --name scylla --hostname scylla -d scylladb/scylla:5.2.9 --smp 1
docker logs scylla -f
```

docker exec -it some-scylla nodetool status

docker exec -it some-scylla cqlsh

https://manager.docs.scylladb.com/stable/docker/

https://university.scylladb.com/setup-a-scylla-cluster/
https://university.scylladb.com/courses/using-scylla-drivers/lessons/golang-and-scylla-part-1/

https://hub.docker.com/r/scylladb/scylla/tags

https://github.com/scylladb/gocql

https://github.com/scylladb/gocqlx