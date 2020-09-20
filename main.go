package main

import (
	"fmt"
	"log"
	"os"

	"./controller"
	"./converter"
	"./exchangeRepo"
)

func main() {
	repo := &exchangeRepo.Exchange{URL: "https://api.exchangeratesapi.io/latest?base=%s"}
	domain := &converter.AmountConverter{Repo: repo}
	ctrl := controller.InputData{Args: os.Args[1:], Domain: domain}

	result, err := ctrl.ProccessInputData()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(result)
}
