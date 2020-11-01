package yaml

import (
	"bytes"
	"github.com/ghodss/yaml"
	goyaml "github.com/go-yaml/yaml"
	"io"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// ref - https://gist.github.com/yanniszark/c6f347421a1eeb75057ff421e03fd57c
func SplitYAML(resources []byte) ([][]byte, error) {

	dec := goyaml.NewDecoder(bytes.NewReader(resources))

	var res [][]byte
	for {
		var value interface{}
		err := dec.Decode(&value)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		valueBytes, err := goyaml.Marshal(value)
		if err != nil {
			return nil, err
		}
		res = append(res, valueBytes)
	}
	return res, nil
}

// https://github.com/kubernetes/client-go/issues/193
func YAMLToCRD(crdBytes []byte) (*v1.CustomResourceDefinition, error) {

	var crd v1.CustomResourceDefinition
	err := yaml.Unmarshal(crdBytes, &crd)
	if err != nil {
		return nil, err
	}
	return &crd, nil
}
