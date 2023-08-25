migrate:
	migrate -source file://postgres/migrations \
		-database postgres://bon:1234567890@localhost:5432/sqlc_todos up

rollback:
	migrate -source file://postgres/migrations \
		-database postgres://bon:1234567890@localhost:5432/sqlc_todos down
	
drop:
	migrate -source file://postgres/migrations \
		-database postgres://bon:1234567890@localhost:5432/sqlc_todos drop

migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir postgres/migrations $$name

sqlc:
	sqlc generate

dev:
	go run main.go

