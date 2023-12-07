package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func ReadBlocks(filePath string) (blocks [][]string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			blocks = append(blocks, lines)
			lines = []string{}
		} else {
			lines = append(lines, line)
		}
	}
	blocks = append(blocks, lines)

	return
}
