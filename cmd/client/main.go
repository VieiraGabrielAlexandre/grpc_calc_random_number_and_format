package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_calc_random_number_and_format/pb"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	client := pb.NewCalcClient(conn)

	a := int64(generateNumber())
	b := int64(generateNumber())
	symbol := generateSymbol()

	req := &pb.CalcRequest{
		A:                 a,
		B:                 b,
		SimboloMatematico: symbol,
	}

	println(symbol)

	res, err := client.ExecCalc(context.Background(), req)

	if err != nil {
		panic(err)
	}

	println(res.GetResult())
}

func generateNumber() int {
	return rand.Intn(rand.Intn(time.Now().Second()))
}

func generateSymbol() string {
	symbol := "*-/+"
	index := rand.Intn(len(symbol))

	return strconv.Itoa(int(symbol[index]))
}
