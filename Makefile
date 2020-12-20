# kartoffel-notification/Makefile
build:
	docker build -t \
	kartoffel-notification \
	.
	
run-nats:
	docker run -d --rm \
	--name nats \
	-p 4222:4222 \
	nats

run-bg:
	docker run -d --rm \
	--name kartoffel-notification \
	-p 30010 \
	-e MICRO_SERVER_ADDRESS=:30010 \
	-e MICRO_REGISTRY=mdns \
	kartoffel-notification:1.0

run-fg:
	docker run --rm \
	--name kartoffel-notification \
	-p 30010 \
	-e MICRO_SERVER_ADDRESS=:30010 \
	-e MICRO_REGISTRY=mdns \
	kartoffel-notification:1.0