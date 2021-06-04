package main

import (
	"bufio"
	"io"
    "fmt"
    "log"
	"os"
	"reflect"
	"runtime"

	"github.com/sukumar-varma/go-icecream/icecream"
)

func getFuncCallFileInfo(str string) (string, string, int, string) {
	pc1, file, line, ok1 := runtime.Caller(1)
	pc2, _, _, ok2 := runtime.Caller(0)
	if ok1 && ok2 {
		return runtime.FuncForPC(pc1).Name(), file, line, runtime.FuncForPC(pc2).Name()
	}
	return "", "", -1, ""
}

func getLine(r io.Reader, lineNum int) (line string, lastLine int, err error) {
    sc := bufio.NewScanner(r)
    for sc.Scan() {
        lastLine++
        if lastLine == lineNum {
            // you can return sc.Bytes() if you need output in []bytes
            return sc.Text(), lastLine, sc.Err()
        }
	}
	if err := sc.Err(); err != nil {
        log.Fatal(err)
    }

    return line, lastLine, io.EOF
}

func main() {
	var str string = "hello"
	funcName, filePath, lineNumber, callerName := getFuncCallFileInfo(str,)
	fmt.Println(funcName, filePath, lineNumber, callerName)

	file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()

	// fmt.Println(getLine(file, lineNumber))
	icecream.CheckBrackets(filePath, lineNumber)

	scanner := bufio.NewScanner(file)
	fmt.Println(reflect.TypeOf(scanner))
}

