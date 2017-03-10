package httpfib

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/uudashr/fibweb"
)

type service struct {
	baseURL string
	client  *http.Client
}

// NewFibonacciService create HTTP based fibonacci service
func NewFibonacciService(baseURL string) fibweb.FibonacciService {
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 1 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 2 * time.Second,
	}

	client := &http.Client{
		Timeout:   1 * time.Second,
		Transport: transport,
	}

	return service{
		baseURL: baseURL,
		client:  client,
	}
}

func (svc service) Seq(limit int) ([]int, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/numbers", svc.baseURL), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("limit", strconv.Itoa(limit))
	req.URL.RawQuery = q.Encode()

	resp, err := svc.client.Do(req)
	if err != nil {
		return nil, err
	}

	var numbers []int
	err = json.NewDecoder(resp.Body).Decode(&numbers)
	if err != nil {
		return nil, err
	}
	return numbers, nil
}
