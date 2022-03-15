package main

import (
	"context"
	pb "crud/client/proto/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	routehttp := "/Find"
	client := pb.NewBookServiceClient(conn)

	if routehttp == "/Find" {
		req := &pb.BookId{Id: "23"}
		res, err := client.Find(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(res)
	} else if routehttp == "/Create" {
		req := &pb.Book{
			Name:        "Livro Teste 1",
			Description: "Um livro feito para teste '0'",
			MediumPrice: 10.99,
			Author:      "Carlos da Pamonha da Silva",
			ImgUrl:      "acarajé.com.br",
		}
		res, err := client.CreateBook(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(res)
	}

}
