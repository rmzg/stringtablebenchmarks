// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/rmzg/go-mph"

	"golang.org/x/text/unicode/rangetable"
	"golang.org/x/text/unicode/runenames"
)

func main() {
	fmt.Println("vim-go")

	//header, err := os.Open("header.go")
	//if err != nil {
	//fmt.Printf("Error: %s\n", err)
	//panic(1)
	//}

	//stat, _ := header.Stat()

	//data := make([]byte, stat.Size())

	//header.Read(data)

	file, err := os.Create("test.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		panic(1)
	}

	fmt.Fprintf(file, "package mphtest2\n\n import \"github.com/rmzg/go-mph\"\n")
	table := rangetable.Assigned("11.0.0")

	keys := make([]string, 0, 22000)
	runes := make([]rune, 0, 22000)
	dupeTest := make(map[string]bool)

	rangetable.Visit(table, func(r rune) {
		name := runenames.Name(r)
		if dupeTest[name] {
			return
		}
		dupeTest[name] = true

		keys = append(keys, name)
		runes = append(runes, r)
	})

	t := mph.New(keys)

	fmt.Fprintf(file, "var MyTable = %#v\nvar Names = %#v\n var Runes = %#v\n\n", t, keys, runes)

	fmt.Fprintf(file, `func KV(s string) rune { return Runes[MyTable.Query(s)] }`)

}
