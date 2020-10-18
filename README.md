# gosaas
SaaS starter for Go and VueJS

## gRPC

```
# Setup and start
make cert
make run

# Call a method with params
grpcurl -d '{"name": "Vlad"}' -insecure localhost:9000 Greeter/SayHello
```