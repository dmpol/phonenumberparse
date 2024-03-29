package phonenumber

import (
	"errors"
	"regexp"
	"strconv"
)

var code = map[int]string{
	1: "7",
	2: "8",
	3: "+7",
}

/*
Функция Parse на вход принимает строку в виде номера телефона
(например: 8(123)345-67-89, +7123-345-6789 и т.д.), а так же код префикса
для определения начала итогового номера, где 1: 7, 2: 8, 3: +7.

Пример:

	на вход подается строка вида "8(123)345-67-89" и префиксом 1
	на выходе будет строка вида "71233456789".
*/
func Parse(phoneNumber string, prefix int) (string, error) {
	if len(phoneNumber) != 0 {
		//Удаляем всё, кроме цифр
		regex := regexp.MustCompile("[^0-9]+")
		cleanNumber := regex.ReplaceAllString(phoneNumber, "")

		if len(cleanNumber) == 11 {
			operator := cleanNumber[1:4]
			err := operatorValidation(operator)
			if err != nil {
				return "", err
			}
			//Меняем первую цифру на значение из map
			resNumber := code[prefix] + cleanNumber[1:]
			return resNumber, nil
		} else if len(cleanNumber) == 10 {
			operator := cleanNumber[:3]
			err := operatorValidation(operator)
			if err != nil {
				return "", err
			}
			//Добавляем значение из map
			resNumber := code[prefix] + cleanNumber
			return resNumber, nil
		}
	}
	return "", errors.New("Некорректный номер")
}

// operatorValidation проверяет номер оператора мобильного номера в пределах от 900 до 1000
func operatorValidation(operator string) error {
	num, _ := strconv.Atoi(operator)
	if 900 <= num && num < 1000 {
		return nil
	}
	return errors.New("Некорректный код оператора номера")
}
