build:
	go build -o echo_server echo_server.go
run: build 
	./echo_server
test:
	go run client/tcp_client.go