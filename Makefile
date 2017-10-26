gen:
	protoc -I test_server/pb --go_out=plugins=grpc:pb test_server/pb/example.proto
