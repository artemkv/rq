package main

import (
	"fmt"
	"net/http"
)

func printRequest(req *requestData) {
	fmt.Printf("%s %s\n", req.Method, req.Url)
	for header, headerValue := range req.Headers {
		fmt.Printf("%s: %s\n", header, headerValue)
	}
	fmt.Println("")
	fmt.Printf("%s\n", req.Body)
}

func printResponse(resp *http.Response, body []byte) error {
	fmt.Printf("%v\n", resp.Status)
	for header := range resp.Header {
		for _, headerValue := range resp.Header[header] {
			fmt.Printf("%s: %s\n", header, headerValue)
		}
	}
	fmt.Println("")
	fmt.Printf("%s\n", string(body[:]))
	return nil
}

func verbosePrintRequest(requestId string, request *requestData) {
	fmt.Println("==========================")
	fmt.Printf("Request: %s\n", requestId)
	fmt.Println("==========================")
	printRequest(request)
	fmt.Println("--------------------------")
}

func verbosePrintResponse(resp *http.Response, body []byte) error {
	err := printResponse(resp, body)
	if err != nil {
		return err
	}
	fmt.Println("==========================")
	return nil
}
