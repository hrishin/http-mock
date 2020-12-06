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
