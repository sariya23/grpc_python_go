package main

import (
	"context"
	"fmt"
	"log"
	"time"

	p "main/pkg/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	defer timer("main")()
	ctx := context.Background()
	addr := "localhost:9999"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := p.NewSolverClient(conn)
	chanRes := make(chan float64)
	go func() {
		resp, err := client.Solve(ctx, &p.SolverRequest{A: 2, B: 5, Sign: "+"})
		if err != nil {
			log.Fatal(err)
		}
		chanRes <- resp.Result
	}()
loop:
	for {
		select {
		case res := <-chanRes:
			fmt.Printf("res is %v", res)
			break loop
		default:
			time.Sleep(time.Second)
			fmt.Println("waiting")
		}
	}
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
