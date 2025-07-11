package service

import (
	"strings"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"errors"
)

func DefineText(data string) (error, string) {
	if data == "" {
		return errors.New("an empty string is received"), ""
	}
	if strings.ContainsAny(data, ".- ") {
		return nil, morse.ToText(data)
	}
	return nil, morse.ToMorse(data)
}