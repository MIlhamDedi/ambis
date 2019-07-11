pb:
	protoc -I protobuf kirito.proto --go_out=plugins=grpc:./kirito/pb
	protoc -I protobuf heathcliff.proto --go_out=plugins=grpc:./heathcliff/pb
	protoc -I protobuf sterben.proto --go_out=plugins=grpc:./sterben/pb
	protoc -I protobuf kirito.proto heathcliff.proto sterben.proto --go_out=plugins=grpc:./aincrad/pb
	protoc -I protobuf kirito.proto --js_out=import_style=commonjs:./yui/app/pb \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./yui/app/pb
	# web app servers handle end user api call until grpc-web is implemented
	protoc -I protobuf kirito.proto --go_out=plugins=grpc:./yui/pb

build-envoy:
	docker build -t helloworld/envoy -f ./alfheim/envoy.Dockerfile .

run-envoy:
	docker run -d -p 8080:8080 helloworld/envoy