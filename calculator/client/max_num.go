package main

import (
	"context"
	"io"
	"log"

	pb "github.com/vova4o/grpc-go-study/calculator/proto"
)

// клиентская функция на много замороченее.
// поскольку у нас стрим в две стороны, нужно возвращать и получать данные стрима
// это делается в го рутинах.
func doMaxNum(c pb.MaxNumberServiceClient) {
	log.Println("doMaxNum was called")

	numbers := []int64{10, 30, 20, 50, 40, 70, 60, 90, 80, 100}

	// создаем стрим в обе стороны
	stream, err := c.GetMaxNumber(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GetMaxNumber: %v\n", err)
	}

	// первая го рутина отвечает за отправку чисел
	go func() {
		// в цикле передаем числа стриму
		for _, number := range numbers {
			log.Printf("Sending number: %v\n", number)
			// используя метод стрим передаем данные
			// Number назначен в прото файле!
			stream.Send(&pb.GetMaxNumberRequest{
				Number: number,
			})
		}
		// закрываем стрим
		stream.CloseSend()
	}()

	// вторая го рутина отвечает за получение чисел
	// делаем канал для ожидания окончания работы
	waitc := make(chan struct{})

	go func() {
		// бесконечный цикл для получения чисел
		for {
			// используя метод стрим получаем данные
			res, err := stream.Recv()
			// если мы получили конец файла прерываем цикл и завершаем го рутину
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while reading stream: %v\n", err)
			}

			// печатаем полученное число
			log.Printf("GetMaxNumber: %d\n", res.Result)
		}
		// когда цикл закончится, закрываем канал
		close(waitc)
	}()

	// ждем окончания работы
	<-waitc
}
