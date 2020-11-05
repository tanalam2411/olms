package yaml

import (
	"bytes"
	"github.com/ghodss/yaml"
	goyaml "github.com/go-yaml/yaml"
	"io"
	v14 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	v13 "k8s.io/api/rbac/v1"
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

func YAMLToNamespace(nsBytes []byte) (*v12.Namespace, error) {

	var ns v12.Namespace
	err := yaml.Unmarshal(nsBytes, &ns)
	if err != nil {
		return nil, err
	}
	return &ns, nil
}

func YAMLToServiceAccount(saBytes []byte) (*v12.ServiceAccount, error) {

	var sa v12.ServiceAccount
	err := yaml.Unmarshal(saBytes, &sa)
	if err != nil {
		return nil, err
	}
	return &sa, nil
}

func YAMLToClusterRole(crBytes []byte) (*v13.ClusterRole, error) {

	var cr v13.ClusterRole
	err := yaml.Unmarshal(crBytes, &cr)
	if err != nil {
		return nil, err
	}
	return &cr, nil
}

func YAMLToClusterRoleBinding(crbBytes []byte) (*v13.ClusterRoleBinding, error) {

	var crb v13.ClusterRoleBinding
	err := yaml.Unmarshal(crbBytes, &crb)
	if err != nil {
		return nil, err
	}
	return &crb, nil
}

func YAMLToDeployment(deployBytes []byte) (*v14.Deployment, error) {

	var deploy v14.Deployment
	err := yaml.Unmarshal(deployBytes, &deploy)
	if err != nil {
		return nil, err
	}
	return &deploy, nil
}
