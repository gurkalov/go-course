package copy

import (
	"io"
	"io/ioutil"
	"os"
)

func CopyOne(from, to string) error {
	input, err := ioutil.ReadFile(from)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(to, input, 0644)
}

func CopyTwo(from, to string) error {
	_, err := os.Stat(from)
	if err != nil {
		return err
	}

	source, err := os.Open(from)
	if err != nil {
		return err
	}
	defer source.Close()

	dest, err := os.Create(to)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, source)
	return err
}
