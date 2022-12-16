/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta1 "github.com/FrangipaneTeam/provider-flexibleengine/apis/iam/v1beta1"
	v1beta11 "github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this L7Policy.
func (mg *L7Policy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ListenerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ListenerIDRef,
		Selector:     mg.Spec.ForProvider.ListenerIDSelector,
		To: reference.To{
			List:    &ListenerList{},
			Managed: &Listener{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ListenerID")
	}
	mg.Spec.ForProvider.ListenerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ListenerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RedirectListenerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RedirectListenerIDRef,
		Selector:     mg.Spec.ForProvider.RedirectListenerIDSelector,
		To: reference.To{
			List:    &ListenerList{},
			Managed: &Listener{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RedirectListenerID")
	}
	mg.Spec.ForProvider.RedirectListenerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RedirectListenerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RedirectPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RedirectPoolIDRef,
		Selector:     mg.Spec.ForProvider.RedirectPoolIDSelector,
		To: reference.To{
			List:    &PoolList{},
			Managed: &Pool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RedirectPoolID")
	}
	mg.Spec.ForProvider.RedirectPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RedirectPoolIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1beta1.ProjectList{},
			Managed: &v1beta1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this L7Rule.
func (mg *L7Rule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.L7PolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.L7PolicyIDRef,
		Selector:     mg.Spec.ForProvider.L7PolicyIDSelector,
		To: reference.To{
			List:    &L7PolicyList{},
			Managed: &L7Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.L7PolicyID")
	}
	mg.Spec.ForProvider.L7PolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.L7PolicyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1beta1.ProjectList{},
			Managed: &v1beta1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Listener.
func (mg *Listener) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DefaultPoolIDRef,
		Selector:     mg.Spec.ForProvider.DefaultPoolIDSelector,
		To: reference.To{
			List:    &PoolList{},
			Managed: &Pool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DefaultPoolID")
	}
	mg.Spec.ForProvider.DefaultPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DefaultPoolIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadbalancerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LoadbalancerIDRef,
		Selector:     mg.Spec.ForProvider.LoadbalancerIDSelector,
		To: reference.To{
			List:    &LoadBalancerList{},
			Managed: &LoadBalancer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadbalancerID")
	}
	mg.Spec.ForProvider.LoadbalancerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadbalancerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1beta1.ProjectList{},
			Managed: &v1beta1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LoadBalancer.
func (mg *LoadBalancer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &v1beta11.SecurityGroupList{},
			Managed: &v1beta11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1beta1.ProjectList{},
			Managed: &v1beta1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VipSubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VipSubnetIDRef,
		Selector:     mg.Spec.ForProvider.VipSubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.VPCSubnetList{},
			Managed: &v1beta11.VPCSubnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VipSubnetID")
	}
	mg.Spec.ForProvider.VipSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VipSubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Member.
func (mg *Member) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PoolIDRef,
		Selector:     mg.Spec.ForProvider.PoolIDSelector,
		To: reference.To{
			List:    &PoolList{},
			Managed: &Pool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PoolID")
	}
	mg.Spec.ForProvider.PoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PoolIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.VPCSubnetList{},
			Managed: &v1beta11.VPCSubnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1beta1.ProjectList{},
			Managed: &v1beta1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Monitor.
func (mg *Monitor) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PoolIDRef,
		Selector:     mg.Spec.ForProvider.PoolIDSelector,
		To: reference.To{
			List:    &PoolList{},
			Managed: &Pool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PoolID")
	}
	mg.Spec.ForProvider.PoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PoolIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1beta1.ProjectList{},
			Managed: &v1beta1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Pool.
func (mg *Pool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ListenerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ListenerIDRef,
		Selector:     mg.Spec.ForProvider.ListenerIDSelector,
		To: reference.To{
			List:    &ListenerList{},
			Managed: &Listener{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ListenerID")
	}
	mg.Spec.ForProvider.ListenerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ListenerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadbalancerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LoadbalancerIDRef,
		Selector:     mg.Spec.ForProvider.LoadbalancerIDSelector,
		To: reference.To{
			List:    &LoadBalancerList{},
			Managed: &LoadBalancer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadbalancerID")
	}
	mg.Spec.ForProvider.LoadbalancerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadbalancerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1beta1.ProjectList{},
			Managed: &v1beta1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Whitelist.
func (mg *Whitelist) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ListenerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ListenerIDRef,
		Selector:     mg.Spec.ForProvider.ListenerIDSelector,
		To: reference.To{
			List:    &ListenerList{},
			Managed: &Listener{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ListenerID")
	}
	mg.Spec.ForProvider.ListenerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ListenerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1beta1.ProjectList{},
			Managed: &v1beta1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}
