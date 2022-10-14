// Use gRPC to send the diddocument between two nodes.
package main

// server

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"

	"fusion/node/sendDID/proto/pb/proto_demo"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// gRPC服务地址
	Address = "127.0.0.1:9988"
)

type helloService struct{}

var HelloService = helloService{}

func (h helloService) ProcessRequest(ctx context.Context, in *proto_demo.SignedRequest) (*proto_demo.SignedResponse, error) {
	resp := new(proto_demo.SignedResponse)
	requestBytes := "This is RequestBytes from server."
	signature := "This is Signature from server."
	resp.RequestBytes, resp.Signature = []byte(requestBytes), []byte(signature)
	return resp, nil
}

func (h helloService) HandleRequest_DID(ctx context.Context, in *proto_demo.SignedRequest) (*proto_demo.DID, error) {
	resDID := new(proto_demo.DID)
	scheme := "did:example:q7ckgxeq1lxmra0r"
	method := []string{
		"id: did:example:123#z6MkpzW2izkFjNwMBwwvKqmELaQcH8t54QL5xmBdJg9Xh1y4",
		"type: Ed25519VerificationKey2018",
		"controller: did:example:123",
		"publicKeyBase58: BYEz8kVpPqSt5T7DeGoPVUrcTZcDeX5jGkGhUQBWmoBg",
	}
	methodSpecificID := "controller: did:example:123"
	sc, _ := json.Marshal(scheme)
	me, _ := json.Marshal(method)
	meth, _ := json.Marshal(methodSpecificID)
	resDID.Scheme, resDID.Method, resDID.MethodSpecificID = sc, me, meth
	return resDID, nil
}

func main() {
	var user string
	flag.StringVar(&user, "u", "", "默认为启动服务端")
	flag.Parse()
	if user == "client" {
		conn, err := grpc.Dial(Address, grpc.WithInsecure())
		if err != nil {
			grpclog.Fatalln(err)
		}
		defer conn.Close()

		c := proto_demo.NewHandleClient(conn)

		req := &proto_demo.SignedRequest{
			RequestBytes: []byte("This is RequestBytes from client"),
			Signature:    []byte("This is Signature"),
		}
		res, err := c.HandleRequest_DID(context.Background(), req)
		if err != nil {
			grpclog.Fatalln(err)
		}

		fmt.Printf("%v\n", string(res.Scheme))
		fmt.Printf("%v\n", string(res.Method))
		fmt.Printf("%v\n", string(res.MethodSpecificID))
	} else if user == "server" {
		listen, err := net.Listen("tcp", Address)
		if err != nil {
			grpclog.Fatalf("Failed to listen: %v", err)
		}

		s := grpc.NewServer()

		proto_demo.RegisterHandleServer(s, HelloService)
		fmt.Println("Listen on " + Address)
		//grpclog.Println("Listen on " + Address)
		s.Serve(listen)
	} else {
		fmt.Println("Value of \"user\"(-u) should be \"client\" or \"server\"")
	}
}
