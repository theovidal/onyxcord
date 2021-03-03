package onyxcord

import (
	"io/ioutil"

	"github.com/fatih/color"
)

// OpenFile opens a file from a path
func OpenFile(path string) (data []byte, err error) {
	data, err = ioutil.ReadFile(path)
	return
}

// Red is a tool to display Red color into the term
var Red = color.New(color.FgRed)

// Green is a tool to display Green color into the term
var Green = color.New(color.FgGreen)
