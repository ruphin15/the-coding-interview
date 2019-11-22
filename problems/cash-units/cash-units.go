package main

import (
	"fmt"
	"math"

	"github.com/google/go-cmp/cmp"
)

func main() {
	cashUnits := getCashUnit(1972.78)

	fmt.Println(cmp.Equal(cashUnits[500], 3))
	fmt.Println(cmp.Equal(cashUnits[200], 2))
	fmt.Println(cmp.Equal(cashUnits[100], 0))
	fmt.Println(cmp.Equal(cashUnits[50], 1))
	fmt.Println(cmp.Equal(cashUnits[20], 1))
	fmt.Println(cmp.Equal(cashUnits[10], 0))
	fmt.Println(cmp.Equal(cashUnits[5], 0))
	fmt.Println(cmp.Equal(cashUnits[2], 1))
	fmt.Println(cmp.Equal(cashUnits[1], 0))
	fmt.Println(cmp.Equal(cashUnits[0.5], 1))
	fmt.Println(cmp.Equal(cashUnits[0.2], 1))
	fmt.Println(cmp.Equal(cashUnits[0.1], 0))
	fmt.Println(cmp.Equal(cashUnits[0.05], 1))
	fmt.Println(cmp.Equal(cashUnits[0.02], 1))
	fmt.Println(cmp.Equal(cashUnits[0.01], 1))
}

func getCashUnit(cash float64) map[float64]int {
	units := []float64{500.0, 200.0, 100.0, 50.0, 20.0, 10.0, 5.0, 2.0, 1.0, 0.5, 0.2, 0.1, 0.05, 0.02, 0.01}
	cashUnits := make(map[float64]int)

	for _, unit := range units {
		if cash < unit {
			continue
		}

		cashUnits[unit] = int(math.Floor(cash / unit))
		cash = math.Round((cash-(float64(cashUnits[unit])*unit))*100) / 100

		if cash == 0 {
			break
		}
	}

	return cashUnits
}
