# Connect to container
docker exec -it server sh

# Run the application
go run main.go

# Output xo model
xo --escape-all 'mysql://root:test@mysql/sample?parseTime=true&sql_mode=ansi' -o model