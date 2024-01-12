migration_create:
	migrate create -ext=sql -dir=sql/migrations -seq initial

migrate_up:
	migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/orders" -verbose up

migrate_down:
	migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/orders" -verbose down

.PHONY: migration_create migrate_up migrate_down