package utils

import "math"

type CompoundInterestInterface interface {
	GetFutualValue(prevent int) float64
	GetPrevent(futualValue int) float64
	GetInterest() float64
	GetInterestType() InterestType
	GetFrequency() int
}

type InterestType int

const (
	Daily InterestType = iota + 1
	Monthy
	Annual
)

type CompoundInterest struct {
	interest     float64
	interestType InterestType
	frequency    int
}

func (ci CompoundInterest) GetInterest() float64 {
	return ci.interest
}

func (ci CompoundInterest) GetInterestType() InterestType {
	return ci.interestType
}

func (ci CompoundInterest) GetFrequency() int {
	return ci.frequency
}

func (ci CompoundInterest) GetFutualValue(prevent int) float64 {
	return float64(prevent) * math.Pow((1+ci.interest/100), float64(ci.frequency))
}

func (ci CompoundInterest) GetPrevent(futualValue int) float64 {
	return float64(futualValue) / math.Pow((1+ci.interest/100), float64(ci.frequency))
}
