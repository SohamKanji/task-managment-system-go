postgres:
	docker run --name task-managment-system-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine
createdb:
	docker exec -it task-managment-system-db createdb --username=root --owner=root task-managment-system-db 

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/task-managment-system-db?sslmode=disable" -verbose up

migratedown:	
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/task-managment-system-db?sslmode=disable" -verbose down

test: 
	go test -v -cover ./...