package lolz

import (
	"io"
)

// Panik panics
func Panik(i any) {
	panic(i)
}

// HideThePain throws an error to io.Discard
func HideThePain(err error) {
	_, _ = io.Discard.Write([]byte(err.Error()))
}