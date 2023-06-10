package utils

import "bytes"
import "compress/gzip"

func GZip(source []byte) []byte {

	var buffer bytes.Buffer

	writer := gzip.NewWriter(&buffer)
	_, err := writer.Write(source)
	defer writer.Close()

	if err == nil {
		return buffer.Bytes()
	}

	return []byte{}

}
