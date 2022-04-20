gen:
	# 生成proto目录下的所有protobuf文件，插件使用grpc，生成的文件存放在当前目录下（具体存放路径在.proto文件中通过 go_package指定）
	protoc --proto_path proto proto/*.proto --plugin=grpc --go_out=. --go-grpc_out=.

clean:
	rm pb/*.go

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

test:
	go test -cover -race ./...

.PHONY: client