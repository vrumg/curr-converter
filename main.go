package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

const exchangeByBaseCur string = "https://api.exchangeratesapi.io/latest?base=%s"

type exchangeBySrcCurrencyStruct struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Error string             `json:"error"`
}

func main() {
	//read args
	args := os.Args[1:]
	if len(args) != 3 {
		log.Fatal("Failed to parse args: need exact 3 parameters")
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		log.Fatal("Failed to parse amount: non-float64 type")
	}

	srcCur := args[1]
	if len(srcCur) != 3 {
		log.Fatal("Failed to parse src curr: need exact 3-letter currency name")
	}

	dstCur := args[2]
	if len(dstCur) != 3 {
		log.Fatal("Failed to parse dst curr: need exact 3-letter currency name")
	}

	res, err := GetExchangeString(amount, srcCur, dstCur)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res)

}

func GetExchangeString(amount float64, srcCur string, dstCur string) (string, error) {

	rates, err := GetRateListBySrcCur(srcCur)
	if err != nil {
		return "", err
	}

	rate, err := GetCurPair(rates, dstCur)
	if err != nil {
		return "", err
	}

	resString := fmt.Sprintf("%f %s %s", amount*rate, srcCur, dstCur)

	return resString, nil

}

func GetRateListBySrcCur(srcCur string) (map[string]float64, error) {

	res, err := http.Get(fmt.Sprintf(exchangeByBaseCur, srcCur))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get exchange rate with error: %s", err.Error()))
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to read exchange rate response with error: %s", err.Error()))
	}
	res.Body.Close()

	resData := &exchangeBySrcCurrencyStruct{}
	err = json.Unmarshal(data, resData)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to unmarshal exchange rate response with error: %s", err.Error()))
	}
	if resData.Error != "" {
		return nil, errors.New(fmt.Sprintf("Failed to unmarshal exchange rate response with error: %s", resData.Error))
	}

	return resData.Rates, nil
}

func GetCurPair(rates map[string]float64, dstCur string) (float64, error) {

	rate, ok := rates[dstCur]
	if !ok {
		return 0, errors.New("Failed to find destination currency")
	}

	return rate, nil
}
