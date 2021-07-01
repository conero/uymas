package syntax

import (
	"fmt"
	"testing"
)

func TestNumberOpera(t *testing.T) {
	var (
		eqal   string
		refRes float64
		res    float64
	)

	// Case 1
	//eqal = "21+7+2.5(1-0.3)+(1-0.5*0.21)*112-(1949-2021)/365"
	eqal = "21+7+2.5*(1-0.3)+(1-0.5*0.21)*112-(1949-2021)/365"
	refRes = 130.187260273973
	res = NumberOpera(eqal)
	if fmt.Sprintf("%.6f", res) != fmt.Sprintf("%.6f", refRes) {
		t.Fatalf("eqal [%v] -> %v != %v", eqal, res, refRes)
	}

	// Case 2
	eqal = "365*(1949-2021)"
	refRes = -26280
	res = NumberOpera(eqal)
	if fmt.Sprintf("%.6f", res) != fmt.Sprintf("%.6f", refRes) {
		t.Fatalf("eqal [%v] -> %v != %v", eqal, res, refRes)
	}
}
