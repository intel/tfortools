package main

import (
	"debug/macho"
	"os"

	"github.com/intel/tfortools"
) //START OMIT

func main() {
	if len(os.Args) != 3 {
		panic("Usage: elf_dump script file")
	}

	f, err := os.Open(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()
	ef, err := macho.NewFile(f)
	if err != nil {
		panic(err)
	}

	err = tfortools.OutputToTemplate(os.Stdout, "elf_dump", os.Args[1], ef, nil)
	if err != nil {
		panic(err)
	}
}

//END OMIT
