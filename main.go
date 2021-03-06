package main

import (
	"fmt"
	"os"

	"golang.org/x/text/unicode/rangetable"
	"golang.org/x/text/unicode/runenames"
)

func main() {
	fmt.Println("vim-go")

	header, err := os.Open("header.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		panic(1)
	}

	stat, _ := header.Stat()

	data := make([]byte, stat.Size())

	header.Read(data)

	file, err := os.Create("test.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		panic(1)
	}

	file.Write(data)

	table := rangetable.Assigned("11.0.0")

	dupeTest := make(map[string]bool)

	file.Write([]byte(` var Runes = map[string]rune{`))

	rangetable.Visit(table, func(r rune) {
		name := runenames.Name(r)
		if dupeTest[name] {
			return
		}
		dupeTest[name] = true
		_ = name
		//fmt.Printf("%s = %c\n", name, r)
		//file.Write([]byte(fmt.Sprintf("case %q: return %q\n", name, r)))
		fmt.Fprintf(file, "%q:%q,\n", name, r)
	})

	fmt.Fprintf(file, "}")

	//file.Write([]byte(` func KV(s string) rune {
	//`))
	//file.Write([]byte("\nreturn 0}"))
}
