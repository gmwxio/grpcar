package echosvc

import (
	"context"
	"fmt"
	"io"

	pb "github.com/wxio/grpcar/example/echo"
)

type echosvc struct {
	pb.UnimplementedEchoServer
}

// New EchoServer constructor
func New() pb.EchoServer {
	return &echosvc{}
}

func (svc *echosvc) UnaryEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Printf("unary echoing message %q\n", in.Message)
	return &pb.EchoResponse{Message: in.Message}, nil
}

func (svc *echosvc) ServerStreamingEcho(req *pb.EchoRequest, srv pb.Echo_ServerStreamingEchoServer) error {
	srv.Send(&pb.EchoResponse{Message: req.Message})
	return nil
}

func (svc *echosvc) ClientStreamingEcho(srv pb.Echo_ClientStreamingEchoServer) error {
	msg := ""
	for {
		in, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				srv.SendAndClose(&pb.EchoResponse{Message: msg})
				return nil
			}
			fmt.Printf("server: error receiving from stream: %v\n", err)
			return err
		}
		msg += in.Message
	}
}

func (svc *echosvc) BidirectionalStreamingEcho(stream pb.Echo_BidirectionalStreamingEchoServer) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			fmt.Printf("server: error receiving from stream: %v\n", err)
			return err
		}
		fmt.Printf("bidi echoing message %q\n", in.Message)
		stream.Send(&pb.EchoResponse{Message: in.Message})
	}
}
