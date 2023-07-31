package funclist

import (
	myos "go-samples/samples/os"
	mystrings "go-samples/samples/strings"
)

func GetFunc(name string) []func() {
	osfunc := []func(){
		myos.OsSample001,
		myos.OsSample002,
		myos.OsSample003,
		myos.OsSample004,
		myos.OsSample005,
		myos.OsSample006,
		myos.OsSample007,
		myos.OsSample008,
		myos.OsSample009,
		myos.OsSample010,
		myos.OsSample011,
		myos.OsSample012,
	}
	stringsfunc := []func(){
		mystrings.StringsSample001,
		mystrings.StringsSample002,
	}

	switch name {
	case "os":
		return osfunc
	case "strings":
		return stringsfunc
	default:
		return nil
	}
}
