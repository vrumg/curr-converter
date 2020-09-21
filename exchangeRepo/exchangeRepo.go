package exchangeRepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"curr-converter/converter"
)

type Exchange struct {
	URL string
}

type exchangeBySrcCurrencyStruct struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Error string             `json:"error"`
}

func (s *Exchange) GetRates(srcCur string) (*converter.Rates, error) {

	res, err := http.Get(fmt.Sprintf(s.URL, srcCur))
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

	ret := &converter.Rates{Rates: resData.Rates, SrcCur: srcCur}

	return ret, nil
}
