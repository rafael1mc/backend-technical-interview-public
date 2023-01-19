# !make
include .env

migration-create:
	migrate create -ext sql -dir migrations -seq $(name)

migration-up:
	migrate -path "./migrations" -database ${DATABASE_URL} up

migration-down:
	migrate -path "./migrations" -database ${DATABASE_URL} down 1