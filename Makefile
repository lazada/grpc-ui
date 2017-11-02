gen:
	protoc -I test_server/pb --go_out=plugins=grpc:test_server/pb test_server/pb/example.proto
gen-fixtures:
	protoc -I reflection/fixtures/simple --go_out=plugins=grpc:reflection/fixtures/simple reflection/fixtures/simple/simple.proto
build:
	npm install
	webpack
	gostatic2lib -package http_server -path static/ -out http_server/static.go
