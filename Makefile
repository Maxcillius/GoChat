migrate-up:
	migrate -path platforms/db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

migrate-down:
	migrate -path platforms/db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" down

db-proto:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative platforms/db/proto/db.proto