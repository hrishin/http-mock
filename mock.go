package httpmock

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

// Response is to mock the HTTP response
// to use with the MockClient
type MultiResponse []*Response

type Response struct {
	URI        string
	Body       string
	StatusCode int
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// Client create the mock http.client type
// to mock HTTP response with a given MockResponse type
func Client(response *Response) *http.Client {
	fn := func(req *http.Request) *http.Response {
		if strings.Contains(req.URL.String(), response.URI) {
			return &http.Response{
				StatusCode: response.StatusCode,
				Body:       ioutil.NopCloser(strings.NewReader(response.Body)),
				Header:     make(http.Header),
			}
		}

		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       ioutil.NopCloser(bytes.NewBufferString("Resource not found")),
			Header:     make(http.Header),
		}
	}

	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

// MultiResponseClient create the mock http.client type
// to mock HTTP response with a given MockResponse type
func MultiResponseClient(response MultiResponse) *http.Client {
	fn := func(req *http.Request) *http.Response {
		for _, r := range response {
			if strings.HasSuffix(req.URL.String(), r.URI) {
				return &http.Response{
					StatusCode: r.StatusCode,
					Body:       ioutil.NopCloser(strings.NewReader(r.Body)),
					Header:     make(http.Header),
				}
			}
		}

		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       ioutil.NopCloser(bytes.NewBufferString("Resource not found")),
			Header:     make(http.Header),
		}
	}

	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
