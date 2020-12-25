```
docker build . -f Docker-protoc -t protoc:compile

docker run --rm -v $(pwd):/compile protoc:compile \
protoc --go_out=/compile/src/lib -I /compile /compile/protos/*.proto

```
