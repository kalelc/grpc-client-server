build-server:
	docker build --platform linux/amd64 -f DockerfileServer -t grpc-example-server .
