# LOCAL DEVELOPMENT UTILITY SHELL APPS
migrate-up:
	sql-migrate up -config=local.dbconfig.yml -env="local"

migrate-up-dry:
	sql-migrate up -config=local.dbconfig.yml -env="local" -dryrun

migrate-rollback:
	sql-migrate down -config=local.dbconfig.yml -env="local"
	
up:
	docker compose up -d --remove-orphans --build

down:
	docker compose down -v