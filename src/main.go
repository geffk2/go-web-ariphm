package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func handle(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	if request.Form["operation"] != nil && request.Form["first"] != nil && request.Form["second"] != nil {
		first, err1 := strconv.Atoi(request.Form["first"][0])
		second, err2 := strconv.Atoi(request.Form["second"][0])
		if err1 != nil || err2 != nil {
			return
		}
		result := float32(first)
		operation := rune(request.Form["operation"][0][0])
		if operation == '+' {		// +
			result += float32(second)
		}
		if operation == '-' {		// -
			result -= float32(second)
		}
		if operation == '*' {		// *
			result *= float32(second)
		}
		if operation == '/' {		// /
			result /= float32(second)
		}
		fmt.Printf("[LOG] %f %c %f = %f\n", float32(first), operation, float32(second), result)
		fmt.Fprintf(response, "%f\n", result)
	}
}

func main() {
	fmt.Println("Starting server at localhost:8080 ...")
	http.HandleFunc("/math", handle)
	http.ListenAndServe(":8080", nil)
}
