package helpers

import (
	"crypto/rand"
	"math/big"

	uuid "github.com/nu7hatch/gouuid"
)

// Function to check whether an array contains an element or not
func IsArrayContains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}

// Function to generate random string consists of only letter and number, with specified length
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

// This function is used to remove any field on a map which have empty string value
func RemoveEmptyField(data map[string]string) map[string]string {
	for k, v := range data {
		if v == "" {
			delete(data, k)
		}
	}
	return data
}

// Generate UUID
func GenerateUUID() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

// Extract values from nullable string (if null, we will set the value to an empty string)
func ExtractNullableString(val *string) string {
	if val != nil {
		return *val
	}
	return ""
}

// Convert a string to a pointer string
func CreatePointerString(str string) *string {
	tempStr := str
	return &tempStr
}
