package controller

import (
	"errors"
	"strconv"

	"curr-converter/converter"
)

type InputData struct {
	Args   []string
	Domain *converter.AmountConverter
}

func (s *InputData) getValidationError() error {

	var err error

	if len(s.Args) != 3 {
		return errors.New("Failed to parse args: provide exact 3 parameters")
	}

	_, err = strconv.ParseFloat(s.Args[0], 64)
	if err != nil {
		return errors.New("Failed to parse amount: provide float64 type")
	}

	if len(s.Args[1]) != 3 {
		return errors.New("Failed to parse src curr: provide exact 3-letter currency name")
	}

	if len(s.Args[2]) != 3 {
		return errors.New("Failed to parse dst curr: provide exact 3-letter currency name")
	}

	return nil
}

func (s *InputData) ProccessInputData() (string, error) {

	err := s.getValidationError()
	if err != nil {
		return "", err
	}

	amount, _ := strconv.ParseFloat(s.Args[0], 64)
	s.Domain.SetCurrencyAmount(amount, s.Args[1], s.Args[2])
	ret, err := s.Domain.GetResult()
	if err != nil {
		return "", err
	}

	//call domain and get data
	//return data

	return ret, nil
}
