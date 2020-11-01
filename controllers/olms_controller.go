/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	olmsgv1alpha1 "github.com/tanalam2411/olms/api/v1alpha1"
	"github.com/tanalam2411/olms/utils/k8s"
	"github.com/tanalam2411/olms/utils/rest"
	"github.com/tanalam2411/olms/utils/yaml"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
)

const (
	olmCRDUrl                 = "https://github.com/operator-framework/operator-lifecycle-manager/releases/download/0.16.1/crds.yaml"
	olmResourcesDefinitionUrl = "https://github.com/operator-framework/operator-lifecycle-manager/releases/download/0.16.1/olm.yaml"
)

// OLMSReconciler reconciles a OLMS object
type OLMSReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=olmsg.olms.com,resources=olms,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=olmsg.olms.com,resources=olms/status,verbs=get;update;patch

func (r *OLMSReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("olms", strings.ToLower(olmsgv1alpha1.OLMSKind), "request", req.NamespacedName)

	olms := &olmsgv1alpha1.OLMS{}
	err := r.Get(ctx, req.NamespacedName, olms)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Info("OLMS resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get OLMS")
		return ctrl.Result{}, err
	}

	config, err := k8s.GetClusterConfig()
	if err != nil {
		log.Error(err, "Failed to fetch Cluster Config")
	}

	kubeClient, err := apiextension.NewForConfig(config)
	if err != nil {
		log.Error(err, "Failed to create client.")
	}

	// Get all CRDs([][]byte) of OLM from OLM's release, create in the cluster
	crds, err := GetOLMCrds()
	if err != nil {
		log.Error(err, "Failed to fetch OLM's CRDs")
	}

	for _, crd := range crds {
		crdObj, err := yaml.YAMLToCRD(crd)
		if err != nil {
			log.Error(err, "Failed to decode from YAMl to CRD")
		}
		fmt.Printf("crdObj: %v\n", crdObj)

		// https://www.velotio.com/engineering-blog/extending-kubernetes-apis-with-custom-resource-definitions-crds
		createdCrd, err := kubeClient.ApiextensionsV1().CustomResourceDefinitions().Create(context.TODO(), crdObj, v1.CreateOptions{})
		if err != nil {
			log.Error(err, "Failed to Create CRD object within the Cluster.")
		}
		fmt.Printf("CRD got created: %v \n", createdCrd)

	}

	// Get all the OLM Resources([][]byte) from OLM's release, create in the cluster
	resDefinitions, err := GetOLMResourcesDefinitions()
	if err != nil {
		log.Error(err, "Failed to fetch OLM's CRDs")
	}

	fmt.Println("============================================================================")
	for _, resDef := range resDefinitions {
		sch := runtime.NewScheme()
		_ = clientgoscheme.AddToScheme(sch)
		_ = apiextv1beta1.AddToScheme(sch)

		decode := serializer.NewCodecFactory(sch).UniversalDeserializer().Decode
		obj, _, err := decode([]byte(resDef), nil, nil)
		if err != nil {
			fmt.Printf("%#v", err)
		}

		gvk := obj.GetObjectKind()
		if gvk != nil {
			fmt.Printf("Object kind: %v \n", gvk)
		} else {
			fmt.Println("Object gvk nil : ", gvk)
		}
		fmt.Println("============================================================================")
	}

	return ctrl.Result{}, nil
}

func GetOLMCrds() ([][]byte, error) {
	olmCRDs, err := rest.HttpGET(olmCRDUrl)

	if err != nil {
		return nil, err
	}

	crds, err := yaml.SplitYAML(olmCRDs)
	if err != nil {
		return nil, err
	}

	return crds, nil
}

func GetOLMResourcesDefinitions() ([][]byte, error) {
	olmResDefinitions, err := rest.HttpGET(olmResourcesDefinitionUrl)

	if err != nil {
		return nil, err
	}

	resDefinitions, err := yaml.SplitYAML(olmResDefinitions)
	if err != nil {
		return nil, err
	}

	return resDefinitions, nil
}

func (r *OLMSReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&olmsgv1alpha1.OLMS{}).
		Complete(r)
}
