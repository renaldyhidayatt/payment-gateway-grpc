package randomvcc

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomVCC() (string, error) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	randomNumber := ""
	for i := 0; i < 15; i++ {
		randomNumber += strconv.Itoa(rand.Intn(10))
	}

	checkDigit := calculateCheckDigit("4" + randomNumber)

	creditCardNumber := "4" + randomNumber + strconv.Itoa(checkDigit)

	return creditCardNumber, nil
}

func calculateCheckDigit(number string) int {
	sum := 0
	alternate := false
	for i := len(number) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(number[i]))
		if alternate {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		alternate = !alternate
	}
	return (10 - (sum % 10)) % 10
}
