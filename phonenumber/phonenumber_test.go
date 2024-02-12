package phonenumber

import (
	"testing"
)

func TestParse(t *testing.T) {
	phoneNumber := "8-(900)-456-78-90"
	prefix := 3
	result := "+79004567890"

	resFunc, err := Parse(phoneNumber, prefix)
	if err != nil {
		t.Errorf("%s", err)
	}

	if resFunc != result {
		t.Errorf("Ожидаем: %s, получили: %s", result, resFunc)
	}
}
