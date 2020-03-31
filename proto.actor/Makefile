docker-images:
	docker build -t server -f server.dockerfile .
	docker build -t client -f client.dockerfile .

create-network:
	docker network create actors

rm-network:
	docker network rm actors

start-server:
	docker run --rm --network=actors --name=server server

start-client:
	docker run --rm --network=actors --name=client client