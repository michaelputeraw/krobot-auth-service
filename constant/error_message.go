package constant

import "fmt"

var (
	ErrEmailAlreadyTaken = map[string]string{
		LANG_DEFAULT: fmt.Sprintf("%s already taken", Fields["email"][LANG_DEFAULT]),
		LANG_ID:      fmt.Sprintf("%s sudah digunakan", Fields["email"][LANG_ID]),
	}
)
