.PHONY: protoc
protoc:
	docker-compose run --rm protoc protoc -I proto/ proto/api.proto --go_out=plugins=grpc:proto
