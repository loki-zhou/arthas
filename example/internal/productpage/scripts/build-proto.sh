docker run -v ${PWD}/../../:/defs namely/gen-grpc-gateway:1.29_4 -f ../../schemas/detail/detail.proto -s Service -o ../

docker run --rm -v  ${PWD}/../../schemas:/defs namely/gen-grpc-gateway:1.29_4 -f detail/detail.proto -s Service -o ../