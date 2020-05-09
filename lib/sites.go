package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

//ShowLogs shows and print content in log.txt
func ShowLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error while open logs file")
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(string(file))
}

//WriteLog writes a log entry in log.txt
func WriteLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error readin file")
		fmt.Println(err)
		os.Exit(-1)
	}

	timestamp := time.Now().Format("02/01/2006 15:04:05")
	text := site + " - online: " + strconv.FormatBool(status) + "\n"

	file.WriteString(timestamp + " - " + text)

	file.Close()
}
