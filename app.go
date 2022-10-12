package main

import (
	"fmt"
	"net/http"
	delivery "play3/internal/delivery"
	repository "play3/internal/repository"
	usecase "play3/internal/usecase"
)

func main() {
	delivery.InitDelivery(usecase.InitUsecase(repository.InitRepository()))
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
