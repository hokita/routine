sql:
	docker-compose exec db psql -U app -d routine

restart:
	docker-compose restart
