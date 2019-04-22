# sample code url
https://github.com/grpc/grpc-go/tree/master/examples/helloworld

# Run the application from outside Docker
docker exec server go run main.go

# Connect to container
docker exec -it server sh

# Run the application
go run main.go