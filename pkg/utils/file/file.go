package file

import (
	"io/ioutil"
	"mime/multipart"
)

func ReadFile(filepath string) ([]byte, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func ReadAll(file multipart.File) ([]byte, error) {
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}
