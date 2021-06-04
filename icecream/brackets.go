package icecream

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	PARENTHESES = map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
		"<": ">",
	}
	OPEN_PARENTHESES  = "({["
	CLOSE_PARENTHESES = ")}]"
)

type element struct {
	value interface{}
	next  *element
}

type stack struct {
	top  *element
	size int
}

func (stack *stack) len() int {
	return stack.size
}

func (stack *stack) push(value interface{}) {
	stack.top = &element{value, stack.top}
	stack.size++
}

func (stack *stack) pop() (value interface{}) {
	if stack.size > 0 {
		value, stack.top = stack.top.value, stack.top.next
		stack.size--
		return
	}
	return nil
}

// CheckBrackets will do this
func CheckBrackets(filePath string, lineNum int) (lastLine int) {
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
			var isValid bool
			stack := new(stack)

			parentheses = scanner.Text()
			fmt.Println(parentheses)
			fmt.Println((strings.Split(parentheses, "getFuncCallFileInfo"))[1])
			if len(parentheses) > 0 {
				isValid = true
				for _, p := range strings.Split(parentheses, "getFuncCallFileInfo(") {
					if strings.Index(OPEN_PARENTHESES, p) != -1 {
						stack.push(p)
						continue
					}
					if strings.Index(CLOSE_PARENTHESES, p) != -1 {
						if stack.len() > 0 {
							lastParentheses := stack.pop().(string)
							if PARENTHESES[lastParentheses] == p {
								continue
							}
						}
						isValid = false
						break
					}
				}
			}

			if isValid && stack.len() == 0 {
				fmt.Println("True")
			} else {
				fmt.Println("False")
			}

		}
	}
	return -1
}
