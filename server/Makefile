
run:
	docker-compose --project-name scylla up -d

logs:
	docker-compose logs -f
	
status:
	docker exec -it some-scylla nodetool status

cqlsh:
	docker exec -it some-scylla cqlsh
