.PHONY: deploy
deploy:
	gcloud app deploy ./app/api/app.yaml

.PHONY: protoc
protoc: protoc-grpc protoc-grpc-gateway protoc-grpc-gateway-swagger

.PHONY: protoc-grpc
protoc-grpc:
	docker-compose run --rm protoc \
		protoc \
		-I app/api/proto/ \
		-I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:app/api/proto \
		app/api/proto/api.proto

.PHONY: protoc-grpc-gateway
protoc-grpc-gateway:
	docker-compose run --rm protoc \
		protoc \
		-I app/api/proto/ \
		-I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:app/api/proto \
		app/api/proto/api.proto
	docker-compose run --rm protoc ./script/add-octrace-to-gw.sh

.PHONY: protoc-grpc-gateway-swagger
protoc-grpc-gateway-swagger:
	docker-compose run --rm protoc \
		protoc \
		-I app/api/proto/ \
		-I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:app/api/proto \
		app/api/proto/api.proto
