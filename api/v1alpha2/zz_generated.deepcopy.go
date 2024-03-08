//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentDeployment) DeepCopyInto(out *AgentDeployment) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(v1.PodSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentDeployment.
func (in *AgentDeployment) DeepCopy() *AgentDeployment {
	if in == nil {
		return nil
	}
	out := new(AgentDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentDeploymentAutoscaling) DeepCopyInto(out *AgentDeploymentAutoscaling) {
	*out = *in
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.TargetWorkspaces != nil {
		in, out := &in.TargetWorkspaces, &out.TargetWorkspaces
		*out = new([]TargetWorkspace)
		if **in != nil {
			in, out := *in, *out
			*out = make([]TargetWorkspace, len(*in))
			copy(*out, *in)
		}
	}
	if in.CooldownPeriodSeconds != nil {
		in, out := &in.CooldownPeriodSeconds, &out.CooldownPeriodSeconds
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentDeploymentAutoscaling.
func (in *AgentDeploymentAutoscaling) DeepCopy() *AgentDeploymentAutoscaling {
	if in == nil {
		return nil
	}
	out := new(AgentDeploymentAutoscaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentDeploymentAutoscalingStatus) DeepCopyInto(out *AgentDeploymentAutoscalingStatus) {
	*out = *in
	if in.DesiredReplicas != nil {
		in, out := &in.DesiredReplicas, &out.DesiredReplicas
		*out = new(int32)
		**out = **in
	}
	if in.LastScalingEvent != nil {
		in, out := &in.LastScalingEvent, &out.LastScalingEvent
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentDeploymentAutoscalingStatus.
func (in *AgentDeploymentAutoscalingStatus) DeepCopy() *AgentDeploymentAutoscalingStatus {
	if in == nil {
		return nil
	}
	out := new(AgentDeploymentAutoscalingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentPool) DeepCopyInto(out *AgentPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentPool.
func (in *AgentPool) DeepCopy() *AgentPool {
	if in == nil {
		return nil
	}
	out := new(AgentPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentPoolList) DeepCopyInto(out *AgentPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AgentPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentPoolList.
func (in *AgentPoolList) DeepCopy() *AgentPoolList {
	if in == nil {
		return nil
	}
	out := new(AgentPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentPoolSpec) DeepCopyInto(out *AgentPoolSpec) {
	*out = *in
	in.Token.DeepCopyInto(&out.Token)
	if in.AgentTokens != nil {
		in, out := &in.AgentTokens, &out.AgentTokens
		*out = make([]*AgentToken, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AgentToken)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.AgentDeployment != nil {
		in, out := &in.AgentDeployment, &out.AgentDeployment
		*out = new(AgentDeployment)
		(*in).DeepCopyInto(*out)
	}
	if in.AgentDeploymentAutoscaling != nil {
		in, out := &in.AgentDeploymentAutoscaling, &out.AgentDeploymentAutoscaling
		*out = new(AgentDeploymentAutoscaling)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentPoolSpec.
func (in *AgentPoolSpec) DeepCopy() *AgentPoolSpec {
	if in == nil {
		return nil
	}
	out := new(AgentPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentPoolStatus) DeepCopyInto(out *AgentPoolStatus) {
	*out = *in
	if in.AgentTokens != nil {
		in, out := &in.AgentTokens, &out.AgentTokens
		*out = make([]*AgentToken, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AgentToken)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.AgentDeploymentAutoscalingStatus != nil {
		in, out := &in.AgentDeploymentAutoscalingStatus, &out.AgentDeploymentAutoscalingStatus
		*out = new(AgentDeploymentAutoscalingStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentPoolStatus.
func (in *AgentPoolStatus) DeepCopy() *AgentPoolStatus {
	if in == nil {
		return nil
	}
	out := new(AgentPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentToken) DeepCopyInto(out *AgentToken) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(int64)
		**out = **in
	}
	if in.LastUsedAt != nil {
		in, out := &in.LastUsedAt, &out.LastUsedAt
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentToken.
func (in *AgentToken) DeepCopy() *AgentToken {
	if in == nil {
		return nil
	}
	out := new(AgentToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationVersionStatus) DeepCopyInto(out *ConfigurationVersionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationVersionStatus.
func (in *ConfigurationVersionStatus) DeepCopy() *ConfigurationVersionStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigurationVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsumerWorkspace) DeepCopyInto(out *ConsumerWorkspace) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsumerWorkspace.
func (in *ConsumerWorkspace) DeepCopy() *ConsumerWorkspace {
	if in == nil {
		return nil
	}
	out := new(ConsumerWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomPermissions) DeepCopyInto(out *CustomPermissions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomPermissions.
func (in *CustomPermissions) DeepCopy() *CustomPermissions {
	if in == nil {
		return nil
	}
	out := new(CustomPermissions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomProjectPermissions) DeepCopyInto(out *CustomProjectPermissions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomProjectPermissions.
func (in *CustomProjectPermissions) DeepCopy() *CustomProjectPermissions {
	if in == nil {
		return nil
	}
	out := new(CustomProjectPermissions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module) DeepCopyInto(out *Module) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module.
func (in *Module) DeepCopy() *Module {
	if in == nil {
		return nil
	}
	out := new(Module)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Module) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleList) DeepCopyInto(out *ModuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Module, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleList.
func (in *ModuleList) DeepCopy() *ModuleList {
	if in == nil {
		return nil
	}
	out := new(ModuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleOutput) DeepCopyInto(out *ModuleOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleOutput.
func (in *ModuleOutput) DeepCopy() *ModuleOutput {
	if in == nil {
		return nil
	}
	out := new(ModuleOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleSource) DeepCopyInto(out *ModuleSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleSource.
func (in *ModuleSource) DeepCopy() *ModuleSource {
	if in == nil {
		return nil
	}
	out := new(ModuleSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleSpec) DeepCopyInto(out *ModuleSpec) {
	*out = *in
	in.Token.DeepCopyInto(&out.Token)
	if in.Module != nil {
		in, out := &in.Module, &out.Module
		*out = new(ModuleSource)
		**out = **in
	}
	if in.Workspace != nil {
		in, out := &in.Workspace, &out.Workspace
		*out = new(ModuleWorkspace)
		**out = **in
	}
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make([]ModuleVariable, len(*in))
		copy(*out, *in)
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make([]ModuleOutput, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleSpec.
func (in *ModuleSpec) DeepCopy() *ModuleSpec {
	if in == nil {
		return nil
	}
	out := new(ModuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleStatus) DeepCopyInto(out *ModuleStatus) {
	*out = *in
	if in.ConfigurationVersion != nil {
		in, out := &in.ConfigurationVersion, &out.ConfigurationVersion
		*out = new(ConfigurationVersionStatus)
		**out = **in
	}
	if in.Run != nil {
		in, out := &in.Run, &out.Run
		*out = new(RunStatus)
		**out = **in
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = new(OutputStatus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleStatus.
func (in *ModuleStatus) DeepCopy() *ModuleStatus {
	if in == nil {
		return nil
	}
	out := new(ModuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleVariable) DeepCopyInto(out *ModuleVariable) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleVariable.
func (in *ModuleVariable) DeepCopy() *ModuleVariable {
	if in == nil {
		return nil
	}
	out := new(ModuleVariable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleWorkspace) DeepCopyInto(out *ModuleWorkspace) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleWorkspace.
func (in *ModuleWorkspace) DeepCopy() *ModuleWorkspace {
	if in == nil {
		return nil
	}
	out := new(ModuleWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Notification) DeepCopyInto(out *Notification) {
	*out = *in
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]NotificationTrigger, len(*in))
		copy(*out, *in)
	}
	if in.EmailAddresses != nil {
		in, out := &in.EmailAddresses, &out.EmailAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EmailUsers != nil {
		in, out := &in.EmailUsers, &out.EmailUsers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Notification.
func (in *Notification) DeepCopy() *Notification {
	if in == nil {
		return nil
	}
	out := new(Notification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputStatus) DeepCopyInto(out *OutputStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputStatus.
func (in *OutputStatus) DeepCopy() *OutputStatus {
	if in == nil {
		return nil
	}
	out := new(OutputStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Project) DeepCopyInto(out *Project) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Project.
func (in *Project) DeepCopy() *Project {
	if in == nil {
		return nil
	}
	out := new(Project)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Project) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectList) DeepCopyInto(out *ProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Project, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectList.
func (in *ProjectList) DeepCopy() *ProjectList {
	if in == nil {
		return nil
	}
	out := new(ProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSpec) DeepCopyInto(out *ProjectSpec) {
	*out = *in
	in.Token.DeepCopyInto(&out.Token)
	if in.TeamAccess != nil {
		in, out := &in.TeamAccess, &out.TeamAccess
		*out = make([]*ProjectTeamAccess, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ProjectTeamAccess)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSpec.
func (in *ProjectSpec) DeepCopy() *ProjectSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectStatus) DeepCopyInto(out *ProjectStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectStatus.
func (in *ProjectStatus) DeepCopy() *ProjectStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectTeamAccess) DeepCopyInto(out *ProjectTeamAccess) {
	*out = *in
	out.Team = in.Team
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = new(CustomProjectPermissions)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectTeamAccess.
func (in *ProjectTeamAccess) DeepCopy() *ProjectTeamAccess {
	if in == nil {
		return nil
	}
	out := new(ProjectTeamAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteStateSharing) DeepCopyInto(out *RemoteStateSharing) {
	*out = *in
	if in.Workspaces != nil {
		in, out := &in.Workspaces, &out.Workspaces
		*out = make([]*ConsumerWorkspace, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ConsumerWorkspace)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteStateSharing.
func (in *RemoteStateSharing) DeepCopy() *RemoteStateSharing {
	if in == nil {
		return nil
	}
	out := new(RemoteStateSharing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunStatus) DeepCopyInto(out *RunStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunStatus.
func (in *RunStatus) DeepCopy() *RunStatus {
	if in == nil {
		return nil
	}
	out := new(RunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunTrigger) DeepCopyInto(out *RunTrigger) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunTrigger.
func (in *RunTrigger) DeepCopy() *RunTrigger {
	if in == nil {
		return nil
	}
	out := new(RunTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHKey) DeepCopyInto(out *SSHKey) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHKey.
func (in *SSHKey) DeepCopy() *SSHKey {
	if in == nil {
		return nil
	}
	out := new(SSHKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetWorkspace) DeepCopyInto(out *TargetWorkspace) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetWorkspace.
func (in *TargetWorkspace) DeepCopy() *TargetWorkspace {
	if in == nil {
		return nil
	}
	out := new(TargetWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Team) DeepCopyInto(out *Team) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Team.
func (in *Team) DeepCopy() *Team {
	if in == nil {
		return nil
	}
	out := new(Team)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamAccess) DeepCopyInto(out *TeamAccess) {
	*out = *in
	out.Team = in.Team
	out.Custom = in.Custom
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamAccess.
func (in *TeamAccess) DeepCopy() *TeamAccess {
	if in == nil {
		return nil
	}
	out := new(TeamAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Token) DeepCopyInto(out *Token) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Token.
func (in *Token) DeepCopy() *Token {
	if in == nil {
		return nil
	}
	out := new(Token)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueFrom) DeepCopyInto(out *ValueFrom) {
	*out = *in
	if in.ConfigMapKeyRef != nil {
		in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueFrom.
func (in *ValueFrom) DeepCopy() *ValueFrom {
	if in == nil {
		return nil
	}
	out := new(ValueFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Variable) DeepCopyInto(out *Variable) {
	*out = *in
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(ValueFrom)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Variable.
func (in *Variable) DeepCopy() *Variable {
	if in == nil {
		return nil
	}
	out := new(Variable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionControl) DeepCopyInto(out *VersionControl) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionControl.
func (in *VersionControl) DeepCopy() *VersionControl {
	if in == nil {
		return nil
	}
	out := new(VersionControl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workspace) DeepCopyInto(out *Workspace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspace.
func (in *Workspace) DeepCopy() *Workspace {
	if in == nil {
		return nil
	}
	out := new(Workspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workspace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceAgentPool) DeepCopyInto(out *WorkspaceAgentPool) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceAgentPool.
func (in *WorkspaceAgentPool) DeepCopy() *WorkspaceAgentPool {
	if in == nil {
		return nil
	}
	out := new(WorkspaceAgentPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceList) DeepCopyInto(out *WorkspaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceList.
func (in *WorkspaceList) DeepCopy() *WorkspaceList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceProject) DeepCopyInto(out *WorkspaceProject) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceProject.
func (in *WorkspaceProject) DeepCopy() *WorkspaceProject {
	if in == nil {
		return nil
	}
	out := new(WorkspaceProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceRunTask) DeepCopyInto(out *WorkspaceRunTask) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceRunTask.
func (in *WorkspaceRunTask) DeepCopy() *WorkspaceRunTask {
	if in == nil {
		return nil
	}
	out := new(WorkspaceRunTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSpec) DeepCopyInto(out *WorkspaceSpec) {
	*out = *in
	in.Token.DeepCopyInto(&out.Token)
	if in.AgentPool != nil {
		in, out := &in.AgentPool, &out.AgentPool
		*out = new(WorkspaceAgentPool)
		**out = **in
	}
	if in.RunTasks != nil {
		in, out := &in.RunTasks, &out.RunTasks
		*out = make([]WorkspaceRunTask, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
	if in.TeamAccess != nil {
		in, out := &in.TeamAccess, &out.TeamAccess
		*out = make([]*TeamAccess, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(TeamAccess)
				**out = **in
			}
		}
	}
	if in.EnvironmentVariables != nil {
		in, out := &in.EnvironmentVariables, &out.EnvironmentVariables
		*out = make([]Variable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TerraformVariables != nil {
		in, out := &in.TerraformVariables, &out.TerraformVariables
		*out = make([]Variable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RemoteStateSharing != nil {
		in, out := &in.RemoteStateSharing, &out.RemoteStateSharing
		*out = new(RemoteStateSharing)
		(*in).DeepCopyInto(*out)
	}
	if in.RunTriggers != nil {
		in, out := &in.RunTriggers, &out.RunTriggers
		*out = make([]RunTrigger, len(*in))
		copy(*out, *in)
	}
	if in.VersionControl != nil {
		in, out := &in.VersionControl, &out.VersionControl
		*out = new(VersionControl)
		**out = **in
	}
	if in.SSHKey != nil {
		in, out := &in.SSHKey, &out.SSHKey
		*out = new(SSHKey)
		**out = **in
	}
	if in.Notifications != nil {
		in, out := &in.Notifications, &out.Notifications
		*out = make([]Notification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(WorkspaceProject)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSpec.
func (in *WorkspaceSpec) DeepCopy() *WorkspaceSpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceStatus) DeepCopyInto(out *WorkspaceStatus) {
	*out = *in
	out.Run = in.Run
	out.DeleteRun = in.DeleteRun
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceStatus.
func (in *WorkspaceStatus) DeepCopy() *WorkspaceStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceStatus)
	in.DeepCopyInto(out)
	return out
}
