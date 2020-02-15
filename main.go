package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/alecthomas/mph"

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
	builder := mph.Builder()

	rangetable.Visit(table, func(r rune) {
		name := runenames.Name(r)
		if dupeTest[name] {
			return
		}
		dupeTest[name] = true
		_ = name
		//fmt.Printf("%s = %c\n", name, r)
		//file.Write([]byte(fmt.Sprintf("case %q: return %q\n", name, r)))
		builder.Add([]byte(name), []byte(string(r)))
	})

	chd, err := builder.Build()
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		panic(1)
	}

	buf := bytes.NewBuffer(nil)

	chd.Write(buf)

	fmt.Fprintf(file, "var data = %#v", buf.Bytes())

	//file.Write([]byte(` func KV(s string) rune {
	//`))
	//file.Write([]byte("\nreturn 0}"))
}
