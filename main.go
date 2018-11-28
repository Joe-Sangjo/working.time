package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	filePath string
)

func init() {
	flag.StringVar(&filePath, "f", "", "working time file")
	flag.Parse()
}

func main() {

	timeData, err := os.Open(filePath)
	defer timeData.Close()
	if err != nil {
		fmt.Println("file is not exist\n")
		return
	}
	inp := bufio.NewScanner(timeData)
	layout := "15:04:05"

	totalWorkingTime, _ := time.Parse(layout, "00:00:00")
	for inp.Scan() {
		timeData := strings.Split(inp.Text(), ":")
		hour, _ := time.ParseDuration(timeData[0] + "h")
		minitue, _ := time.ParseDuration(timeData[1] + "m")

		totalWorkingTime = totalWorkingTime.Add(hour).Add(minitue)
	}

	fmt.Println((totalWorkingTime.Day()-1)*24+totalWorkingTime.Hour(), ":", totalWorkingTime.Minute())
}
