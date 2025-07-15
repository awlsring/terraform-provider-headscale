package service

import (
	"context"

	"github.com/awlsring/terraform-provider-headscale/internal/gen/client/headscale_service"
	"github.com/awlsring/terraform-provider-headscale/internal/gen/models"
	"github.com/go-openapi/strfmt"
)

func (h *HeadscaleService) ListRoutes(ctx context.Context) ([]*models.V1Route, error) {
	request := headscale_service.NewHeadscaleServiceGetRoutesParams()
	request.SetContext(ctx)

	resp, err := h.client.HeadscaleService.HeadscaleServiceGetRoutes(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Routes, nil
}

func (h *HeadscaleService) DeleteRoute(ctx context.Context, route string) error {
	request := headscale_service.NewHeadscaleServiceDeleteRouteParams()
	request.SetContext(ctx)
	request.SetRouteID(route)

	_, err := h.client.HeadscaleService.HeadscaleServiceDeleteRoute(request)
	if err != nil {
		return handleRequestError(err)
	}
	return nil
}

func (h *HeadscaleService) DisableRoute(ctx context.Context, route string) error {
	request := headscale_service.NewHeadscaleServiceDisableRouteParams()
	request.SetContext(ctx)
	request.SetRouteID(route)

	_, err := h.client.HeadscaleService.HeadscaleServiceDisableRoute(request)
	if err != nil {
		return handleRequestError(err)
	}

	return nil
}

func (h *HeadscaleService) EnableRoute(ctx context.Context, route string) error {
	request := headscale_service.NewHeadscaleServiceEnableRouteParams()
	request.SetContext(ctx)
	request.SetRouteID(route)

	_, err := h.client.HeadscaleService.HeadscaleServiceEnableRoute(request)
	if err != nil {
		return handleRequestError(err)
	}

	return nil
}
