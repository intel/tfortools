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
