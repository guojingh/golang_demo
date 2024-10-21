package main

import (
	"bookstore_client/pb"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	//建立连接
	conn, err := grpc.Dial("127.0.0.1:8091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("did not connect: %v\n", err)
	}

	defer conn.Close()

	//创建客户端
	c := pb.NewBookstoreClient(conn)

	res, err := c.ListBooks(context.Background(), &pb.ListBooksRequest{Shelf: 3})
	if err != nil {
		fmt.Printf("could not list books: %v\n", err)
		return
	}

	fmt.Printf("next_page_token: %s\n", res.NextPageToken)
	for i, book := range res.Books {
		fmt.Printf("%d: %#v\n", i, book)
	}
}
