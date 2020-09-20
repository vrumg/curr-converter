package converter

import (
	"errors"
	"fmt"
)

type CurrencyRepo interface {
	GetRates(srcCur string) (*Rates, error)
}

type Rates struct {
	Rates  map[string]float64
	SrcCur string
}

type AmountConverter struct {
	Repo   CurrencyRepo
	rates  *Rates
	srcCur string
	dstCur string
	amount float64
}

func (s *AmountConverter) SetCurrencyAmount(amountIn float64, srcCurIn string, dstCurIn string) {
	s.amount = amountIn
	s.srcCur = srcCurIn
	s.dstCur = dstCurIn
}

func (s *AmountConverter) GetResult() (string, error) {
	var err error

	s.rates, err = s.Repo.GetRates(s.srcCur)
	if err != nil {
		return "", err
	}

	rate, err := s.getCurPair()
	if err != nil {
		return "", err
	}
	resString := fmt.Sprintf("%f %s %s", s.amount*rate, s.srcCur, s.dstCur)

	return resString, nil
}

func (s *AmountConverter) getCurPair() (float64, error) {

	rate, ok := s.rates.Rates[s.dstCur]
	if !ok {
		return 0, errors.New("Failed to find destination currency. Provide appropriate currency name")
	}

	return rate, nil
}
