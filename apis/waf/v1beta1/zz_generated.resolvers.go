/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta1 "github.com/FrangipaneTeam/provider-flexibleengine/apis/ecs/v1beta1"
	v1beta11 "github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1"
	common "github.com/FrangipaneTeam/provider-flexibleengine/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this DedicatedDomain.
func (mg *DedicatedDomain) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CertificateID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CertificateIDRef,
		Selector:     mg.Spec.ForProvider.CertificateIDSelector,
		To: reference.To{
			List:    &DedicatedCertificateList{},
			Managed: &DedicatedCertificate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CertificateID")
	}
	mg.Spec.ForProvider.CertificateID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CertificateIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PolicyIDRef,
		Selector:     mg.Spec.ForProvider.PolicyIDSelector,
		To: reference.To{
			List:    &DedicatedPolicyList{},
			Managed: &DedicatedPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyID")
	}
	mg.Spec.ForProvider.PolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DedicatedInstance.
func (mg *DedicatedInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.GroupIDRef,
		Selector:     mg.Spec.ForProvider.GroupIDSelector,
		To: reference.To{
			List:    &v1beta1.ServerGroupList{},
			Managed: &v1beta1.ServerGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupID")
	}
	mg.Spec.ForProvider.GroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroup),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupSelector,
		To: reference.To{
			List:    &v1beta11.SecurityGroupList{},
			Managed: &v1beta11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroup")
	}
	mg.Spec.ForProvider.SecurityGroup = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      common.IDExtractor(),
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
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1beta11.VPCList{},
			Managed: &v1beta11.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Domain.
func (mg *Domain) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CertificateID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CertificateIDRef,
		Selector:     mg.Spec.ForProvider.CertificateIDSelector,
		To: reference.To{
			List:    &CertificateList{},
			Managed: &Certificate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CertificateID")
	}
	mg.Spec.ForProvider.CertificateID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CertificateIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PolicyIDRef,
		Selector:     mg.Spec.ForProvider.PolicyIDSelector,
		To: reference.To{
			List:    &PolicyList{},
			Managed: &Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyID")
	}
	mg.Spec.ForProvider.PolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RuleAlarmMasking.
func (mg *RuleAlarmMasking) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PolicyIDRef,
		Selector:     mg.Spec.ForProvider.PolicyIDSelector,
		To: reference.To{
			List:    &PolicyList{},
			Managed: &Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyID")
	}
	mg.Spec.ForProvider.PolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RuleBlacklist.
func (mg *RuleBlacklist) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PolicyIDRef,
		Selector:     mg.Spec.ForProvider.PolicyIDSelector,
		To: reference.To{
			List:    &PolicyList{},
			Managed: &Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyID")
	}
	mg.Spec.ForProvider.PolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RuleCcProtection.
func (mg *RuleCcProtection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PolicyIDRef,
		Selector:     mg.Spec.ForProvider.PolicyIDSelector,
		To: reference.To{
			List:    &PolicyList{},
			Managed: &Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyID")
	}
	mg.Spec.ForProvider.PolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RuleDataMasking.
func (mg *RuleDataMasking) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PolicyIDRef,
		Selector:     mg.Spec.ForProvider.PolicyIDSelector,
		To: reference.To{
			List:    &PolicyList{},
			Managed: &Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyID")
	}
	mg.Spec.ForProvider.PolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RulePreciseProtection.
func (mg *RulePreciseProtection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PolicyIDRef,
		Selector:     mg.Spec.ForProvider.PolicyIDSelector,
		To: reference.To{
			List:    &PolicyList{},
			Managed: &Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyID")
	}
	mg.Spec.ForProvider.PolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RuleWebTamperProtection.
func (mg *RuleWebTamperProtection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PolicyIDRef,
		Selector:     mg.Spec.ForProvider.PolicyIDSelector,
		To: reference.To{
			List:    &PolicyList{},
			Managed: &Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyID")
	}
	mg.Spec.ForProvider.PolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyIDRef = rsp.ResolvedReference

	return nil
}
