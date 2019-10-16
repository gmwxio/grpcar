package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

var (
	exit = os.Exit
)

func main() {

}

type grpcproxy struct {
	ReflHeaders    []string `help:"Additional reflection headers in 'name: value' format. May specify more than one via multiple flags. These headers will only be used during reflection requests and will be excluded when invoking the requested RPC method."`
	ConnectTimeout float64  `help:"The maximum time, in seconds, to wait for connection to be established. Defaults to 10 seconds."`
	Target         string   `opts:"mode=arg" help:"target grpc servers"`
	PlainText      bool
	Insecure       bool
	Cacert         string
	Cert           string
	Key            string
	ServerName     string
	KeepaliveTime  float64
	MaxMsgSz       int `help:"The maximum encoded size of a message that grpcui will accept. If not specified, defaults to 4mb."`
}

func fail(err error, msg string, args ...interface{}) {
	if err != nil {
		msg += ": %v"
		args = append(args, err)
	}
	fmt.Fprintf(os.Stderr, msg, args...)
	fmt.Fprintln(os.Stderr)
	if err != nil {
		exit(1)
	} else {
		// nil error means it was CLI usage issue
		fmt.Fprintf(os.Stderr, "Try '%s -help' for more details.\n", os.Args[0])
		exit(2)
	}
}

func (gp *grpcproxy) Run() error {
	ctx := context.Background()
	dialTime := 10 * time.Second
	if gp.ConnectTimeout > 0 {
		dialTime = time.Duration(gp.ConnectTimeout * float64(time.Second))
	}
	ctx, cancel := context.WithTimeout(ctx, dialTime)
	defer cancel()
	//
	var opts []grpc.DialOption
	if gp.KeepaliveTime > 0 {
		timeout := time.Duration(gp.KeepaliveTime * float64(time.Second))
		opts = append(opts, grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    timeout,
			Timeout: timeout,
		}))
	}
	if gp.MaxMsgSz > 0 {
		opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(gp.MaxMsgSz)))
	}
	// if *authority != "" {
	// 	opts = append(opts, grpc.WithAuthority(*authority))
	// }
	//
	var creds credentials.TransportCredentials
	if !gp.PlainText {
		var err error
		creds, err = grpcurl.ClientTransportCredentials(gp.Insecure, gp.Cacert, gp.Cert, gp.Key)
		if err != nil {
			fail(err, "Failed to configure transport credentials")
		}
		if gp.ServerName != "" {
			if err := creds.OverrideServerName(gp.ServerName); err != nil {
				fail(err, "Failed to override server name as %q", gp.ServerName)
			}
		}
	}
	//
	network := "tcp"
	// if isUnixSocket != nil && isUnixSocket() {
	// 	network = "unix"
	// }
	cc, err := grpcurl.BlockingDial(ctx, network, gp.Target, creds, opts...)
	if err != nil {
		fail(err, "Failed to dial target host %q", gp.Target)
	}
	var descSource grpcurl.DescriptorSource
	var refClient *grpcreflect.Client
	md := grpcurl.MetadataFromHeaders(gp.ReflHeaders)
	refCtx := metadata.NewOutgoingContext(ctx, md)
	refClient = grpcreflect.NewClient(refCtx, reflectpb.NewServerReflectionClient(cc))
	descSource = grpcurl.DescriptorSourceFromServer(ctx, refClient)
	// arrange for the RPCs to be cleanly shutdown
	reset := func() {
		if refClient != nil {
			refClient.Reset()
			refClient = nil
		}
		if cc != nil {
			cc.Close()
			cc = nil
		}
	}
	defer reset()
	exit = func(code int) {
		// since defers aren't run by os.Exit...
		reset()
		os.Exit(code)
	}
	//
	methods, err := getMethods(descSource)
	if err != nil {
		fail(err, "Failed to compute set of methods to expose")
	}
	allFiles, err := grpcurl.GetAllFiles(descSource)
	if err != nil {
		fail(err, "Failed to enumerate all proto files")
	}
	// can go ahead and close reflection client now
	if refClient != nil {
		refClient.Reset()
		refClient = nil
	}

	return nil
}

type proxy struct {
	Source grpcurl.DescriptorSource
	ch     grpcdynamic.Channel
}

func (pr *proxy) unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	fmt.Printf("ctx:%+v\nreq:%+v\ninfo:%+v\nhandler:%+v\n", ctx, req, info, handler)
	// authentication (token verification)
	md, ok := metadata.FromIncomingContext(ctx)
	var headers []string
	for k, v := range md {
		for _, v2 := range v {
			headers = append(headers, k+": "+v2)
		}
	}

	var iehandler grpcurl.InvocationEventHandler
	var requestData grpcurl.RequestSupplier

	grpcurl.InvokeRPC(
		ctx,
		pr.Source,
		pr.ch,
		info.FullMethod,
		headers,
		iehandler,
		requestData,
	)
	// if !valid(md["authorization"]) {
	// 	return nil, errInvalidToken
	// }
	// m, err := handler(ctx, req)
	// if err != nil {
	// 	logger("RPC failed with error %v", err)
	// }
	return m, err
}

func getMethods(source grpcurl.DescriptorSource) ([]*desc.MethodDescriptor, error) {
	allServices, err := source.ListServices()
	if err != nil {
		return nil, err
	}
	var descs []*desc.MethodDescriptor
	for _, svc := range allServices {
		if svc == "grpc.reflection.v1alpha.ServerReflection" {
			continue
		}
		d, err := source.FindSymbol(svc)
		if err != nil {
			return nil, err
		}
		sd, ok := d.(*desc.ServiceDescriptor)
		if !ok {
			return nil, fmt.Errorf("%s should be a service descriptor but instead is a %T", d.GetFullyQualifiedName(), d)
		}
		for _, md := range sd.GetMethods() {
			descs = append(descs, md)
		}
	}
	return descs, nil
}
