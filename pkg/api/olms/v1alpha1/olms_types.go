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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	OLMSKind = "OLMS"
)

type OLMCrdDetails struct {
	CrdApiVersion metav1.APIVersions                      `json:"olmCrdApiVersion"`
	Crd           *apiextensions.CustomResourceDefinition `json:"olmCrd"`
	CrdName       string                                  `json:"olmCrdName"`
	CrdExists     bool                                    `json:"olmCrdExists"` // CRD's name: exists or not
}

type OLMNamespace struct {
	NSApiVersion metav1.APIVersions `json:"olmNSApiVersion"`
	NS           v1.Namespace       `json:"olmNS"`
	NSName       string             `json:"olmNSName"`
	NSExists     bool               `json:"olmNSExists"`
}

type OLMServiceAccount struct {
	SAApiVersion metav1.APIVersions `json:"olmSAApiVersion"`
	SA           v1.ServiceAccount  `json:"olmSA"`
	SAName       string             `json:"olmSAName"`
	SAExists     bool               `json:"olmSAExists"`
}

type OLMClusterRole struct {
	CRApiVersion metav1.APIVersions `json:"olmCRApiVersion"`
	CR           v1.ServiceAccount  `json:"olmCR"`
	CRName       string             `json:"olmCRName"`
	CRExists     bool               `json:"olmCRExists"`
}

type OLMClusterRoleBinding struct {
	CRBApiVersion metav1.APIVersions `json:"olmCRBApiVersion"`
	CRB           v1.ServiceAccount  `json:"olmCRB"`
	CRBName       string             `json:"olmCRBName"`
	CRBExists     bool               `json:"olmCRBExists"`
}

type OLMDeployment struct {
	DeploymentApiVersion metav1.APIVersions `json:"olmDeploymentApiVersion"`
	Deployment           v1.ServiceAccount  `json:"olmDeployment"`
	DeploymentName       string             `json:"olmDeploymentName"`
	DeploymentExists     bool               `json:"olmDeploymentExists"`
}

type OLMOperatorGroup struct {
	OGApiVersion metav1.APIVersions `json:"olmOGApiVersion"`
	OG           v1.ServiceAccount  `json:"olmOG"`
	OGName       string             `json:"olmOGName"`
	OGExists     bool               `json:"olmOGExists"`
}

type OLMClusterServiceVersion struct {
	CSVApiVersion metav1.APIVersions `json:"olmCRBApiVersion"`
	CSV           v1.ServiceAccount  `json:"olmCSV"`
	CSVName       string             `json:"olmCSVName"`
	CSVExists     bool               `json:"olmCSVExists"`
}

type OLMCatalogSource struct {
	CSApiVersion metav1.APIVersions `json:"olmCSApiVersion"`
	CS           v1.ServiceAccount  `json:"olmCS"`
	CSName       string             `json:"olmCSName"`
	CSExists     bool               `json:"olmCSExists"`
}

//type CustomCatalogSource struct {
//	metav1.TypeMeta `json:",inline"`
//	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
//}

// OLMSSpec defines the desired state of OLMS
type OLMSSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// OLMReleaseVersion - The olm release version to be installed
	OLMReleaseVersion string `json:"olmReleaseVersion"`

	// OLMNamespace - Namespace under which olm operator will be installed
	OLMNamespace string `json:"olmNamespace"`

	// OLMCatalogSource v1alpha1.CatalogSource // TODO: CustomCatalogSource details
}

// OLMSStatus defines the observed state of OLMS
type OLMSStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Whether OLM is installed in the cluster
	OLMIsInstalled string `json:"olmIsInstalled,omitempty"`

	// Version of OLM installed in the cluster
	OLMInstalledVersion string `json:"olmInstalledVersion,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// OLMS is the Schema for the olms API
type OLMS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OLMSSpec   `json:"spec,omitempty"`
	Status OLMSStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OLMSList contains a list of OLMS
type OLMSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OLMS `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OLMS{}, &OLMSList{})
}
