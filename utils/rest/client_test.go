package rest

import (
	"fmt"
	"testing"
)

const (
	olmCRDUrl = "https://github.com/operator-framework/operator-lifecycle-manager/releases/download/0.16.1/crds.yaml"
)

// temp function, will be replaced by rest client interface
// test HttpGet function
func TestHttpGET(t *testing.T) {

	fmt.Println("InputURL: ", olmCRDUrl)
	result, err := HttpGET(olmCRDUrl)

	if err != nil {
		t.Errorf("HttpGET failed, error: %v", err)
	}

	fmt.Printf("Result: %v", string(result))
}
