## Tesco Microservice

## Dependencies

 - go
 - docker

## Running Tests

 - Run all tests with `make docker_tests`
 - Before running `make docker_tests` set the $TESCO_API_KEY environment variable on your local system

## Manual Testing

  - Download [grpcc](https://github.com/njpatel/grpcc)
  - Start the server
  - Run `grpcc -i --proto grpc/tesco_products.proto --address localhost:3030`
  - Select `tescogrpc`
  - Enter `client.getProducts({barcode:'{Enter Barcode Here}'}, (e,r) => console.log(e, r))`


