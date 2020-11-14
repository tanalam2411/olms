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
	v15 "github.com/operator-framework/api/pkg/operators/v1"
	"github.com/operator-framework/api/pkg/operators/v1alpha1"
	olmsgv1alpha1 "github.com/tanalam2411/olms/api/v1alpha1"
	"github.com/tanalam2411/olms/utils/k8s"
	"github.com/tanalam2411/olms/utils/olm"
	"github.com/tanalam2411/olms/utils/rest"
	"github.com/tanalam2411/olms/utils/yaml"
	v14 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	v13 "k8s.io/api/rbac/v1"
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

	log.Info("Found OLMS with configurations", "object:: ", olms)

	config, err := k8s.GetClusterConfig()
	if err != nil {
		log.Error(err, "Failed to fetch Cluster Config")
	}

	apiExtkubeClient, err := apiextension.NewForConfig(config)
	if err != nil {
		log.Error(err, "Failed to create Api Extension KubeClient.")
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

		// https://www.velotio.com/engineering-blog/extending-kubernetes-apis-with-custom-resource-definitions-crds
		createdCrd, err := apiExtkubeClient.ApiextensionsV1().CustomResourceDefinitions().Create(context.TODO(), crdObj, v1.CreateOptions{})
		if err != nil {
			log.Error(err, "Failed to Create CRD object within the Cluster.")
		} else {
			fmt.Printf("CRD got created: %v \n", createdCrd)
		}

	}

	// Get all the OLM Resources([][]byte) from OLM's release, create in the cluster
	resDefinitions, err := GetOLMResourcesDefinitions()
	if err != nil {
		log.Error(err, "Failed to fetch OLM's CRDs")
	}

	kubeClient, err := k8s.GetClientSet(config)
	if err != nil {
		log.Error(err, "Failed to create Cluster ClientSet")
	}

	olmClient, err := olm.GetOLMClientSet(config)
	if err != nil {
		log.Error(err, "Failed to create OLM ClientSet")
	}

	for _, resDef := range resDefinitions {
		sch := runtime.NewScheme()
		_ = clientgoscheme.AddToScheme(sch)
		_ = apiextv1beta1.AddToScheme(sch)
		_ = v15.AddToScheme(sch)
		_ = v1alpha1.AddToScheme(sch)

		decode := serializer.NewCodecFactory(sch).UniversalDeserializer().Decode
		obj, _, err := decode([]byte(resDef), nil, nil)
		if err != nil {
			fmt.Printf("%#v", err)
		}

		if obj != nil {

			gvk := obj.GetObjectKind()
			if gvk != nil {

				switch obj.(type) {
				case *v12.Namespace:
					nsObj, err := yaml.YAMLToNamespace(resDef)
					if err != nil {
						log.Error(err, "Failed to convert YAMl to Namespace")
					}

					nsClient := kubeClient.CoreV1().Namespaces()
					_, err = nsClient.Get(context.TODO(), nsObj.Name, v1.GetOptions{})
					if err != nil {
						log.Error(err, fmt.Sprintf("Failed to get Namespace by name: %v", nsObj.Name))

						_, err := nsClient.Create(context.TODO(), nsObj, v1.CreateOptions{})
						if err != nil {
							log.Error(err, "Failed to create Namespace")
						}
						log.Info(fmt.Sprintf("Created Namespace: %T, Value: %v", nsObj, nsObj))
					}

				case *v12.ServiceAccount:
					saObj, err := yaml.YAMLToServiceAccount(resDef)
					if err != nil {
						log.Error(err, "Failed to convert YAMl to ServiceAccount")
					}

					saClient := kubeClient.CoreV1().ServiceAccounts(saObj.Namespace)
					_, err = saClient.Get(context.TODO(), saObj.Name, v1.GetOptions{})
					if err != nil {
						log.Error(err, fmt.Sprintf("Failed to get ServiceAccount by name: %v", saObj.Name))

						_, err := saClient.Create(context.TODO(), saObj, v1.CreateOptions{})
						if err != nil {
							log.Error(err, "Failed to create ServiceAccount")
						}
						log.Info(fmt.Sprintf("Created ServiceAccount: %T, Value: %v", saObj, saObj))
					}

				case *v13.ClusterRole:
					crObj, err := yaml.YAMLToClusterRole(resDef)
					if err != nil {
						log.Error(err, "Failed to convert YAMl to ClusterRole")
					}

					crClient := kubeClient.RbacV1().ClusterRoles()
					_, err = crClient.Get(context.TODO(), crObj.Name, v1.GetOptions{})
					if err != nil {
						log.Error(err, fmt.Sprintf("Failed to get ClusterRole by name: %v", crObj.Name))

						_, err := crClient.Create(context.TODO(), crObj, v1.CreateOptions{})
						if err != nil {
							log.Error(err, "Failed to create ClusterRole")
						}
						log.Info(fmt.Sprintf("Created ClusterRole: %T, Value: %v", crObj, crObj))
					}

				case *v13.ClusterRoleBinding:
					crbObj, err := yaml.YAMLToClusterRoleBinding(resDef)
					if err != nil {
						log.Error(err, "Failed to convert YAMl to ClusterRoleBinding")
					}

					crbClient := kubeClient.RbacV1().ClusterRoleBindings()
					_, err = crbClient.Get(context.TODO(), crbObj.Name, v1.GetOptions{})
					if err != nil {
						log.Error(err, fmt.Sprintf("Failed to get ClusterRole by name: %v", crbObj.Name))

						_, err := crbClient.Create(context.TODO(), crbObj, v1.CreateOptions{})
						if err != nil {
							log.Error(err, "Failed to create ClusterRoleBinding")
						}
						log.Info(fmt.Sprintf("Created ClusterRoleBinding: %T, Value: %v", crbObj, crbObj))
					}

				case *v14.Deployment:
					deployObj, err := yaml.YAMLToDeployment(resDef)
					if err != nil {
						log.Error(err, "Failed to convert YAMl to Deployment")
					}

					deployClient := kubeClient.AppsV1().Deployments(deployObj.Namespace)
					_, err = deployClient.Get(context.TODO(), deployObj.Name, v1.GetOptions{})
					if err != nil {
						log.Error(err, fmt.Sprintf("Failed to get Deployment by name: %v", deployObj.Name))

						_, err := deployClient.Create(context.TODO(), deployObj, v1.CreateOptions{})
						if err != nil {
							log.Error(err, "Failed to create Deployment")
						}
						log.Info(fmt.Sprintf("Created Deployment: %T, Value: %v", deployObj, deployObj))
					}

				case *v15.OperatorGroup:
					ogObj, err := yaml.YAMLToOperatorGroup(resDef)
					if err != nil {
						log.Error(err, "Failed to convert YAMl to OperatorGroup")
					}

					ogClient := olmClient.OperatorsV1().OperatorGroups(ogObj.Namespace)
					_, err = ogClient.Get(context.TODO(), ogObj.Name, v1.GetOptions{})
					if err != nil {
						log.Error(err, fmt.Sprintf("Failed to get OperatorGroup by name: %v", ogObj.Name))

						_, err := ogClient.Create(context.TODO(), ogObj, v1.CreateOptions{})
						if err != nil {
							log.Error(err, "Failed to create OperatorGroup")
						}
						log.Info(fmt.Sprintf("Created OperatorGroup: %T, Value: %v", ogObj, ogObj))
					}

				case *v1alpha1.ClusterServiceVersion:
					csvObj, err := yaml.YAMLToClusterServiceVersion(resDef)
					if err != nil {
						log.Error(err, "Failed to convert YAMl to ClusterServiceVersion")
					}

					csvClient := olmClient.OperatorsV1alpha1().ClusterServiceVersions(csvObj.Namespace)
					_, err = csvClient.Get(context.TODO(), csvObj.Name, v1.GetOptions{})
					if err != nil {
						log.Error(err, fmt.Sprintf("Failed to get ClusterServiceVersion by name: %v", csvObj.Name))

						_, err := csvClient.Create(context.TODO(), csvObj, v1.CreateOptions{})
						if err != nil {
							log.Error(err, "Failed to create ClusterServiceVersion")
						}
						log.Info(fmt.Sprintf("Created ClusterServiceVersion: %T, Value: %v", csvObj, csvObj))
					}

				case *v1alpha1.CatalogSource:
					csObj, err := yaml.YAMLToCatalogSource(resDef)
					if err != nil {
						log.Error(err, "Failed to convert YAMl to CatalogSource")
					}

					csClient := olmClient.OperatorsV1alpha1().CatalogSources(csObj.Namespace)
					_, err = csClient.Get(context.TODO(), csObj.Name, v1.GetOptions{})
					if err != nil {
						log.Error(err, fmt.Sprintf("Failed to get CatalogSource by name: %v", csObj.Name))

						_, err := csClient.Create(context.TODO(), csObj, v1.CreateOptions{})
						if err != nil {
							log.Error(err, "Failed to create CatalogSource")
						}
						log.Info(fmt.Sprintf("Created CatalogSource: %T, Value: %v", csObj, csObj))
					}

				}

			}

		}

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
