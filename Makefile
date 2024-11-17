migrate-up:
	migrate -database "postgres://postgres:asutp93~~@localhost:5432/site?sslmode=disable" -path migrations up $(c)

migrate-down:
	migrate -database "postgres://postgres:asutp93~~@localhost:5432/site?sslmode=disable" -path migrations down 


migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

swag-init:
	swag init -g ./cmd/main/main.go