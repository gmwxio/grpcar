module github.com/wxio/grpcar

go 1.13

replace github.com/fullstorydev/grpcurl => /home/garym/devel/github.com/fullstorydev/grpcurl

replace google.golang.org/grpc => /home/garym/devel/github.com/grpc/grpc-go

require (
	github.com/fullstorydev/grpcurl v1.4.0
	github.com/golang/protobuf v1.3.2
	github.com/jhump/protoreflect v1.5.0
	github.com/stretchr/testify v1.4.0
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v1.24.0
)
