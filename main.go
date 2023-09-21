package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	xpub := os.Getenv("XPUB")
	if xpub == "" {
		fmt.Println("XPUB environment variable not set")
		return
	}

	command := exec.Command("xpub", "derive", "-p", "p2wpkh", "-c", "0", "-i", "0", "-n", "3", xpub)

	cmd := strings.Join(command.Args, " ")
	fmt.Printf("cmd: %s\n", cmd)

	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("%s: %s\n", err.Error(), output)
		return
	}

	// fmt.Printf("output: %s\n", output)
	fmt.Printf("\n")

	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Add header row
	header := []string{"address"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	outputLines := strings.Split(string(output), "\n")

	// fmt.Printf("outputLines: %s\n", outputLines)

	for _, line := range outputLines {
		if line == "" {
			continue
		}
		fmt.Printf("addr: %s\n", line)
		err := writer.Write([]string{line})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	fmt.Printf("\n")

}
