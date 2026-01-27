# 生成 apprpc.proto 的 grpc 文件
goctl rpc protoc apprpc.proto --go_out=.  --go-grpc_out=.  --zrpc_out=.