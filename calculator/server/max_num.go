package main

import (
	"io"
	"log"

	pb "github.com/vova4o/grpc-go-study/calculator/proto"
)

// GetMaxNumber получаем числа через стрим
// и через стрим их возвращаем, таким образом эта функция
// будет возвращать максимальное число каждый раз, когда оно появляется
func (s *Server) GetMaxNumber(stream pb.MaxNumberService_GetMaxNumberServer) error {
	log.Println("GetMaxNumber function was invoked")

	max := int64(0)
	// Бесконечный цикл
	for {
		// Читаем стрим по одному числу
		req, err := stream.Recv()
		// для того чтобы понять что данных больше нет,
		// необходимо использовать конец файла, таким образом программа
		// будет ждать пока не появится новое значение
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		// обраточка
		// полученое значение req.Number сравниваем с максимальным
		if req.Number > max {
			max = req.Number
			// отправляем максимальное значение
			sendErr := stream.Send(&pb.GetMaxNumberResponse{
				Result: max,
			})
			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v\n", sendErr)
			}
		}
	}
}
