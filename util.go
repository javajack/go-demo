package main

import (
	"io/ioutil"
	"net/http"
)

// ProxyCall fetches google ersults via http client
func ProxyCall() []byte {
	resp, _ := http.Get("https://www.google.com?q=AdZapier")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
