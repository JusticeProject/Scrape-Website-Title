package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io"

	"github.com/andybalholm/brotli"
)

//*************************************************************************************************
//*************************************************************************************************

func handleGzip(body []byte) bytes.Buffer {
	reader := bytes.NewReader(body)
	zipreader, err := gzip.NewReader(reader)
	defer zipreader.Close()
	if err != nil {
		DebugLog("Error in gzip:", err)
		return bytes.Buffer{}
	}

	writer := bytes.Buffer{}
	if _, err := io.Copy(&writer, zipreader); err != nil {
		DebugLog("Error in gzip io.Copy", err)
		return bytes.Buffer{}
	}

	return writer
}

//*************************************************************************************************
//*************************************************************************************************

func handleDeflate(body []byte) bytes.Buffer {
	reader := bytes.NewReader(body)
	zlibReader, err := zlib.NewReader(reader)
	defer zlibReader.Close()
	if err != nil {
		DebugLog("Error in deflate:", err)
		return bytes.Buffer{}
	}

	writer := bytes.Buffer{}
	_, err = io.Copy(&writer, zlibReader)
	if err != nil {
		DebugLog("Error in zlib io.Copy", err)
		return bytes.Buffer{}
	}

	return writer
}

//*************************************************************************************************
//*************************************************************************************************

func handleBrotli(body []byte) bytes.Buffer {
	reader := bytes.NewReader(body)
	brReader := brotli.NewReader(reader)
	// brotli.Reader does not have a Close function

	writer := bytes.Buffer{}
	_, err := io.Copy(&writer, brReader)
	if err != nil {
		DebugLog("Error in br io.Copy", err)
		return bytes.Buffer{}
	}

	return writer
}

//*************************************************************************************************
//*************************************************************************************************

func Decompress(decompressed bool, encoded bool, encoding string, body []byte) string {
	if !encoded || decompressed {
		return string(body)
	}

	DebugLog("found encoding:", encoding)

	if encoding == "gzip" {
		result := handleGzip(body)
		return result.String()
	} else if encoding == "deflate" || encoding == "flate" || encoding == "zlib" {
		result := handleDeflate(body)
		return result.String()
	} else if encoding == "br" {
		result := handleBrotli(body)
		return result.String()
	}

	DebugLog("Unknown encoding")
	return ""
}
