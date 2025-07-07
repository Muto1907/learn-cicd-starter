package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	type test struct {
		input          http.Header
		errExpected    bool
		expectedString string
	}
	tests := []test{
		{input: http.Header{}, errExpected: true, expectedString: ""},
		{input: http.Header{"Authorization": {"whatevertheKeyIS"}}, errExpected: true, expectedString: ""},
		{input: http.Header{"Authorization": {"ApiKey ws4343ewdwsd"}}, errExpected: false, expectedString: "ws4343ewdwsd"},
	}

	for _, tcase := range tests {
		resStr, resErr := GetAPIKey(tcase.input)
		if (resErr != nil) != tcase.errExpected {
			t.Fatalf("expected err: %t for error: %s got: %t ", tcase.errExpected, resErr, !tcase.errExpected)
		}
		if resStr != tcase.expectedString {
			t.Fatalf("expected string: %s got: %s", tcase.expectedString, resStr)
		}
	}
}
