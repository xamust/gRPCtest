package main

import (
	"bufio"
	"client/pkg/api"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

//client gRPC for checking transmit data (для проверки)...

func main() {

	var myParam, myBook, myWriter string
	conn, err := grpc.Dial(":8080", grpc.WithInsecure()) //withinsecure, для корректного игнорирования
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	for {
		fmt.Println("Что(Кого) ищем? (book or writer, exit выход):")
		fmt.Scan(&myParam)
		if myParam == "book" {
			fmt.Println("Укажите имя книги:")
			fmt.Scan(&myBook)
			c := api.NewSearchingClient(conn)
			resp, err := c.Search(context.Background(), &api.SearchRequest{Book: myBook})
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Printf("Книгу %s написал %s\n", myBook, resp.GetWriter())

		} else if myParam == "writer" {
			fmt.Println("Укажите писателя:")
			//bufio для чтения строки с пробелами...
			myscanner := bufio.NewScanner(os.Stdin)
			myscanner.Scan()
			myWriter = myscanner.Text()
			c := api.NewSearchingClient(conn)
			resp, err := c.Search(context.Background(), &api.SearchRequest{Writer: myWriter})
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Printf("Писатель %s написал книги %s\n", myWriter, resp.GetBook())
		} else if myParam == "exit" {
			break
		}
	}

}
