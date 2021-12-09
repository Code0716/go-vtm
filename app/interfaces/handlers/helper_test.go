package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
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

	if !Compare(t, want, gotBody) {
		t.Errorf("got = %v\n, want = %v\n", gotBody, want)
		return false
	}

	return true

}

func Compare(t *testing.T, want, got interface{}) bool {
	t.Helper()
	var isSame bool
	switch reflect.TypeOf(want).Kind() {
	case reflect.Struct:
		isSame = compareStruct(t, want, got)
	default:
		if !reflect.DeepEqual(want, got) {
			t.Errorf("got = %v\n, want = %v\n", got, want)
			isSame = false
		}
	}

	return isSame
}

func compareStruct(t *testing.T, want, got interface{}) bool {
	opt := cmp.AllowUnexported(want, got)
	if !cmp.Equal(want, got, opt) {
		t.Errorf("%v value is mismatch (-want +got):\n%s", reflect.TypeOf(want), cmp.Diff(want, got, opt))
		return false
	}

	return true
}
