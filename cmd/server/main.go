package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc_calc_random_number_and_format/pb"
	"net"
)

type Server struct {
	pb.UnimplementedCalcServer
}

func (s *Server) ExecCalc(ctx context.Context, in *pb.CalcRequest) (*pb.CalcResponse, error) {
	println(in.SimboloMatematico)
	if in.GetSimboloMatematico() == "*" {
		return &pb.CalcResponse{Result: in.GetA() * in.GetB()}, nil
	}

	if in.GetSimboloMatematico() == "-" {
		return &pb.CalcResponse{Result: in.GetA() - in.GetB()}, nil
	}

	if in.GetSimboloMatematico() == "/" {
		return &pb.CalcResponse{Result: in.GetA() / in.GetB()}, nil
	}

	return &pb.CalcResponse{Result: in.A + in.B}, nil
}

func main() {
	println("Running gRPC server...")

	listener, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterCalcServer(server, &Server{})

	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
