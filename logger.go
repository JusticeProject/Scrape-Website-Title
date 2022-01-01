package main

import (
	"bufio"
	"fmt"
	"os"
)

func DebugLog(v ...interface{}) {
	if len(debugFileName) > 0 {
		data := fmt.Sprintln(v...)
		data_slice := []byte(data)
		f, _ := os.OpenFile(debugFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		f.Write(data_slice)
		f.Close()
	}
}

// save the html data to a file
func DebugSaveHTML(html *string) {
	if len(debugFileName) > 0 {
		f, _ := os.Create("test.html")
		defer f.Close()
		w := bufio.NewWriter(f)
		w.WriteString(*html)
		w.Flush()
	}
}
