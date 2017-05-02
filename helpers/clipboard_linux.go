package helpers

import (
	"errors"
	"os/exec"
)

var Clipboard clipboard

type clipboard struct{}

func (self clipboard) Add(str string) error {
	if _, err := exec.LookPath(xclip); err == nil {
		return errors.New("Error xclip needed to use clipboard")
	}

	copy := exec.Command(xclip, "-in", "-selection", "clipboard")

	in, err := copy.StdinPipe()
	if err != nil {
		return err
	}

	if err := copy.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(str)); err != nil {
		return err
	}
	if err := in.Close(); err != nil {
		return err
	}
	copy.Wait()

	return nil
}
