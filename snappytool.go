package main

import (
	"fmt"
	"github.com/golang/snappy"
	"io/ioutil"
	"log"
)

func main() {

	src, err := ioutil.ReadFile("/Users/frankw/golang/src/github.com/minio/perftest/dicomimport/data/CT.dcm")
	if err != nil {
		log.Fatal(err)
	}

	dst := make([]byte, snappy.MaxEncodedLen(len(src)))
	dst = snappy.Encode(dst, src)

	fmt.Printf("%v, %v\n", len(dst), float64(len(dst))/float64(len(src)))
}
