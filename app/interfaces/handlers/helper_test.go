package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

// wantRes is common expected response type.
//   code is expected http status code.
//   body is expected response body
type wantRes struct {
	code int
	body interface{}
}

func newTestEchoContext(t *testing.T, req *http.Request) (echo.Context, *httptest.ResponseRecorder) {
	t.Helper()

	rec := httptest.NewRecorder()

	e := echo.New()

	c := e.NewContext(req, rec)

	return c, rec
}

//　TODO:このままでは、member_test.goなど、ネストにポインターがあるものが比較できない。
func testJSON(t *testing.T, got []byte, want interface{}) bool {
	t.Helper()

	if want == nil {
		t.Fatal("testJSON() want is nil")
		return false
	}

	if len(got) == 0 {
		t.Fatal("testJSON() got is 0 length")
		return false
	}

	gotType := reflect.TypeOf(want)
	el := reflect.New(gotType)

	gotBody := el.Interface()
	err := json.Unmarshal(got, &gotBody)
	if err != nil {
		t.Fatal(err)
	}

	if reflect.TypeOf(want).Kind() != reflect.Ptr {
		gotBody = reflect.ValueOf(gotBody).Elem().Interface()
	}

	if !reflect.DeepEqual(gotBody, want) {
		t.Errorf("got = %v, want = %v", gotBody, want)
		return false
	}

	return true

}
