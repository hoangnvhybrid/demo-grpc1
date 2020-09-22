.PHONY: compile
compile: ## Compile the proto file.
	protoc -I pb/ pb/helloworld.proto --go_out=plugins=grpc:pb/

.PHONY: server
server: ## Build and run server.
	go build -race -ldflags "-s -w" -o bin/server greeter_server/main.go
	bin/server

.PHONY: client
client: ## Build and run client.
	go build -race -ldflags "-s -w" -o bin/client greeter_client/main.go
	bin/client