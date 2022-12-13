/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta1 "github.com/FrangipaneTeam/provider-flexibleengine/apis/ecs/v1beta1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AntiDDoS.
func (mg *AntiDDoS) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FloatingIPID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FloatingIPIDRef,
		Selector:     mg.Spec.ForProvider.FloatingIPIDSelector,
		To: reference.To{
			List:    &v1beta1.FloatingIpList{},
			Managed: &v1beta1.FloatingIp{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FloatingIPID")
	}
	mg.Spec.ForProvider.FloatingIPID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FloatingIPIDRef = rsp.ResolvedReference

	return nil
}