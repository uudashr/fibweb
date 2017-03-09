package fibweb

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

// FibonacciService provide all fibonacci operation
type FibonacciService interface {
	Seq(int) ([]int, error)
}

type httpFibonacciService struct {
	baseURL string
	client  *http.Client
}

// NewHTTPFibonacciService create HTTP based fibonacci service
func NewHTTPFibonacciService(baseURL string) FibonacciService {
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

	return httpFibonacciService{
		baseURL: baseURL,
		client:  client,
	}
}

func (svc httpFibonacciService) Seq(limit int) ([]int, error) {
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
