package helpers

import "os/exec"

var Clipboard clipboard

type clipboard struct{}

func (self clipboard) Add(str string) error {
	copy := exec.Command("pbcopy")

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
