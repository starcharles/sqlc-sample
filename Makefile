migrate-create:
	dbmate new ${ARG} 
migrate-up:
	dbmate up

migrate-down:
	dbmate down

sqlc-generate:
	sqlc generate