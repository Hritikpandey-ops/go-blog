# Variables
APP_NAME=blog-site

# Start server
run:
	go run server/cmd/main.go  # <-- THIS MUST BE A TAB, NOT SPACES

# Build binary
build:
	go build -o $(APP_NAME) main.go

# Start built binary
start: build
	./$(APP_NAME)

# Format code
fmt:
	go fmt ./...

# Clean binary
clean:
	rm -f $(APP_NAME)

# Run Docker containers (Postgres)
db-up:
	docker-compose up -d

# Stop Docker containers
db-down:
	docker-compose down

# Reset DB (Dangerous in prod)
db-reset:
	docker-compose down -v
	docker-compose up -d

# Tidy Go modules
tidy:
	go mod tidy

.PHONY: run build start fmt clean db-up db-down db-reset tidy


# make run	Runs your Go server (main.go)
# make db-up	Starts PostgreSQL using your docker-compose.yml
# make db-down	Stops PostgreSQL
# make db-reset	Stops PostgreSQL and deletes all volumes (fresh DB)
# make tidy	Runs go mod tidy
# make build	Builds the server binary
# make start	Builds and runs the binary
# make clean	Removes built binary

