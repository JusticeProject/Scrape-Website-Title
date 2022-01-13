package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var debugFileName string = ""

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: scrape-website-title.exe https://www.example.com [debugFilename.txt]")
		return
	}

	// check if we need to print debug statements
	if len(os.Args) == 3 {
		debugFileName = os.Args[2]
		DebugLog(os.Args)
	}

	// grab the url from the user
	url := os.Args[1]
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		DebugLog("Error in http.NewRequest:", err)
		fmt.Println(",1000")
		return
	}

	// add the custom HTTP headers
	// req.Header.Add("Connection", "keep-alive")
	// req.Header.Add(`Sec-Ch-Ua`, `" Not A;Brand";v="99", "Chromium";v="96", "Google Chrome";v="96"`)
	// req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	// req.Header.Add("Sec-Ch-Ua-Platform", "Windows")
	// req.Header.Add("Upgrade-Insecure-Requests", "1")
	// req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	// req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9`)
	// req.Header.Add("Sec-Fetch-Site", "none")
	// req.Header.Add("Sec-Fetch-Mode", "navigate")
	// req.Header.Add("Sec-Fetch-User", "?1")
	// req.Header.Add("Sec-Fetch-Dest", "document")
	// req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	// req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.1 Mobile/15E148 Safari/604.1")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		DebugLog("Error in client.Do:", err)
		fmt.Println(",1001")
		return
	}

	// check the HTTP status code from the server
	if resp.StatusCode > 299 {
		DebugLog("HTTP Status code:", resp.StatusCode)
	}

	// log the headers that were received from the server
	for key, value := range resp.Header {
		DebugLog(key, "=", value)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		DebugLog("Error in io.ReadAll:", err)
		fmt.Println(",1002")
		return
	}
	defer resp.Body.Close()

	// check which type of compression was used
	content_encoding, encoded := resp.Header["Content-Encoding"]
	var encoding string
	if encoded {
		encoding = content_encoding[0]
	}
	DebugSaveBinary(body)
	decoded := Decompress(resp.Uncompressed, encoded, encoding, body)
	DebugSaveHTML(&decoded)

	// extract title
	title := ExtractTitle(decoded)
	fmt.Printf("%s,%d\n", title, resp.StatusCode)
}
