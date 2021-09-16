package main

import (
	"io"
	"net/http"
	"time"
)

func GetHTTP(url string, dst io.Writer) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err2 := io.Copy(dst, resp.Body)
	return err2
}
