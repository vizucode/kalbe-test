package formatting

import (
	"math"

	"github.com/leekchan/accounting"
)

func ToIDR(num float64) (resp string) {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 0, Thousand: ","}
	flooredNum := math.Floor(float64(num))
	resp = ac.FormatMoney(flooredNum)
	return
}
