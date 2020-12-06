# HTTP Mock

<p align="left">
  <a href="https://github.com/hrishin/httpmock/actions"><img alt="GitHub Actions CI status" src="https://github.com/hrishin/httpmock/workflows/test/badge.svg"></a>
</p>

The `httpmock` is a simple mocking library to mock the [golang http client](https://golang.org/src/net/http/client.go) response.

Its intent is to reduce the boiler plate code required to write the unit test where a unit of code to test is wrapping the [http.Client](https://golang.org/src/net/http/client.go) or passing [http.Client](https://golang.org/src/net/http/client.go) as a dependency.


# Usage
1) Execute the `go get github.com/hrishin/httpmock`

2) Let's mock the response for `GET /foo` request which return the `HTTP 200 OK` response with some response body content

    ```
    import (
        "testing"
        "github.com/hrishin/httpmock"
    )


    func Test_some_test(t *testing.T) { // I prefer using _ as convention(readability)

        mockResponse := httpmock.Response{
            URI:        "/foo",
            StatusCode: 200,
            Body:       "bar response",
        }

        api := &someAPI{
            client: httpmock.Client(&mockResponse),
            API: "https://some.com/api/",
        }

        got := api.do()

        // assert got response from API
    }
    ```

In this example, `someAPI` is using [http.Client]() as a dependency in order to execute the HTTP request for a given API URL.