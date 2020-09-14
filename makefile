.PHONY: front app
sql:
	docker-compose exec db psql -U app -d routine

front:
	docker-compose exec front bash

app:
	docker-compose exec app bash
