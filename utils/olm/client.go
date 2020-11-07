package olm

import (
	"k8s.io/client-go/rest"
	olmClient "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned"
)

func GetOLMClientSet(config *rest.Config) (*olmClient.Clientset, error) {

	clientset, err := olmClient.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
