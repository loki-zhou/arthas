#build proto windows 必须在cmd下， 其他类linux终端出错
docker run --rm -v  ${PWD}/../../:/defs namely/gen-grpc-gateway:1.29_4 -f schemas/detail/detail.proto -s Service -o productpage