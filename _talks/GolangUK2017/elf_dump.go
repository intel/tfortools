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

import "os"

//START OMIT
import "debug/elf"
import "github.com/intel/tfortools"

func main() {
	if len(os.Args) != 3 {
		panic("Usage: elf_dump script file")
	}

	f, err := os.Open(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()
	ef, err := elf.NewFile(f)
	if err != nil {
		panic(err)
	}

	err = tfortools.OutputToTemplate(os.Stdout, "elf_dump", os.Args[1], ef, nil)
	if err != nil {
		panic(err)
	}
}

//END OMIT
