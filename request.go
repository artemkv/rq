package main

import (
	"fmt"
	"net/http"
	"strings"
)

func makeRequest(req *requestData) (*http.Response, error) {
	if strings.EqualFold(req.Method, "get") {
		resp, err := Get(req.Url, req.Headers)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	if strings.EqualFold(req.Method, "post") {
		resp, err := Post(req.Url, req.Headers, []byte(req.Body))
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	if strings.EqualFold(req.Method, "put") {
		resp, err := Put(req.Url, req.Headers, []byte(req.Body))
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	if strings.EqualFold(req.Method, "delete") {
		resp, err := Delete(req.Url, req.Headers)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, fmt.Errorf("'%s' method is not supported", req.Method)
}

func parametrize(req requestData, inputs map[string]string) *requestData {
	var parametrized = &requestData{}

	parametrized.Method = req.Method
	parametrized.Url = substitute(req.Url, inputs)
	parametrized.Headers = make(map[string]string, len(req.Headers))
	for header, val := range req.Headers {
		parametrized.Headers[header] = substitute(val, inputs)
	}
	parametrized.Body = substitute(req.Body, inputs)

	return parametrized
}

func substitute(txt string, replacements map[string]string) string {
	for placeholder, val := range replacements {
		// TODO: work for GC, could it be optimized?
		txt = strings.ReplaceAll(txt, fmt.Sprintf("${%s}", placeholder), val)
	}
	return txt
}
