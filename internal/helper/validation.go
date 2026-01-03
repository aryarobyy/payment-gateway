package helper

import (
	"fmt"
	"strings"
)

func IsEmailValid(e string) bool {
	if e == "" {
		return false
	}

	for _, r := range e {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '@' || r ==
			'.' || r == '_' || r == '-') {
			return false
		}
	}
	return true
}

func EmptyCheck(fields map[string]string) error {
	for name, value := range fields {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("%s cannot be empty", name)
		}
	}
	return nil
}
