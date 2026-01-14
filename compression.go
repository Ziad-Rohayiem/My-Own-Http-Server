package main

import (
	"bytes"
	"compress/gzip"
)

// CompressGzip compresses data using gzip
func CompressGzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)

	_, err := gzipWriter.Write(data)
	if err != nil {
		gzipWriter.Close()
		return nil, err
	}

	err = gzipWriter.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// SupportsGzip checks if the client supports gzip encoding
func SupportsGzip(schemes []string) bool {
	for _, scheme := range schemes {
		if scheme == "gzip" {
			return true
		}
	}
	return false
}
