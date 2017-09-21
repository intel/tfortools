package main

import (
	"os"
	"time"

	"github.com/intel/tfortools"
)

type stock struct {
	Ticker    string
	Name      string
	LastTrade time.Time
	Current   float64
	High      float64
	Low       float64
	Volume    int
}

//START OMIT
var fictionalStocks = []interface{}{ // HL
	stock{"BCOM.L", "Big Company", time.Now(), 120.23, 150.00, 119.00, 7500000},
}

func main() {
	const script = `{{ describe . }}` // HL
	err := tfortools.OutputToTemplate(os.Stdout, "filter", script, fictionalStocks, nil)
	//END OMIT
	if err != nil {
		panic(err)
	}
}
