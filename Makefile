server:
	docker exec server go run main.go
test-server:
	docker exec server go run test/server.go
test-client:
	docker exec server go run test/client.go
