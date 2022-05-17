cd $( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )/../api
protoc -I. --go_out="." --go-grpc_out="." "api.proto"