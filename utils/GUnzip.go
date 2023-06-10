package utils

import "bytes"
import "compress/gzip"
import "io/ioutil"

func GUnzip(source []byte) []byte {

	var buffer = bytes.NewBuffer(source)

	reader, err1 := gzip.NewReader(buffer)

	if err1 == nil {

		decompressed, err2 := ioutil.ReadAll(reader)
		defer reader.Close()

		if err2 == nil {
			return decompressed
		}

	}

	return []byte{}

}
