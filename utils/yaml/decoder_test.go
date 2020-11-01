package yaml

import (
	"github.com/tanalam2411/olms/utils/rest"
	"testing"
)

func SplitYAMLDocumentTestData() ([]byte, error) {
	olmCRDUrl := "https://github.com/operator-framework/operator-lifecycle-manager/releases/download/0.16.1/crds.yaml"
	result, err := rest.HttpGET(olmCRDUrl)
	return result, err
}

func TestSplitYAMLDocument(t *testing.T) {

	testData, err := SplitYAMLDocumentTestData()

	if err != nil {
		t.Errorf("Failed to get TestData, err: %v", err)
	}

	_, err = SplitYAML(testData)
	if err != nil {
		t.Errorf("Failed, err: %v", err)
	}

}
