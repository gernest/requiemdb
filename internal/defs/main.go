package main

import (
	"bytes"
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()
	data, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	var b bytes.Buffer
	b.WriteString("export const rqDefinitions=`")
	b.Write(data)
	b.WriteByte('`')
	err = os.WriteFile(flag.Arg(1), b.Bytes(), 0600)
	if err != nil {
		log.Fatal(err)
	}
}
