.PHONY: deploy
deploy:
	gcloud app deploy ./app/api/app.yaml

.PHONY: protoc
protoc: protoc-grpc protoc-grpc-gateway

.PHONY: protoc-grpc
protoc-grpc:
	docker-compose run --rm protoc \
		protoc \
		-I proto/ \
		-I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:proto \
		proto/api.proto

.PHONY: protoc-grpc-gateway
protoc-grpc-gateway:
	docker-compose run --rm protoc \
		protoc \
		-I proto/ \
		-I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:proto \
		proto/api.proto
