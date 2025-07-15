package service

import (
	"context"
	"strings"

	"github.com/awlsring/terraform-provider-headscale/internal/gen/client/headscale_service"
	"github.com/awlsring/terraform-provider-headscale/internal/gen/models"
	"github.com/go-openapi/strfmt"
)

func (h *HeadscaleService) ListPreAuthKeys(ctx context.Context, user string) ([]*models.V1PreAuthKey, error) {
	request := headscale_service.NewHeadscaleServiceListPreAuthKeysParams()
	request.SetContext(ctx)
	request.SetUser(&user)

	resp, err := h.client.HeadscaleService.HeadscaleServiceListPreAuthKeys(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.PreAuthKeys, nil
}

type CreatePreAuthKeyInput struct {
	User       string
	Reusable   bool
	Ephemeral  bool
	Expiration *strfmt.DateTime
	ACLTags    []string
}

func (h *HeadscaleService) CreatePreAuthKey(ctx context.Context, input CreatePreAuthKeyInput) (*models.V1PreAuthKey, error) {
	request := headscale_service.NewHeadscaleServiceCreatePreAuthKeyParams()
	request.SetContext(ctx)
	body := &models.V1CreatePreAuthKeyRequest{
		User:      input.User,
		Reusable:  input.Reusable,
		Ephemeral: input.Ephemeral,
		ACLTags:   input.ACLTags,
	}

	if input.Expiration != nil {
		body.Expiration = *input.Expiration
	}
	request.SetBody(body)

	resp, err := h.client.HeadscaleService.HeadscaleServiceCreatePreAuthKey(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.PreAuthKey, nil
}

func (h *HeadscaleService) ExpirePreAuthKey(ctx context.Context, user string, key string) error {
	request := headscale_service.NewHeadscaleServiceExpirePreAuthKeyParams()
	request.SetContext(ctx)
	request.SetBody(&models.V1ExpirePreAuthKeyRequest{
		User: user,
		Key:  key,
	})

	_, err := h.client.HeadscaleService.HeadscaleServiceExpirePreAuthKey(request)
	if err != nil {
		if e, ok := err.(*headscale_service.HeadscaleServiceExpirePreAuthKeyDefault); ok {
			if strings.Contains(e.Payload.Message, "AuthKey expired") {
				return nil
			}
		}
		if e, ok := err.(*headscale_service.HeadscaleServiceExpirePreAuthKeyDefault); ok {
			if strings.Contains(e.Payload.Message, "AuthKey has already been used") {
				return nil
			}
		}
		return handleRequestError(err)
	}
	return nil
}
