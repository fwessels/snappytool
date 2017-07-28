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

func benchmarkEncode(b *testing.B, filename string) {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	benchEncode(b, src)
}

func BenchmarkEncodeCT(b *testing.B) { benchmarkEncode(b,"../../minio/perftest/dicomimport/data/CT.dcm") }
//func BenchmarkEncodeCT2(b *testing.B) { benchmarkEncode(b,"../../minio/perftest/dicomimport/data/CT-2.dcm") }
func BenchmarkEncodeMR(b *testing.B) { benchmarkEncode(b,"../../minio/perftest/dicomimport/data/MR.dcm") }

func benchDecode(b *testing.B, src []byte) {
	encoded := snappy.Encode(nil, src)
	// Bandwidth is in amount of uncompressed data.
	b.SetBytes(int64(len(src)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		snappy.Decode(src, encoded)
	}
}

func benchmarkDecode(b *testing.B, filename string) {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	benchDecode(b, src)
}

func BenchmarkDecodeCT(b *testing.B) { benchmarkEncode(b,"../../minio/perftest/dicomimport/data/CT.dcm") }
//func BenchmarkDecodeCT2(b *testing.B) { benchmarkEncode(b,"../../minio/perftest/dicomimport/data/CT-2.dcm") }
func BenchmarkDecodeMR(b *testing.B) { benchmarkEncode(b,"../../minio/perftest/dicomimport/data/MR.dcm") }
