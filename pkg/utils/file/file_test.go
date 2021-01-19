package file

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	content, err := ReadFile("../../../manifests/olm/v0.16.1/olm.yaml")
	if err != nil {
		t.Errorf("Failed to read file, err: %v", err)
	}

	fmt.Println(string(content))
}
