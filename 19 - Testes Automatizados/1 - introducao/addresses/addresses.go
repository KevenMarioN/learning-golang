package addresses

import "strings"

// TypeAddress is function for validation the address, is valid if contain rua,avenida, estrada or rodóvia
func TypeAddress(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodóvia"}

	firstWordAddress := strings.Split(strings.ToLower(address), " ")[0]
	isValidAddress := false
	for _, typeAddress := range validTypes {
		if typeAddress == firstWordAddress {
			isValidAddress = true
		}
	}
	if isValidAddress {
		return strings.Title(firstWordAddress)
	} else {
		return "Address invalid!"
	}
}
