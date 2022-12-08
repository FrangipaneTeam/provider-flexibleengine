/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this EIPAssociate.
func (mg *EIPAssociate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PortIDRef,
		Selector:     mg.Spec.ForProvider.PortIDSelector,
		To: reference.To{
			List:    &NetworkPortList{},
			Managed: &NetworkPort{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortID")
	}
	mg.Spec.ForProvider.PortID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Endpoint.
func (mg *Endpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceIDRef,
		Selector:     mg.Spec.ForProvider.ServiceIDSelector,
		To: reference.To{
			List:    &EndpointServiceList{},
			Managed: &EndpointService{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceID")
	}
	mg.Spec.ForProvider.ServiceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &VpcList{},
			Managed: &Vpc{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EndpointApproval.
func (mg *EndpointApproval) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceIDRef,
		Selector:     mg.Spec.ForProvider.ServiceIDSelector,
		To: reference.To{
			List:    &EndpointServiceList{},
			Managed: &EndpointService{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceID")
	}
	mg.Spec.ForProvider.ServiceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EndpointService.
func (mg *EndpointService) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PortIDRef,
		Selector:     mg.Spec.ForProvider.PortIDSelector,
		To: reference.To{
			List:    &NetworkPortList{},
			Managed: &NetworkPort{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortID")
	}
	mg.Spec.ForProvider.PortID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &VpcList{},
			Managed: &Vpc{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NetworkPort.
func (mg *NetworkPort) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NetworkSubnet.
func (mg *NetworkSubnet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Peering.
func (mg *Peering) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &VpcList{},
			Managed: &Vpc{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Route.
func (mg *Route) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RouteTableID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RouteTableIDRef,
		Selector:     mg.Spec.ForProvider.RouteTableIDSelector,
		To: reference.To{
			List:    &RouteTableList{},
			Managed: &RouteTable{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RouteTableID")
	}
	mg.Spec.ForProvider.RouteTableID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RouteTableIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RouteTable.
func (mg *RouteTable) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Subnets),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetsRefs,
		Selector:      mg.Spec.ForProvider.SubnetsSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Subnets")
	}
	mg.Spec.ForProvider.Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &VpcList{},
			Managed: &Vpc{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RouterInterface.
func (mg *RouterInterface) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PortIDRef,
		Selector:     mg.Spec.ForProvider.PortIDSelector,
		To: reference.To{
			List:    &NetworkPortList{},
			Managed: &NetworkPort{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortID")
	}
	mg.Spec.ForProvider.PortID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RouterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RouterIDRef,
		Selector:     mg.Spec.ForProvider.RouterIDSelector,
		To: reference.To{
			List:    &RouterList{},
			Managed: &Router{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RouterID")
	}
	mg.Spec.ForProvider.RouterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RouterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityGroupRule.
func (mg *SecurityGroupRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RemoteGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RemoteGroupIDRef,
		Selector:     mg.Spec.ForProvider.RemoteGroupIDSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RemoteGroupID")
	}
	mg.Spec.ForProvider.RemoteGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RemoteGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecurityGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupID")
	}
	mg.Spec.ForProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecurityGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Subnet.
func (mg *Subnet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &VpcList{},
			Managed: &Vpc{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Vip.
func (mg *Vip) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VipAssociate.
func (mg *VipAssociate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.PortIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.PortIdsRefs,
		Selector:      mg.Spec.ForProvider.PortIdsSelector,
		To: reference.To{
			List:    &NetworkPortList{},
			Managed: &NetworkPort{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortIds")
	}
	mg.Spec.ForProvider.PortIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.PortIdsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VipID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VipIDRef,
		Selector:     mg.Spec.ForProvider.VipIDSelector,
		To: reference.To{
			List:    &VipList{},
			Managed: &Vip{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VipID")
	}
	mg.Spec.ForProvider.VipID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VipIDRef = rsp.ResolvedReference

	return nil
}