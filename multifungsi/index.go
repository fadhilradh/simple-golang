package multifungsi

import (
	"fmt"
	"errors"
)


func Add(num1, num2 int) (int, error) {
	if (num1 < 0) || (num2 < 0) {
		return 0, errors.New("You can't put below zero there, buddy !")
	}
	result := num1 + num2
	return result, nil
}

func Subtract(num1, num2 int) (int, error) {
	if (num2 > num1) {
		return 0, errors.New("Wait, that would be minus, and we don't do that here")
	}
	result := num1 - num2
	return result, nil
}

func Multiply(num1, num2 int) (int, error) {
	if (num1 < 1) || (num2 < 1) {
		return 0, errors.New("We don't do 0 or minus multiplication here, buddy")
	}
	result := num1 * num2
	return result, nil
}

func Divide(num1, num2 int) (int, error) {
	if (num1 == 0) || (num2 == 0) {
		return 0, errors.New("We don't divide by zero here, you better use calculator for that :)")
	}
	result := num1 - num2
	return result, nil
}

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("Are you forgetting something?")
	}
	message := fmt.Sprintf("Howdy, %v ?", name)
	return message, nil
}