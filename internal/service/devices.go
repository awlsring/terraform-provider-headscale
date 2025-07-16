package service

import (
	"context"

	"github.com/awlsring/terraform-provider-headscale/internal/gen/client/headscale_service"
	"github.com/awlsring/terraform-provider-headscale/internal/gen/models"
	"github.com/go-openapi/strfmt"
)

func (h *HeadscaleService) ListDevices(ctx context.Context, user *string) ([]*models.V1Node, error) {
	request := headscale_service.NewHeadscaleServiceListNodesParams().
		WithContext(ctx).
		WithUser(user)

	resp, err := h.client.HeadscaleService.HeadscaleServiceListNodes(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Nodes, nil
}

func (h *HeadscaleService) GetDevice(ctx context.Context, deviceId string) (*models.V1Node, error) {
	request := headscale_service.NewHeadscaleServiceGetNodeParams()
	request.SetContext(ctx)
	request.SetNodeID(deviceId)

	resp, err := h.client.HeadscaleService.HeadscaleServiceGetNode(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Node, nil
}

func (h *HeadscaleService) CreateDevice(ctx context.Context, user string, key string) (*models.V1Node, error) {
	request := headscale_service.NewHeadscaleServiceRegisterNodeParams()
	request.SetContext(ctx)
	request.SetKey(&key)
	request.SetUser(&user)

	resp, err := h.client.HeadscaleService.HeadscaleServiceRegisterNode(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Node, nil
}

func (h *HeadscaleService) ExpireDevice(ctx context.Context, deviceId string) (*models.V1Node, error) {
	request := headscale_service.NewHeadscaleServiceExpireNodeParams()
	request.SetContext(ctx)
	request.SetNodeID(deviceId)

	resp, err := h.client.HeadscaleService.HeadscaleServiceExpireNode(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Node, nil
}

func (h *HeadscaleService) DeleteDevice(ctx context.Context, deviceId string) error {
	request := headscale_service.NewHeadscaleServiceDeleteNodeParams()
	request.SetContext(ctx)
	request.SetNodeID(deviceId)

	_, err := h.client.HeadscaleService.HeadscaleServiceDeleteNode(request)
	if err != nil {
		return handleRequestError(err)
	}
	return nil
}

func (h *HeadscaleService) RenameDevice(ctx context.Context, deviceId string, name string) (*models.V1Node, error) {
	request := headscale_service.NewHeadscaleServiceRenameNodeParams()
	request.SetContext(ctx)
	request.SetNodeID(deviceId)
	request.SetNewName(name)

	resp, err := h.client.HeadscaleService.HeadscaleServiceRenameNode(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Node, nil
}

func (h *HeadscaleService) TagDevice(ctx context.Context, deviceId string, tags []string) (*models.V1Node, error) {
	request := headscale_service.NewHeadscaleServiceSetTagsParams()
	request.SetContext(ctx)
	request.SetNodeID(deviceId)

	request.SetBody(&models.HeadscaleServiceSetTagsBody{
		Tags: tags,
	})

	resp, err := h.client.HeadscaleService.HeadscaleServiceSetTags(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Node, nil
}

func (h *HeadscaleService) MoveDevice(ctx context.Context, deviceId string, user string) (*models.V1Node, error) {
	request := headscale_service.NewHeadscaleServiceMoveNodeParams()
	request.SetContext(ctx)
	request.SetNodeID(deviceId)
	request.SetBody(&models.HeadscaleServiceMoveNodeBody{
		User: user,
	})

	resp, err := h.client.HeadscaleService.HeadscaleServiceMoveNode(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Node, nil
}
