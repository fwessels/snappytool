package main

import (
	"github.com/golang/snappy"
	"io/ioutil"
	"testing"
	"log"
)

func benchEncode(b *testing.B, src []byte) {
	// Bandwidth is in amount of uncompressed data.
	b.SetBytes(int64(len(src)))
	dst := make([]byte, snappy.MaxEncodedLen(len(src)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		snappy.Encode(dst, src)
	}
}

func BenchmarkEncodeCT(b *testing.B) {
	src, err := ioutil.ReadFile("../../minio/perftest/dicomimport/data/CT.dcm")
	if err != nil {
		log.Fatal(err)
	}

	benchEncode(b, src)
}

func benchDecode(b *testing.B, src []byte) {
	encoded := snappy.Encode(nil, src)
	// Bandwidth is in amount of uncompressed data.
	b.SetBytes(int64(len(src)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		snappy.Decode(src, encoded)
	}
}

func BenchmarkDecodeCT(b *testing.B) {
	// Note: the file is OS-language dependent so the resulting values are not
	// directly comparable for non-US-English OS installations.
	src, err := ioutil.ReadFile("../../minio/perftest/dicomimport/data/CT.dcm")
	if err != nil {
		log.Fatal(err)
	}

	benchDecode(b, src)
}
