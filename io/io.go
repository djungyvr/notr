package io

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
)

// CaptureText reads a line from an io.Reader
// It will read and string and strip off the newline
func CaptureText(reader bufio.Reader) (string, error) {
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	if len(text) > 0 {
		text = text[:len(text)-1]
	}

	return text, nil
}

// CaptureVim opens up vim and first writes content to a
// temporary file. It then reads the contents and returns it
// as a byte array along with any errors it may have encountered
func CaptureVim(content []byte) ([]byte, error) {
	tmpfile, err := ioutil.TempFile("", ".edit")

	if err != nil {
		return []byte{}, nil
	}

	// clean up the file
	defer os.Remove(tmpfile.Name())

	// write the contents to the file
	if _, err := tmpfile.Write(content); err != nil {
		return []byte{}, nil
	}

	// open up vim to edit the file
	cmd := exec.Command("vim", tmpfile.Name())
	// set the stdin and out otherwise vim won't open in terminal
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()

	if err != nil {
		return []byte{}, err
	}

	// read the contents from the file once vim finishes
	data, err := ioutil.ReadFile(tmpfile.Name())

	if err != nil {
		return []byte{}, err
	}

	return data, nil
}
