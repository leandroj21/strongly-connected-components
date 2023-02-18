package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// ReadFile reads a text file to extract the nodes list of it
func ReadFile(name string) (maxvalue int, nodeList []intTuple) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Some error just happened.")
		}
	}(file)
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var anod intTuple
	for scanner.Scan() {
		lineStr := scanner.Text()
		_, err2 := fmt.Sscanf(lineStr, "%d %d", &anod[0], &anod[1])
		if err2 != nil {
			return 0, nil
		}
		nodeList = append(nodeList, anod)
		if maxvalue < anod[0] {
			maxvalue = anod[0]
		}
		if maxvalue < anod[1] {
			maxvalue = anod[1]
		}
	}
	return
}

func abs(num int) int {
	if num < 0 {
		num = -1 * num
	}
	return num
}

// CountDigits counts the digits of an int number
func CountDigits(num int) (count int) {
	num = abs(num)
	for num > 0 {
		num = num / 10
		count += 1
	}
	return
}
