package k8s

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// InClusterConfig returns a config object which uses the service account
// kubernetes gives to pods. It's intended for clients that expect to be
// running inside a pod running on kubernetes. It will return ErrNotInCluster
// if called from a process not running in a kubernetes environment.
func GetInClusterConfig() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func GetClusterConfig() (*rest.Config, error) {

	//kubeconfig := flag.String("kubeconfig", "/home/afour/.kube/config", "kubeconfig file")
	//flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", "/home/afour/.kube/config")
	//config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		return nil, err
	}
	return config, nil
}

/*
https://github.com/codeformio/declare/blob/d776b43ab121808d7794befc1182f58f9934d40e/main.go#L62
https://github.com/codeformio/declare/blob/d776b43ab121808d7794befc1182f58f9934d40e/controllers/controller_crd_controller.go
https://github.com/kubernetes/client-go/blob/master/examples/in-cluster-client-configuration/main.go
*/
