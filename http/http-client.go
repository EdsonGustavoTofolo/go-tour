package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func RunHttpClientGet() {
	c := http.Client{
		Timeout: time.Duration(1) * time.Second,
	}

	resp, err := c.Get("http://google.com")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))
}

func RunHttpClientPost() {
	c := http.Client{}

	json := bytes.NewBuffer([]byte(`{"name":"Edson"}`))

	resp, err := c.Post("http://google.com", "application/json", json)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}

func RunHttpClientWithRequestCreationCustomized() {
	req, err := http.NewRequest("GET", "http://google.com", nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")

	c := http.Client{}

	resp, err := c.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}

func RunHttpClientWithRequestCreationCustomizedAndContext() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Microsecond)
	//ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
