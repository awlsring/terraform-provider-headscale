package service

import (
	"context"

	"github.com/awlsring/terraform-provider-headscale/internal/gen/client/headscale_service"
	"github.com/awlsring/terraform-provider-headscale/internal/gen/models"
	"github.com/go-openapi/strfmt"
)

func (h *HeadscaleService) ListAPIKeys(ctx context.Context) ([]*models.V1APIKey, error) {
	request := headscale_service.NewHeadscaleServiceListAPIKeysParams()
	request.SetContext(ctx)

	resp, err := h.client.HeadscaleService.HeadscaleServiceListAPIKeys(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.APIKeys, nil
}

func (h *HeadscaleService) CreateAPIKey(ctx context.Context, expiration *strfmt.DateTime) (string, error) {
	request := headscale_service.NewHeadscaleServiceCreateAPIKeyParams()
	request.SetContext(ctx)
	if expiration != nil {
		request.SetBody(&models.V1CreateAPIKeyRequest{
			Expiration: *expiration,
		})
	}

	resp, err := h.client.HeadscaleService.HeadscaleServiceCreateAPIKey(request)
	if err != nil {
		return "", handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return "", err
	}

	return resp.Payload.APIKey, nil
}

func (h *HeadscaleService) ExpireAPIKey(ctx context.Context, key string) error {
	request := headscale_service.NewHeadscaleServiceExpireAPIKeyParams()
	request.SetContext(ctx)
	request.SetBody(&models.V1ExpireAPIKeyRequest{
		Prefix: key,
	})
	_, err := h.client.HeadscaleService.HeadscaleServiceExpireAPIKey(request)
	if err != nil {
		return handleRequestError(err)
	}
	return nil
}
