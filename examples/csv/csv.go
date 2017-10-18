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
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"github.com/intel/tfortools"
)

var code string

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-f template] file\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, tfortools.GenerateUsageUndecorated([][]string{}))
	}
	flag.StringVar(&code, "f", "{{table (totable .)}}", "string containing the template code to execute")
}

func applyTemplate(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Unable to open %s : %v", path, err)
	}
	defer func() {
		_ = f.Close()
	}()
	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return fmt.Errorf("Unable to read %s : %v", path, err)
	}
	return tfortools.OutputToTemplate(os.Stdout, "csv", code, data, nil)
}

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	if err := applyTemplate(flag.Args()[0]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
