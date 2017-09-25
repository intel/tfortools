//
// Copyright (c) 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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

var fictionalStocks = []stock{
	{"BCOM.L", "Big Company", time.Now(), 120.23, 150.00, 119.00, 7500000},
	{"SMAL.L", "Small Company", time.Date(2017, time.March, 17, 10, 59, 00, 00, time.UTC), 1.06, 1.06, 1.10, 750},
	{"MEDI.L", "Medium Company", time.Date(2017, time.March, 17, 12, 23, 00, 00, time.UTC), 77.00, 75.11, 81.12, 300122},
	{"PICO.L", "Tiny Corp", time.Date(2017, time.March, 16, 16, 01, 00, 00, time.UTC), 0.59, 0.57, 0.63, 155},
	{"HENT.L", "Happy Enterprises", time.Date(2017, time.March, 17, 9, 45, 00, 00, time.UTC), 756.11, 600.00, 10000, 6395624278},
	{"LONL.L", "Lonely Systems", time.Date(2017, time.March, 17, 13, 45, 00, 00, time.UTC), 1245.00, 1200.00, 1245.00, 19003},
}

func main() {
	//START OMIT
const script = `{{ tojson . }}` // HL

err := tfortools.OutputToTemplate(os.Stdout, "tojson", script, fictionalStocks, nil)
	//END OMIT
	if err != nil {
		panic(err)
	}
}
