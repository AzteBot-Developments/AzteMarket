# LOCAL DEVELOPMENT UTILITY SHELL APPS
migrate-up:
	sql-migrate up -config=local.dbconfig.yml -env="local-aztemarket"
	sql-migrate up -config=local.dbconfig.yml -env="local-aztebot"

migrate-up-dry:
	sql-migrate up -config=local.dbconfig.yml -env="local-aztemarket" -dryrun

migrate-rollback:
	sql-migrate down -config=local.dbconfig.yml -env="local-aztemarket"
	
up:
	docker compose up -d --remove-orphans --build

down:
	docker compose down -v

update-envs:
	openssl base64 -A -in .prod.env -out base64.prod.env.out
	openssl base64 -A -in .env -out base64.ci.env.out