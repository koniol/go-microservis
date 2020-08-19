package home_test

import (
	"microservices/m/home"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers_Home(t *testing.T) {
	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		exceptedStatus int
		exceptedBody   string
	}{
		{
			name: "Home	Page",
			in:             httptest.NewRequest("GET", "/", nil),
			out:            httptest.NewRecorder(),
			exceptedBody:   "<h1>Ready to works</h1>",
			exceptedStatus: http.StatusOK,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			h := home.NewHandlers(nil, nil)
			h.Home(test.out, test.in)

			if test.out.Code != test.exceptedStatus {
				t.Logf("excepted code %d recived %d", test.exceptedStatus, test.out.Code)
				t.Fail()
			}

			body := test.out.Body.String()
			if body != test.exceptedBody {
				t.Logf("excepted body %s recived %s", test.exceptedBody, test.out.Body.String())
				t.Fail()
			}
		})
	}
}
