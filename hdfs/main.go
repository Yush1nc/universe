package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/colinmarc/hdfs"
)

const (
	NameNode = ""
	User     = ""
	Filename = ""
	Output   = ""
)

func main() {
	client, err := hdfs.NewClient(hdfs.ClientOptions{
		Addresses: []string{NameNode},
		User:      User,
	})
	if err != nil {
		panic(err)
	}

	reader, err := client.Open(Filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = reader.Close()
	}()

	if Output != "" {
		f, err := os.Create(Output)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err := io.Copy(f, reader); err != nil {
			panic(err)
		}
	} else {
		buffer := bytes.NewBuffer(nil)
		if _, err = io.Copy(buffer, reader); err != nil {
			panic(err)
		}
		fmt.Println(buffer.String())
	}
}
