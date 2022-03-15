package grpc_svc

import (
	"context"
	"crud/internal/database"
	"crud/internal/models"
	"crud/proto/pb"
	"errors"
	"strconv"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedBookServiceServer
}

func (service Server) CreateBook(ctx context.Context, req *pb.Book) (*pb.BookResponse, error) {
	db := database.GetDatabase()
	var book models.Book
	book.Name = req.Name
	book.Author = req.Author
	book.ImageURL = req.ImgUrl
	book.Description = req.Description
	book.MediumPrice = req.MediumPrice

	//Criando
	err := db.Create(&book).Error
	if err != nil {
		return &pb.BookResponse{Status: 500}, errors.New("Erro ao criar o livro")
	}

	response := &pb.BookResponse{
		Status: 200,
	}
	return response, nil
}

func (Server) Find(ctx context.Context, req *pb.BookId) (*pb.Book, error) {
	db := database.GetDatabase()

	var book models.Book
	err := db.First(&book, req.Id).Error
	if err != nil {
		return &pb.Book{}, errors.New("Livro não encontrado")
	}

	response := &pb.Book{
		Id:          strconv.FormatUint(uint64(book.ID), 10),
		Name:        book.Name,
		Description: book.Description,
		MediumPrice: book.MediumPrice,
		Author:      book.Author,
		ImgUrl:      book.ImageURL,
	}
	return response, nil
}

func RegisterService(grpc_server grpc.Server) {
	pb.RegisterBookServiceServer(&grpc_server, &Server{})
}
