package service

import (
	"strings"
	"github.com/magabrotheeeer/sprint6/pkg/morse"
	"errors"
)

func DefineText(data string) (string, error) {
	if data == "" {
		return "", errors.New("an empty string is received")
	}
	if strings.ContainsAny(data, ".- ") {
		return morse.ToText(data), nil
	}
	return morse.ToMorse(data), nil
}