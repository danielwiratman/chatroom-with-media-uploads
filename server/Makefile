migratecreate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrateup:
	migrate -database "postgres://postgres:admin@localhost:5432/chatroom?sslmode=disable" -path db/migrations up	
	
migratedown:
	migrate -database "postgres://postgres:admin@localhost:5432/chatroom?sslmode=disable" -path db/migrations down 
	
postgres:
	docker exec -it pg psql -U postgres
