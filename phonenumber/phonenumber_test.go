package phonenumber

import (
	"testing"
)

func TestParse(t *testing.T) {
	phoneNumber := "8(123)4567890"
	prefix := 1
	result := "71234567890"

	resFunc, err := Parse(phoneNumber, prefix)
	if err != nil {
		t.Errorf("Ошибка в функции")
	}

	if resFunc != result {
		t.Errorf("Ожидаем: %s, получили: %s", result, resFunc)
	}
}
