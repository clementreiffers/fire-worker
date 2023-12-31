//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 clementreiffers.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobBuilder) DeepCopyInto(out *JobBuilder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobBuilder.
func (in *JobBuilder) DeepCopy() *JobBuilder {
	if in == nil {
		return nil
	}
	out := new(JobBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobBuilder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobBuilderList) DeepCopyInto(out *JobBuilderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JobBuilder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobBuilderList.
func (in *JobBuilderList) DeepCopy() *JobBuilderList {
	if in == nil {
		return nil
	}
	out := new(JobBuilderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobBuilderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobBuilderSpec) DeepCopyInto(out *JobBuilderSpec) {
	*out = *in
	if in.ScriptUrls != nil {
		in, out := &in.ScriptUrls, &out.ScriptUrls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ScriptNames != nil {
		in, out := &in.ScriptNames, &out.ScriptNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobBuilderSpec.
func (in *JobBuilderSpec) DeepCopy() *JobBuilderSpec {
	if in == nil {
		return nil
	}
	out := new(JobBuilderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobBuilderStatus) DeepCopyInto(out *JobBuilderStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobBuilderStatus.
func (in *JobBuilderStatus) DeepCopy() *JobBuilderStatus {
	if in == nil {
		return nil
	}
	out := new(JobBuilderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodTemplateWorkerAccount) DeepCopyInto(out *PodTemplateWorkerAccount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodTemplateWorkerAccount.
func (in *PodTemplateWorkerAccount) DeepCopy() *PodTemplateWorkerAccount {
	if in == nil {
		return nil
	}
	out := new(PodTemplateWorkerAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Worker) DeepCopyInto(out *Worker) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Worker.
func (in *Worker) DeepCopy() *Worker {
	if in == nil {
		return nil
	}
	out := new(Worker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerAccount) DeepCopyInto(out *WorkerAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerAccount.
func (in *WorkerAccount) DeepCopy() *WorkerAccount {
	if in == nil {
		return nil
	}
	out := new(WorkerAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerAccountList) DeepCopyInto(out *WorkerAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkerAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerAccountList.
func (in *WorkerAccountList) DeepCopy() *WorkerAccountList {
	if in == nil {
		return nil
	}
	out := new(WorkerAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerAccountSpec) DeepCopyInto(out *WorkerAccountSpec) {
	*out = *in
	in.WorkerReleaseSelector.DeepCopyInto(&out.WorkerReleaseSelector)
	out.PodTemplate = in.PodTemplate
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerAccountSpec.
func (in *WorkerAccountSpec) DeepCopy() *WorkerAccountSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerAccountStatus) DeepCopyInto(out *WorkerAccountStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerAccountStatus.
func (in *WorkerAccountStatus) DeepCopy() *WorkerAccountStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerBundle) DeepCopyInto(out *WorkerBundle) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerBundle.
func (in *WorkerBundle) DeepCopy() *WorkerBundle {
	if in == nil {
		return nil
	}
	out := new(WorkerBundle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerBundle) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerBundleList) DeepCopyInto(out *WorkerBundleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkerBundle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerBundleList.
func (in *WorkerBundleList) DeepCopy() *WorkerBundleList {
	if in == nil {
		return nil
	}
	out := new(WorkerBundleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerBundleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerBundlePodTemplate) DeepCopyInto(out *WorkerBundlePodTemplate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerBundlePodTemplate.
func (in *WorkerBundlePodTemplate) DeepCopy() *WorkerBundlePodTemplate {
	if in == nil {
		return nil
	}
	out := new(WorkerBundlePodTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerBundleSpec) DeepCopyInto(out *WorkerBundleSpec) {
	*out = *in
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = make([]Worker, len(*in))
		copy(*out, *in)
	}
	out.PodTemplate = in.PodTemplate
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerBundleSpec.
func (in *WorkerBundleSpec) DeepCopy() *WorkerBundleSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerBundleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerBundleStatus) DeepCopyInto(out *WorkerBundleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerBundleStatus.
func (in *WorkerBundleStatus) DeepCopy() *WorkerBundleStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerBundleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerDeployment) DeepCopyInto(out *WorkerDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerDeployment.
func (in *WorkerDeployment) DeepCopy() *WorkerDeployment {
	if in == nil {
		return nil
	}
	out := new(WorkerDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerDeploymentList) DeepCopyInto(out *WorkerDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkerDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerDeploymentList.
func (in *WorkerDeploymentList) DeepCopy() *WorkerDeploymentList {
	if in == nil {
		return nil
	}
	out := new(WorkerDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerDeploymentSpec) DeepCopyInto(out *WorkerDeploymentSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerDeploymentSpec.
func (in *WorkerDeploymentSpec) DeepCopy() *WorkerDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerDeploymentStatus) DeepCopyInto(out *WorkerDeploymentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerDeploymentStatus.
func (in *WorkerDeploymentStatus) DeepCopy() *WorkerDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerDeploymentTemplate) DeepCopyInto(out *WorkerDeploymentTemplate) {
	*out = *in
	if in.ScriptsUrls != nil {
		in, out := &in.ScriptsUrls, &out.ScriptsUrls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerDeploymentTemplate.
func (in *WorkerDeploymentTemplate) DeepCopy() *WorkerDeploymentTemplate {
	if in == nil {
		return nil
	}
	out := new(WorkerDeploymentTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerRelease) DeepCopyInto(out *WorkerRelease) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerRelease.
func (in *WorkerRelease) DeepCopy() *WorkerRelease {
	if in == nil {
		return nil
	}
	out := new(WorkerRelease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerRelease) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerReleaseList) DeepCopyInto(out *WorkerReleaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkerRelease, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerReleaseList.
func (in *WorkerReleaseList) DeepCopy() *WorkerReleaseList {
	if in == nil {
		return nil
	}
	out := new(WorkerReleaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerReleaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerReleaseSpec) DeepCopyInto(out *WorkerReleaseSpec) {
	*out = *in
	if in.WorkerVersions != nil {
		in, out := &in.WorkerVersions, &out.WorkerVersions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerReleaseSpec.
func (in *WorkerReleaseSpec) DeepCopy() *WorkerReleaseSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerReleaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerReleaseStatus) DeepCopyInto(out *WorkerReleaseStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerReleaseStatus.
func (in *WorkerReleaseStatus) DeepCopy() *WorkerReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerVersion) DeepCopyInto(out *WorkerVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerVersion.
func (in *WorkerVersion) DeepCopy() *WorkerVersion {
	if in == nil {
		return nil
	}
	out := new(WorkerVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerVersionList) DeepCopyInto(out *WorkerVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkerVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerVersionList.
func (in *WorkerVersionList) DeepCopy() *WorkerVersionList {
	if in == nil {
		return nil
	}
	out := new(WorkerVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerVersionSpec) DeepCopyInto(out *WorkerVersionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerVersionSpec.
func (in *WorkerVersionSpec) DeepCopy() *WorkerVersionSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerVersionStatus) DeepCopyInto(out *WorkerVersionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerVersionStatus.
func (in *WorkerVersionStatus) DeepCopy() *WorkerVersionStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerVersionStatus)
	in.DeepCopyInto(out)
	return out
}
