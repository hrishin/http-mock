package httpmock

import (
	"io/ioutil"
	"testing"
)

func Test_mock_http_response(t *testing.T) {
	tests := []struct {
		title             string
		givenMock         *Response
		givenURI          string
		expectedStatus    int
		expectedtResponse string
	}{
		{
			title: "GET for a valid URI",
			givenMock: &Response{
				URI:        "/hello",
				StatusCode: 200,
				Body:       "hello test",
			},
			givenURI:          "/hello",
			expectedStatus:    200,
			expectedtResponse: "hello test",
		},
		{
			title: "GET for a invalid URI",
			givenMock: &Response{
				URI:        "/hello",
				StatusCode: 200,
				Body:       "hello test",
			},
			givenURI:          "/invalid",
			expectedStatus:    404,
			expectedtResponse: "Resource not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {

			client := Client(tc.givenMock)

			resp, err := client.Get("http://apiurl.com/" + tc.givenURI)
			if err != nil {
				t.Fatalf("Not expecting an error but error occurred executing http request : %v", err)
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Not expecting an error but error occurred reading the http response : %v", err)
			}

			if resp.StatusCode != tc.expectedStatus {
				t.Fatalf("Expecting %d response code but got %d response code", tc.expectedStatus, resp.StatusCode)
			}

			if string(body) != tc.expectedtResponse {
				t.Fatalf("Expecting %s response but got %s response", tc.expectedtResponse, body)
			}
		})
	}
}

func Test_multiple_http_response_mock(t *testing.T) {
	type expected struct {
		URI      string
		status   int
		response string
	}
	var tests = []struct {
		title     string
		givenMock MultiResponse
		expected  []*expected
	}{
		{
			title: "GET for a valid URI's",
			givenMock: MultiResponse{
				&Response{
					URI:        "/hello",
					StatusCode: 200,
					Body:       "hello test",
				},
				&Response{
					URI:        "/foo",
					StatusCode: 200,
					Body:       "foo test",
				},
			},
			expected: []*expected{
				{
					URI:      "/hello",
					status:   200,
					response: "hello test",
				},
				{
					URI:      "/foo",
					status:   200,
					response: "foo test",
				},
			},
		},
		{
			title: "GET for invalid URI's",
			givenMock: MultiResponse{
				&Response{
					URI:        "/hello",
					StatusCode: 200,
					Body:       "hello test",
				},
				&Response{
					URI:        "/foo",
					StatusCode: 200,
					Body:       "foo test",
				},
			},
			expected: []*expected{
				{
					URI:      "/hello",
					status:   200,
					response: "hello test",
				},
				{
					URI:      "/bar",
					status:   404,
					response: "Resource not found",
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {

			client := MultiResponseClient(tc.givenMock)

			for _, expected := range tc.expected {
				resp, err := client.Get("http://apiurl.com/" + expected.URI)
				if err != nil {
					t.Fatalf("Not expecting an error but error occurred executing http request : %v", err)
				}

				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Fatalf("Not expecting an error but error occurred reading the http response : %v", err)
				}

				if resp.StatusCode != expected.status {
					t.Fatalf("Expecting %d response code but got %d response code", expected.status, resp.StatusCode)
				}

				if string(body) != expected.response {
					t.Fatalf("Expecting %s response but got %s response", expected.response, body)
				}
			}
		})
	}
}
