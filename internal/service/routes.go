package service

import (
	"context"
	"slices"

	"github.com/awlsring/terraform-provider-headscale/internal/gen/client/headscale_service"
	"github.com/awlsring/terraform-provider-headscale/internal/gen/models"
)

type Route struct {
	ID      string
	Prefix  string
	Enabled bool
	Node    *models.V1Node
}

func (h *HeadscaleService) ListRoutes(ctx context.Context) ([]*Route, error) {
	nodes, err := h.ListDevices(ctx, nil)
	if err != nil {
		return nil, err
	}
	var routes []*Route
	for _, node := range nodes {
		nodeRoutes, err := h.ListDeviceRoutes(ctx, node.ID)
		if err != nil {
			return nil, err
		}
		routes = append(routes, nodeRoutes...)
	}
	return routes, nil
}

func (h *HeadscaleService) ListDeviceRoutes(ctx context.Context, deviceId string) ([]*Route, error) {
	device, err := h.GetDevice(ctx, deviceId)
	if err != nil {
		return nil, err
	}

	var routes []*Route

	for _, prefix := range device.AvailableRoutes {
		routes = append(routes, &Route{
			ID:      prefix,
			Prefix:  prefix,
			Enabled: slices.Contains(device.ApprovedRoutes, prefix),
			Node:    device,
		})
	}

	return routes, nil
}

func (h *HeadscaleService) EnableDeviceRoutes(ctx context.Context, deviceId string, routes []string) error {
	request := headscale_service.NewHeadscaleServiceSetApprovedRoutesParams().
		WithContext(ctx).
		WithNodeID(deviceId).
		WithBody(&models.HeadscaleServiceSetApprovedRoutesBody{
			Routes: routes,
		})

	_, err := h.client.HeadscaleService.HeadscaleServiceSetApprovedRoutes(request)
	if err != nil {
		return handleRequestError(err)
	}

	return nil
}

func (h *HeadscaleService) DisableDeviceRoutes(ctx context.Context, deviceId string) error {
	request := headscale_service.NewHeadscaleServiceSetApprovedRoutesParams().
		WithContext(ctx).
		WithNodeID(deviceId).
		WithBody(&models.HeadscaleServiceSetApprovedRoutesBody{
			Routes: []string{},
		})

	_, err := h.client.HeadscaleService.HeadscaleServiceSetApprovedRoutes(request)
	if err != nil {
		return handleRequestError(err)
	}

	return nil
}
