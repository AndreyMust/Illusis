package main

import (
	"context"
	"fmt"
	"net/http"
	"encoding/json"
//	"log"

	"google.golang.org/grpc"
	pb "animal" // Замените на путь к вашему файлу protobuf
)

type ResponseData struct {
	Animal string `json:"animal"`
	Color  string `json:"color"`
}

func handler(w http.ResponseWriter, r *http.Request) {
 	// строка сообщения сервера
	fmt.Fprintf(w, "This is Illisionist HTTP-server\n")

}

func handlerShow(w http.ResponseWriter, r *http.Request) {
 	// строка сообщения сервера
	//fmt.Fprintf(w, "Hello Andrew!! from HTTP-server Go!\n")

	//fmt.Println("Client was entered")

	  // Создаем экземпляр структуры
	data := ResponseData{
			Animal: "rabbit",
			Color:  "red",
        }
	data.Color = "white"


	// Кодируем структуру в JSON
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Отправляем JSON в ответ
	w.Write(response)

	getRPC()
}


func getRPC() {
        fmt.Println("start1")

        fmt.Println("start2")
        fmt.Println("start3")
	
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		//log.Fatalf("Failed to connect: %v", err)
		fmt.Printf("Failed to connect: %v \n", err)
		return
	}
	defer conn.Close()

	client := pb.NewHatClient(conn)
	if client == nil {
		//log.Fatalf("Error calling GetColor: %v", err)
		fmt.Printf("Error NewHatClient: %v \n", err)
		return
	}



	response, err := client.GetAnimal(context.Background(), &pb.Empty{})
	if err != nil {
		//log.Fatalf("Error calling GetColor: %v", err)
		fmt.Printf("Error calling: %v \n", err)
		return
	}

	fmt.Printf("Response from server: %s\n", response.Animal)
	fmt.Printf("Response from server: %s\n", response.Color)
}

func main() {
  fmt.Println("Server starting...")
  http.HandleFunc("/", handler)
  http.HandleFunc("/show", handlerShow)
  fmt.Println("Server was started :8080")
  http.ListenAndServe(":8080", nil)
}
