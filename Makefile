gen:
	protoc -I test_server/pb --go_out=plugins=grpc:test_server/pb test_server/pb/example.proto
gen-fixtures:
	protoc -I reflection/fixtures/simple --go_out=plugins=grpc:reflection/fixtures/simple reflection/fixtures/simple/simple.proto
