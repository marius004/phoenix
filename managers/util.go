package managers

import (
	"errors"
	"io/ioutil"
	"os"
)

func writeToFile(path string, data []byte) error {
	// open the file or create a new one in case it does not exist
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	return err
}

func readFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func makeDirectory(path string) error {
	err := os.Mkdir(path, 0755)

	if errors.Is(err, os.ErrExist) {
		return nil
	}

	return err
}

func deleteDirectory(path string) error {
	err := os.RemoveAll(path)

	if errors.Is(err, os.ErrExist) {
		return nil
	}

	return err
}

func renameDirectory(old, new string) error {
	return os.Rename(old, new)
}

func deleteFile(path string) error {
	return os.Remove(path)
}
