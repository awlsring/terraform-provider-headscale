package service

import (
	"context"

	"github.com/awlsring/terraform-provider-headscale/internal/gen/client/headscale_service"
	"github.com/awlsring/terraform-provider-headscale/internal/gen/models"
	"github.com/go-openapi/strfmt"
)

type Policy struct {
	Policy  string
	Updated strfmt.DateTime
}

func (h *HeadscaleService) SetPolicy(ctx context.Context, policyData string) (*Policy, error) {
	request := headscale_service.NewHeadscaleServiceSetPolicyParams().
		WithContext(ctx).
		WithBody(&models.V1SetPolicyRequest{
			Policy: policyData,
		})

	resp, err := h.client.HeadscaleService.HeadscaleServiceSetPolicy(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return &Policy{
		Policy:  resp.Payload.Policy,
		Updated: resp.Payload.UpdatedAt,
	}, nil
}

func (h *HeadscaleService) GetPolicy(ctx context.Context) (*Policy, error) {
	request := headscale_service.NewHeadscaleServiceGetPolicyParams().
		WithContext(ctx)

	resp, err := h.client.HeadscaleService.HeadscaleServiceGetPolicy(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return &Policy{
		Policy:  resp.Payload.Policy,
		Updated: resp.Payload.UpdatedAt,
	}, nil
}
