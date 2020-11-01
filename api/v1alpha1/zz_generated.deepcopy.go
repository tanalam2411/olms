// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OLMCatalogSource) DeepCopyInto(out *OLMCatalogSource) {
	*out = *in
	in.CSApiVersion.DeepCopyInto(&out.CSApiVersion)
	in.CS.DeepCopyInto(&out.CS)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OLMCatalogSource.
func (in *OLMCatalogSource) DeepCopy() *OLMCatalogSource {
	if in == nil {
		return nil
	}
	out := new(OLMCatalogSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OLMClusterRole) DeepCopyInto(out *OLMClusterRole) {
	*out = *in
	in.CRApiVersion.DeepCopyInto(&out.CRApiVersion)
	in.CR.DeepCopyInto(&out.CR)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OLMClusterRole.
func (in *OLMClusterRole) DeepCopy() *OLMClusterRole {
	if in == nil {
		return nil
	}
	out := new(OLMClusterRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OLMClusterRoleBinding) DeepCopyInto(out *OLMClusterRoleBinding) {
	*out = *in
	in.CRBApiVersion.DeepCopyInto(&out.CRBApiVersion)
	in.CRB.DeepCopyInto(&out.CRB)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OLMClusterRoleBinding.
func (in *OLMClusterRoleBinding) DeepCopy() *OLMClusterRoleBinding {
	if in == nil {
		return nil
	}
	out := new(OLMClusterRoleBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OLMClusterServiceVersion) DeepCopyInto(out *OLMClusterServiceVersion) {
	*out = *in
	in.CSVApiVersion.DeepCopyInto(&out.CSVApiVersion)
	in.CSV.DeepCopyInto(&out.CSV)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OLMClusterServiceVersion.
func (in *OLMClusterServiceVersion) DeepCopy() *OLMClusterServiceVersion {
	if in == nil {
		return nil
	}
	out := new(OLMClusterServiceVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OLMCrdDetails) DeepCopyInto(out *OLMCrdDetails) {
	*out = *in
	in.CrdApiVersion.DeepCopyInto(&out.CrdApiVersion)
	if in.Crd != nil {
		in, out := &in.Crd, &out.Crd
		*out = new(apiextensions.CustomResourceDefinition)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OLMCrdDetails.
func (in *OLMCrdDetails) DeepCopy() *OLMCrdDetails {
	if in == nil {
		return nil
	}
	out := new(OLMCrdDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OLMDeployment) DeepCopyInto(out *OLMDeployment) {
	*out = *in
	in.DeploymentApiVersion.DeepCopyInto(&out.DeploymentApiVersion)
	in.Deployment.DeepCopyInto(&out.Deployment)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OLMDeployment.
func (in *OLMDeployment) DeepCopy() *OLMDeployment {
	if in == nil {
		return nil
	}
	out := new(OLMDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OLMNamespace) DeepCopyInto(out *OLMNamespace) {
	*out = *in
	in.NSApiVersion.DeepCopyInto(&out.NSApiVersion)
	in.NS.DeepCopyInto(&out.NS)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OLMNamespace.
func (in *OLMNamespace) DeepCopy() *OLMNamespace {
	if in == nil {
		return nil
	}
	out := new(OLMNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OLMOperatorGroup) DeepCopyInto(out *OLMOperatorGroup) {
	*out = *in
	in.OGApiVersion.DeepCopyInto(&out.OGApiVersion)
	in.OG.DeepCopyInto(&out.OG)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OLMOperatorGroup.
func (in *OLMOperatorGroup) DeepCopy() *OLMOperatorGroup {
	if in == nil {
		return nil
	}
	out := new(OLMOperatorGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OLMS) DeepCopyInto(out *OLMS) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OLMS.
func (in *OLMS) DeepCopy() *OLMS {
	if in == nil {
		return nil
	}
	out := new(OLMS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OLMS) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OLMSList) DeepCopyInto(out *OLMSList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OLMS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OLMSList.
func (in *OLMSList) DeepCopy() *OLMSList {
	if in == nil {
		return nil
	}
	out := new(OLMSList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OLMSList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OLMSSpec) DeepCopyInto(out *OLMSSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OLMSSpec.
func (in *OLMSSpec) DeepCopy() *OLMSSpec {
	if in == nil {
		return nil
	}
	out := new(OLMSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OLMSStatus) DeepCopyInto(out *OLMSStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OLMSStatus.
func (in *OLMSStatus) DeepCopy() *OLMSStatus {
	if in == nil {
		return nil
	}
	out := new(OLMSStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OLMServiceAccount) DeepCopyInto(out *OLMServiceAccount) {
	*out = *in
	in.SAApiVersion.DeepCopyInto(&out.SAApiVersion)
	in.SA.DeepCopyInto(&out.SA)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OLMServiceAccount.
func (in *OLMServiceAccount) DeepCopy() *OLMServiceAccount {
	if in == nil {
		return nil
	}
	out := new(OLMServiceAccount)
	in.DeepCopyInto(out)
	return out
}
