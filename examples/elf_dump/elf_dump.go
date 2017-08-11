package main

import (
	"debug/elf"
	"flag"
	"fmt"
	"os"

	"github.com/intel/tfortools"
)

var code string //The format string

func init() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-f <template>] <filename> \n", os.Args[0])
		flag.PrintDefaults()

		fmt.Fprintf(os.Stderr, "The template passed to the -%s option operates on a\n\n", "f")
		fmt.Fprintln(os.Stderr, tfortools.GenerateUsageUndecorated(elf.File{}))

		fmt.Fprintln(os.Stderr, tfortools.TemplateFunctionHelp(nil))
	}

	flag.StringVar(&code, "f", "", "string containing the template code to execute")

}

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	filename := flag.Args()[0]
	if filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	ef, err := elf.NewFile(f)
	if err != nil {
		panic(err)
	}

	if code == "" {
		fmt.Printf("\nFile Header:\n")
		err = tfortools.OutputToTemplate(os.Stdout, "File", "{{tablealt (sliceof .FileHeader)}}", ef, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			panic(err)
		}

		fmt.Printf("\nProgram Headers:\n")
		err = tfortools.OutputToTemplate(os.Stdout, "Program headers", "{{tablealt (promote .Progs \"ProgHeader\")}}", ef, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			panic(err)
		}

		fmt.Printf("\nSection Headers:\n")
		err = tfortools.OutputToTemplate(os.Stdout, "Sections", "{{tablealt (promote .Sections \"SectionHeader\")}}", ef, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			panic(err)
		}

		fmt.Printf("\nSymbols:\n")
		s, _ := ef.Symbols()
		err = tfortools.OutputToTemplate(os.Stdout, "Symbols", "{{tablealt . }}", s, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			panic(err)
		}
		return
	}

	err = tfortools.OutputToTemplate(os.Stdout, "f", code, ef, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		panic(err)
	}
}
