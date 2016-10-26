package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/groob/plist"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fileType := check(r)
	switch fileType {
	case "plist":
		plistToJSON(r)
	case "json":
		jsonToPLIST(r)
	default:
		fmt.Println("unknown filetype, must be json or plist")
		os.Exit(1)
	}
}

func check(r *bufio.Reader) string {
	data, err := r.Peek(5)
	if err != nil {
		panic(err)
	}

	if bytes.Contains(data, []byte("<")) {
		return "plist"
	}
	return "json"
}

func plistToJSON(r io.Reader) {
	var into interface{}
	if err := plist.NewDecoder(r).Decode(&into); err != nil {
		panic(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(into); err != nil {
		panic(err)
	}
}

func jsonToPLIST(r io.Reader) {
	var into interface{}
	if err := json.NewDecoder(r).Decode(&into); err != nil {
		panic(err)
	}

	enc := plist.NewEncoder(os.Stdout)
	enc.Indent("  ")
	if err := enc.Encode(into); err != nil {
		panic(err)
	}
}
