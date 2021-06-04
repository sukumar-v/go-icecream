package icecream

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// CheckBrackets will do this
func CheckBrackets(filePath string, lineNum int, callFuncName string) (lastLine int, args []string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lastLine++

		if lastLine == lineNum {
			var parentheses string

			parentheses = scanner.Text()
			var inlineArgsString string = strings.Split(parentheses, callFuncName)[1]

			inlineArgsString = strings.Replace(inlineArgsString, "(", "", -1)
			inlineArgsString = strings.Replace(inlineArgsString, ")", "", -1)

			var args []string = strings.Split(inlineArgsString, ",")

			return lastLine, args
		}
	}
	return -1, args
}
