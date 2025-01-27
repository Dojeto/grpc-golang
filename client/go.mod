module github.com/dojeto/grpc-demo-go/client

//For Local Development
replace github.com/dojeto/grpc-demo-go/proto => ../proto

go 1.22.4

require (
	github.com/dojeto/grpc-demo-go/proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.65.0
)

require (
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
