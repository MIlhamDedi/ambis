pb:
	protoc -I protobuf kirito.proto --go_out=plugins=grpc:./kirito/pb
	protoc -I protobuf heathcliff.proto --go_out=plugins=grpc:./heathcliff/pb
	protoc -I protobuf sterben.proto --go_out=plugins=grpc:./sterben/pb
	protoc -I protobuf kirito.proto heathcliff.proto sterben.proto --go_out=plugins=grpc:./aincrad/pb