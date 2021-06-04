package main

import (
    "fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/sukumar-varma/go-icecream/icecream"
)

const (
	icecreamPrefix string = "🍦|"
	// icecreamPrefix string = "ic|"
)

// IC prints in icecream format
func IC(str string) {
	var args []string

	_, file, line, ok1 := runtime.Caller(1)
	pc2, _, _, ok2 := runtime.Caller(0)

	if !ok1 || !ok2 {
		fmt.Println("An error has occured!")
	}

	var (
		filePath string = file
		lineNumber int = line
		callerName string = runtime.FuncForPC(pc2).Name()
	)

	_, args = icecream.CheckBrackets(filePath, lineNumber, strings.Split(callerName, ".")[1])

	for _, arg := range args {
		if arg != "" {
			fmt.Printf("%v %v: %q\n", icecreamPrefix, arg, str)
			fmt.Printf("  | type: %v\n", reflect.TypeOf(str))
		}
	}
}

func main() {
	var hello string = "hello World!"
	IC(hello)
	IC(hello)
}

