package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
//	pb "path/to/your/compiled/proto" // Замените на фактический путь к вашему скомпилированному proto-файлу
//	pb "animal" // Замените на путь к вашему файлу protobuf

	pb "github.com/AndreyMust/Illusis/app_go/animals/animal"
)

//type colorServer struct{
//  pb.UnimplementedColorServiceServer
//}

type hatServer struct{
  pb.UnimplementedHatServer
}
//func (s *colorServer) mustEmbedUnimplementedColorServiceServer() {}

//func (s *colorServer) GetColor(ctx context.Context, req *pb.Empty) (*pb.ColorResponse, error) {
//	return &pb.ColorResponse{Color: "your_color : RED"}, nil // Замените "your_color_here" на фактический цвет
//}

func (s *hatServer) GetAnimal(ctx context.Context, req *pb.Empty) (*pb.ColorAnimal, error) {


	fmt.Println("client was connect")

	return &pb.ColorAnimal{
		Animal: "Lion", // Замените на фактическое животное
		Color:  "Red",  // Замените на фактический цвет
	}, nil
}


func main() {

	fmt.Println("Server starting...")

	lis, err := net.Listen("tcp", ":50051") // Сервер будет слушать порт 50051
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	//pb.RegisterColorServiceServer(grpcServer, &colorServer{})
	pb.RegisterHatServer(grpcServer, &hatServer{})

	fmt.Println("Server is listening on port 50051...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}