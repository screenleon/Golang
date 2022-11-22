package utils

import (
	"math"
	"testing"
)

func TestGetFutureValue(t *testing.T) {
	prevent := 1000
	interest := 5.0
	frequency := 5
	var ci CompoundInterestInterface = CompoundInterest{interest: interest, interestType: Annual, frequency: frequency}
	futureValue, _ := math.Modf(ci.GetFutualValue(prevent))
	println(futureValue)
	if futureValue != float64(1276) {
		t.Errorf("Get wrong futualValue")
	}
}
