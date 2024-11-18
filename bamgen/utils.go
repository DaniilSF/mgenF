package bamgen

import (
	"errors"
	"regexp"
	"strings"
)

func safeGet(parts []string, index int) string {
	if index < len(parts) {
		return parts[index]
	}
	return ""
}

func cleanField(field string) string {
	if strings.EqualFold(field, "NULL") {
		return ""
	}
	return field
}

func isEmail(field string) (bool, error) {
	field = strings.TrimSpace(field)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(field) {
		return false, errors.New("invalid email format")
	}
	return true, nil
}

func isValidUsername(username string) (bool, error) {
	username = strings.TrimSpace(username)
	re := regexp.MustCompile(`^[a-zA-Zа-яА-ЯёЁ0-9._@-]+$`)
	if !re.MatchString(username) {
		return false, errors.New("invalid username format")
	}
	return true, nil
}
