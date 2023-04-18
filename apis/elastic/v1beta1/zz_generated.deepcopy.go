//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudElasticsearch) DeepCopyInto(out *CloudElasticsearch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudElasticsearch.
func (in *CloudElasticsearch) DeepCopy() *CloudElasticsearch {
	if in == nil {
		return nil
	}
	out := new(CloudElasticsearch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudElasticsearch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudElasticsearchList) DeepCopyInto(out *CloudElasticsearchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudElasticsearch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudElasticsearchList.
func (in *CloudElasticsearchList) DeepCopy() *CloudElasticsearchList {
	if in == nil {
		return nil
	}
	out := new(CloudElasticsearchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudElasticsearchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudElasticsearchObservation) DeepCopyInto(out *CloudElasticsearchObservation) {
	*out = *in
	if in.ElasticCloudDeploymentID != nil {
		in, out := &in.ElasticCloudDeploymentID, &out.ElasticCloudDeploymentID
		*out = new(string)
		**out = **in
	}
	if in.ElasticCloudEmailAddress != nil {
		in, out := &in.ElasticCloudEmailAddress, &out.ElasticCloudEmailAddress
		*out = new(string)
		**out = **in
	}
	if in.ElasticCloudSsoDefaultURL != nil {
		in, out := &in.ElasticCloudSsoDefaultURL, &out.ElasticCloudSsoDefaultURL
		*out = new(string)
		**out = **in
	}
	if in.ElasticCloudUserID != nil {
		in, out := &in.ElasticCloudUserID, &out.ElasticCloudUserID
		*out = new(string)
		**out = **in
	}
	if in.ElasticsearchServiceURL != nil {
		in, out := &in.ElasticsearchServiceURL, &out.ElasticsearchServiceURL
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KibanaServiceURL != nil {
		in, out := &in.KibanaServiceURL, &out.KibanaServiceURL
		*out = new(string)
		**out = **in
	}
	if in.KibanaSsoURI != nil {
		in, out := &in.KibanaSsoURI, &out.KibanaSsoURI
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Logs != nil {
		in, out := &in.Logs, &out.Logs
		*out = make([]LogsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MonitoringEnabled != nil {
		in, out := &in.MonitoringEnabled, &out.MonitoringEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudElasticsearchObservation.
func (in *CloudElasticsearchObservation) DeepCopy() *CloudElasticsearchObservation {
	if in == nil {
		return nil
	}
	out := new(CloudElasticsearchObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudElasticsearchParameters) DeepCopyInto(out *CloudElasticsearchParameters) {
	*out = *in
	if in.ElasticCloudEmailAddress != nil {
		in, out := &in.ElasticCloudEmailAddress, &out.ElasticCloudEmailAddress
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Logs != nil {
		in, out := &in.Logs, &out.Logs
		*out = make([]LogsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MonitoringEnabled != nil {
		in, out := &in.MonitoringEnabled, &out.MonitoringEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudElasticsearchParameters.
func (in *CloudElasticsearchParameters) DeepCopy() *CloudElasticsearchParameters {
	if in == nil {
		return nil
	}
	out := new(CloudElasticsearchParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudElasticsearchSpec) DeepCopyInto(out *CloudElasticsearchSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudElasticsearchSpec.
func (in *CloudElasticsearchSpec) DeepCopy() *CloudElasticsearchSpec {
	if in == nil {
		return nil
	}
	out := new(CloudElasticsearchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudElasticsearchStatus) DeepCopyInto(out *CloudElasticsearchStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudElasticsearchStatus.
func (in *CloudElasticsearchStatus) DeepCopy() *CloudElasticsearchStatus {
	if in == nil {
		return nil
	}
	out := new(CloudElasticsearchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilteringTagObservation) DeepCopyInto(out *FilteringTagObservation) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilteringTagObservation.
func (in *FilteringTagObservation) DeepCopy() *FilteringTagObservation {
	if in == nil {
		return nil
	}
	out := new(FilteringTagObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilteringTagParameters) DeepCopyInto(out *FilteringTagParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilteringTagParameters.
func (in *FilteringTagParameters) DeepCopy() *FilteringTagParameters {
	if in == nil {
		return nil
	}
	out := new(FilteringTagParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogsObservation) DeepCopyInto(out *LogsObservation) {
	*out = *in
	if in.FilteringTag != nil {
		in, out := &in.FilteringTag, &out.FilteringTag
		*out = make([]FilteringTagObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SendActivityLogs != nil {
		in, out := &in.SendActivityLogs, &out.SendActivityLogs
		*out = new(bool)
		**out = **in
	}
	if in.SendAzureadLogs != nil {
		in, out := &in.SendAzureadLogs, &out.SendAzureadLogs
		*out = new(bool)
		**out = **in
	}
	if in.SendSubscriptionLogs != nil {
		in, out := &in.SendSubscriptionLogs, &out.SendSubscriptionLogs
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogsObservation.
func (in *LogsObservation) DeepCopy() *LogsObservation {
	if in == nil {
		return nil
	}
	out := new(LogsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogsParameters) DeepCopyInto(out *LogsParameters) {
	*out = *in
	if in.FilteringTag != nil {
		in, out := &in.FilteringTag, &out.FilteringTag
		*out = make([]FilteringTagParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SendActivityLogs != nil {
		in, out := &in.SendActivityLogs, &out.SendActivityLogs
		*out = new(bool)
		**out = **in
	}
	if in.SendAzureadLogs != nil {
		in, out := &in.SendAzureadLogs, &out.SendAzureadLogs
		*out = new(bool)
		**out = **in
	}
	if in.SendSubscriptionLogs != nil {
		in, out := &in.SendSubscriptionLogs, &out.SendSubscriptionLogs
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogsParameters.
func (in *LogsParameters) DeepCopy() *LogsParameters {
	if in == nil {
		return nil
	}
	out := new(LogsParameters)
	in.DeepCopyInto(out)
	return out
}
