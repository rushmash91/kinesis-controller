//go:build !ignore_autogenerated

// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Consumer) DeepCopyInto(out *Consumer) {
	*out = *in
	if in.ConsumerCreationTimestamp != nil {
		in, out := &in.ConsumerCreationTimestamp, &out.ConsumerCreationTimestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Consumer.
func (in *Consumer) DeepCopy() *Consumer {
	if in == nil {
		return nil
	}
	out := new(Consumer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsumerDescription) DeepCopyInto(out *ConsumerDescription) {
	*out = *in
	if in.ConsumerCreationTimestamp != nil {
		in, out := &in.ConsumerCreationTimestamp, &out.ConsumerCreationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.StreamARN != nil {
		in, out := &in.StreamARN, &out.StreamARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsumerDescription.
func (in *ConsumerDescription) DeepCopy() *ConsumerDescription {
	if in == nil {
		return nil
	}
	out := new(ConsumerDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnhancedMetrics) DeepCopyInto(out *EnhancedMetrics) {
	*out = *in
	if in.ShardLevelMetrics != nil {
		in, out := &in.ShardLevelMetrics, &out.ShardLevelMetrics
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnhancedMetrics.
func (in *EnhancedMetrics) DeepCopy() *EnhancedMetrics {
	if in == nil {
		return nil
	}
	out := new(EnhancedMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Record) DeepCopyInto(out *Record) {
	*out = *in
	if in.ApproximateArrivalTimestamp != nil {
		in, out := &in.ApproximateArrivalTimestamp, &out.ApproximateArrivalTimestamp
		*out = (*in).DeepCopy()
	}
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Record.
func (in *Record) DeepCopy() *Record {
	if in == nil {
		return nil
	}
	out := new(Record)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardFilter) DeepCopyInto(out *ShardFilter) {
	*out = *in
	if in.Timestamp != nil {
		in, out := &in.Timestamp, &out.Timestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardFilter.
func (in *ShardFilter) DeepCopy() *ShardFilter {
	if in == nil {
		return nil
	}
	out := new(ShardFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartingPosition) DeepCopyInto(out *StartingPosition) {
	*out = *in
	if in.Timestamp != nil {
		in, out := &in.Timestamp, &out.Timestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartingPosition.
func (in *StartingPosition) DeepCopy() *StartingPosition {
	if in == nil {
		return nil
	}
	out := new(StartingPosition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stream) DeepCopyInto(out *Stream) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stream.
func (in *Stream) DeepCopy() *Stream {
	if in == nil {
		return nil
	}
	out := new(Stream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stream) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamDescription) DeepCopyInto(out *StreamDescription) {
	*out = *in
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.EnhancedMonitoring != nil {
		in, out := &in.EnhancedMonitoring, &out.EnhancedMonitoring
		*out = make([]*EnhancedMetrics, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(EnhancedMetrics)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.HasMoreShards != nil {
		in, out := &in.HasMoreShards, &out.HasMoreShards
		*out = new(bool)
		**out = **in
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
	if in.RetentionPeriodHours != nil {
		in, out := &in.RetentionPeriodHours, &out.RetentionPeriodHours
		*out = new(int64)
		**out = **in
	}
	if in.StreamARN != nil {
		in, out := &in.StreamARN, &out.StreamARN
		*out = new(string)
		**out = **in
	}
	if in.StreamCreationTimestamp != nil {
		in, out := &in.StreamCreationTimestamp, &out.StreamCreationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.StreamModeDetails != nil {
		in, out := &in.StreamModeDetails, &out.StreamModeDetails
		*out = new(StreamModeDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.StreamName != nil {
		in, out := &in.StreamName, &out.StreamName
		*out = new(string)
		**out = **in
	}
	if in.StreamStatus != nil {
		in, out := &in.StreamStatus, &out.StreamStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamDescription.
func (in *StreamDescription) DeepCopy() *StreamDescription {
	if in == nil {
		return nil
	}
	out := new(StreamDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamDescriptionSummary) DeepCopyInto(out *StreamDescriptionSummary) {
	*out = *in
	if in.ConsumerCount != nil {
		in, out := &in.ConsumerCount, &out.ConsumerCount
		*out = new(int64)
		**out = **in
	}
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.EnhancedMonitoring != nil {
		in, out := &in.EnhancedMonitoring, &out.EnhancedMonitoring
		*out = make([]*EnhancedMetrics, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(EnhancedMetrics)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
	if in.OpenShardCount != nil {
		in, out := &in.OpenShardCount, &out.OpenShardCount
		*out = new(int64)
		**out = **in
	}
	if in.RetentionPeriodHours != nil {
		in, out := &in.RetentionPeriodHours, &out.RetentionPeriodHours
		*out = new(int64)
		**out = **in
	}
	if in.StreamARN != nil {
		in, out := &in.StreamARN, &out.StreamARN
		*out = new(string)
		**out = **in
	}
	if in.StreamCreationTimestamp != nil {
		in, out := &in.StreamCreationTimestamp, &out.StreamCreationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.StreamModeDetails != nil {
		in, out := &in.StreamModeDetails, &out.StreamModeDetails
		*out = new(StreamModeDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.StreamName != nil {
		in, out := &in.StreamName, &out.StreamName
		*out = new(string)
		**out = **in
	}
	if in.StreamStatus != nil {
		in, out := &in.StreamStatus, &out.StreamStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamDescriptionSummary.
func (in *StreamDescriptionSummary) DeepCopy() *StreamDescriptionSummary {
	if in == nil {
		return nil
	}
	out := new(StreamDescriptionSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamList) DeepCopyInto(out *StreamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Stream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamList.
func (in *StreamList) DeepCopy() *StreamList {
	if in == nil {
		return nil
	}
	out := new(StreamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StreamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamModeDetails) DeepCopyInto(out *StreamModeDetails) {
	*out = *in
	if in.StreamMode != nil {
		in, out := &in.StreamMode, &out.StreamMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamModeDetails.
func (in *StreamModeDetails) DeepCopy() *StreamModeDetails {
	if in == nil {
		return nil
	}
	out := new(StreamModeDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamSpec) DeepCopyInto(out *StreamSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(int64)
		**out = **in
	}
	if in.StreamModeDetails != nil {
		in, out := &in.StreamModeDetails, &out.StreamModeDetails
		*out = new(StreamModeDetails)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamSpec.
func (in *StreamSpec) DeepCopy() *StreamSpec {
	if in == nil {
		return nil
	}
	out := new(StreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamStatus) DeepCopyInto(out *StreamStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ConsumerCount != nil {
		in, out := &in.ConsumerCount, &out.ConsumerCount
		*out = new(int64)
		**out = **in
	}
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.EnhancedMonitoring != nil {
		in, out := &in.EnhancedMonitoring, &out.EnhancedMonitoring
		*out = make([]*EnhancedMetrics, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(EnhancedMetrics)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
	if in.OpenShardCount != nil {
		in, out := &in.OpenShardCount, &out.OpenShardCount
		*out = new(int64)
		**out = **in
	}
	if in.RetentionPeriodHours != nil {
		in, out := &in.RetentionPeriodHours, &out.RetentionPeriodHours
		*out = new(int64)
		**out = **in
	}
	if in.StreamCreationTimestamp != nil {
		in, out := &in.StreamCreationTimestamp, &out.StreamCreationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.StreamStatus != nil {
		in, out := &in.StreamStatus, &out.StreamStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamStatus.
func (in *StreamStatus) DeepCopy() *StreamStatus {
	if in == nil {
		return nil
	}
	out := new(StreamStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamSummary) DeepCopyInto(out *StreamSummary) {
	*out = *in
	if in.StreamARN != nil {
		in, out := &in.StreamARN, &out.StreamARN
		*out = new(string)
		**out = **in
	}
	if in.StreamCreationTimestamp != nil {
		in, out := &in.StreamCreationTimestamp, &out.StreamCreationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.StreamModeDetails != nil {
		in, out := &in.StreamModeDetails, &out.StreamModeDetails
		*out = new(StreamModeDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.StreamName != nil {
		in, out := &in.StreamName, &out.StreamName
		*out = new(string)
		**out = **in
	}
	if in.StreamStatus != nil {
		in, out := &in.StreamStatus, &out.StreamStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamSummary.
func (in *StreamSummary) DeepCopy() *StreamSummary {
	if in == nil {
		return nil
	}
	out := new(StreamSummary)
	in.DeepCopyInto(out)
	return out
}
