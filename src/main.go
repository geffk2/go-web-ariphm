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
		operation := int32(request.Form["operation"][0][0])
		if operation == 32 {
			result += float32(second)
		}
		if operation == 45 {
			result -= float32(second)
		}
		if operation == 42 {
			result *= float32(second)
		}
		if operation == 47 {
			result /= float32(second)
		}
		fmt.Printf("%c %f\n", operation, result)
		fmt.Fprintf(response, "%f\n", result)
	}
}

func main() {
	fmt.Println("Starting server at localhost:8080 ...")
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}
