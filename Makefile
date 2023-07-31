migrate-create:
	dbmate new ${ARG} 
migrate-up:
	dbmate up

migrate-down:
	dbmate down

sqlc-generate:
	sqlc generate

start-local:
	DATABASE="user:password@tcp(127.0.0.1:3308)/sample_db" go run main.go