// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventConnection) DeepCopyInto(out *EventConnection) {
	*out = *in
	in.From.DeepCopyInto(&out.From)
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = make([]EventDestinationEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventConnection.
func (in *EventConnection) DeepCopy() *EventConnection {
	if in == nil {
		return nil
	}
	out := new(EventConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventConnections) DeepCopyInto(out *EventConnections) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventConnections.
func (in *EventConnections) DeepCopy() *EventConnections {
	if in == nil {
		return nil
	}
	out := new(EventConnections)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventConnections) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventConnectionsList) DeepCopyInto(out *EventConnectionsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventConnections, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventConnectionsList.
func (in *EventConnectionsList) DeepCopy() *EventConnectionsList {
	if in == nil {
		return nil
	}
	out := new(EventConnectionsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventConnectionsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventConnectionsSpec) DeepCopyInto(out *EventConnectionsSpec) {
	*out = *in
	if in.Connections != nil {
		in, out := &in.Connections, &out.Connections
		*out = make([]EventConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventConnectionsSpec.
func (in *EventConnectionsSpec) DeepCopy() *EventConnectionsSpec {
	if in == nil {
		return nil
	}
	out := new(EventConnectionsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventConnectionsStatus) DeepCopyInto(out *EventConnectionsStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventConnectionsStatus.
func (in *EventConnectionsStatus) DeepCopy() *EventConnectionsStatus {
	if in == nil {
		return nil
	}
	out := new(EventConnectionsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventDestinationEndpoint) DeepCopyInto(out *EventDestinationEndpoint) {
	*out = *in
	if in.Https != nil {
		in, out := &in.Https, &out.Https
		*out = new([]HttpsEndpoint)
		if **in != nil {
			in, out := *in, *out
			*out = make([]HttpsEndpoint, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventDestinationEndpoint.
func (in *EventDestinationEndpoint) DeepCopy() *EventDestinationEndpoint {
	if in == nil {
		return nil
	}
	out := new(EventDestinationEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventFunctionImpl) DeepCopyInto(out *EventFunctionImpl) {
	*out = *in
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = make([]EventStatement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventFunctionImpl.
func (in *EventFunctionImpl) DeepCopy() *EventFunctionImpl {
	if in == nil {
		return nil
	}
	out := new(EventFunctionImpl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventGithubRepository) DeepCopyInto(out *EventGithubRepository) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventGithubRepository.
func (in *EventGithubRepository) DeepCopy() *EventGithubRepository {
	if in == nil {
		return nil
	}
	out := new(EventGithubRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventMediationImpl) DeepCopyInto(out *EventMediationImpl) {
	*out = *in
	if in.SendTo != nil {
		in, out := &in.SendTo, &out.SendTo
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(EventMediationSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = new([]EventMediationVariable)
		if **in != nil {
			in, out := *in, *out
			*out = make([]EventMediationVariable, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = make([]EventStatement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventMediationImpl.
func (in *EventMediationImpl) DeepCopy() *EventMediationImpl {
	if in == nil {
		return nil
	}
	out := new(EventMediationImpl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventMediationRepositoryType) DeepCopyInto(out *EventMediationRepositoryType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventMediationRepositoryType.
func (in *EventMediationRepositoryType) DeepCopy() *EventMediationRepositoryType {
	if in == nil {
		return nil
	}
	out := new(EventMediationRepositoryType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventMediationSelector) DeepCopyInto(out *EventMediationSelector) {
	*out = *in
	if in.RepositoryType != nil {
		in, out := &in.RepositoryType, &out.RepositoryType
		*out = new(EventMediationRepositoryType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventMediationSelector.
func (in *EventMediationSelector) DeepCopy() *EventMediationSelector {
	if in == nil {
		return nil
	}
	out := new(EventMediationSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventMediationVariable) DeepCopyInto(out *EventMediationVariable) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.ValueExpression != nil {
		in, out := &in.ValueExpression, &out.ValueExpression
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventMediationVariable.
func (in *EventMediationVariable) DeepCopy() *EventMediationVariable {
	if in == nil {
		return nil
	}
	out := new(EventMediationVariable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventMediator) DeepCopyInto(out *EventMediator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventMediator.
func (in *EventMediator) DeepCopy() *EventMediator {
	if in == nil {
		return nil
	}
	out := new(EventMediator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventMediator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventMediatorList) DeepCopyInto(out *EventMediatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventMediator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventMediatorList.
func (in *EventMediatorList) DeepCopy() *EventMediatorList {
	if in == nil {
		return nil
	}
	out := new(EventMediatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventMediatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventMediatorSourceEndpoint) DeepCopyInto(out *EventMediatorSourceEndpoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventMediatorSourceEndpoint.
func (in *EventMediatorSourceEndpoint) DeepCopy() *EventMediatorSourceEndpoint {
	if in == nil {
		return nil
	}
	out := new(EventMediatorSourceEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventMediatorSpec) DeepCopyInto(out *EventMediatorSpec) {
	*out = *in
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = new([]EventRepository)
		if **in != nil {
			in, out := *in, *out
			*out = make([]EventRepository, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Mediations != nil {
		in, out := &in.Mediations, &out.Mediations
		*out = new([]EventMediationImpl)
		if **in != nil {
			in, out := *in, *out
			*out = make([]EventMediationImpl, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventMediatorSpec.
func (in *EventMediatorSpec) DeepCopy() *EventMediatorSpec {
	if in == nil {
		return nil
	}
	out := new(EventMediatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventMediatorStatus) DeepCopyInto(out *EventMediatorStatus) {
	*out = *in
	if in.UnorderedSummary != nil {
		in, out := &in.UnorderedSummary, &out.UnorderedSummary
		*out = make([]EventStatusSummary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventMediatorStatus.
func (in *EventMediatorStatus) DeepCopy() *EventMediatorStatus {
	if in == nil {
		return nil
	}
	out := new(EventMediatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventRepository) DeepCopyInto(out *EventRepository) {
	*out = *in
	if in.Github != nil {
		in, out := &in.Github, &out.Github
		*out = new(EventGithubRepository)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventRepository.
func (in *EventRepository) DeepCopy() *EventRepository {
	if in == nil {
		return nil
	}
	out := new(EventRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceEndpoint) DeepCopyInto(out *EventSourceEndpoint) {
	*out = *in
	if in.Mediator != nil {
		in, out := &in.Mediator, &out.Mediator
		*out = new(EventMediatorSourceEndpoint)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceEndpoint.
func (in *EventSourceEndpoint) DeepCopy() *EventSourceEndpoint {
	if in == nil {
		return nil
	}
	out := new(EventSourceEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventStatement) DeepCopyInto(out *EventStatement) {
	*out = *in
	if in.If != nil {
		in, out := &in.If, &out.If
		*out = new(string)
		**out = **in
	}
	if in.Assign != nil {
		in, out := &in.Assign, &out.Assign
		*out = new(string)
		**out = **in
	}
	if in.Switch != nil {
		in, out := &in.Switch, &out.Switch
		*out = new([]EventStatement)
		if **in != nil {
			in, out := *in, *out
			*out = make([]EventStatement, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = new([]EventStatement)
		if **in != nil {
			in, out := *in, *out
			*out = make([]EventStatement, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new([]EventStatement)
		if **in != nil {
			in, out := *in, *out
			*out = make([]EventStatement, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventStatement.
func (in *EventStatement) DeepCopy() *EventStatement {
	if in == nil {
		return nil
	}
	out := new(EventStatement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventStatusParameter) DeepCopyInto(out *EventStatusParameter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventStatusParameter.
func (in *EventStatusParameter) DeepCopy() *EventStatusParameter {
	if in == nil {
		return nil
	}
	out := new(EventStatusParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventStatusSummary) DeepCopyInto(out *EventStatusSummary) {
	*out = *in
	if in.Input != nil {
		in, out := &in.Input, &out.Input
		*out = make([]EventStatusParameter, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventStatusSummary.
func (in *EventStatusSummary) DeepCopy() *EventStatusSummary {
	if in == nil {
		return nil
	}
	out := new(EventStatusSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpsEndpoint) DeepCopyInto(out *HttpsEndpoint) {
	*out = *in
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
	if in.UrlExpression != nil {
		in, out := &in.UrlExpression, &out.UrlExpression
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpsEndpoint.
func (in *HttpsEndpoint) DeepCopy() *HttpsEndpoint {
	if in == nil {
		return nil
	}
	out := new(HttpsEndpoint)
	in.DeepCopyInto(out)
	return out
}
