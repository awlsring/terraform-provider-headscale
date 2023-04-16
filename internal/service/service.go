package service

import (
	"context"
	"net/url"

	"github.com/awlsring/terraform-provider-headscale/internal/client"
	"github.com/awlsring/terraform-provider-headscale/internal/client/headscale_service"
	"github.com/awlsring/terraform-provider-headscale/internal/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Headscale interface {
	ListAPIKeys(ctx context.Context) (*models.V1ListAPIKeysResponse, error)
	CreateAPIKey(ctx context.Context, expiration *strfmt.DateTime) (*models.V1CreateAPIKeyResponse, error)
	ExpireAPIKey(ctx context.Context, key string) error
	ListDevices(ctx context.Context, user *string) ([]*models.V1Machine, error)
	GetDevice(ctx context.Context, deviceId string) (*models.V1Machine, error)
	CreateDevice(ctx context.Context, input CreateDeviceInput) (*models.V1RegisterMachineResponse, error)
	ExpireDevice(ctx context.Context, deviceId string) (*models.V1ExpireMachineResponse, error)
	DeleteDevice(ctx context.Context, deviceId string) error
	RenameDevice(ctx context.Context, deviceId string, newName string) (*models.V1RenameMachineResponse, error)
	GetDeviceRoutes(ctx context.Context, deviceId string) (*models.V1GetMachineRoutesResponse, error)
	TagDevice(ctx context.Context, deviceId string, tags []string) (*models.V1SetTagsResponse, error)
	MoveDevice(ctx context.Context, deviceId string, newOwner string) (*models.V1MoveMachineResponse, error)
	ListPreAuthKeys(ctx context.Context, user *string) (*models.V1ListPreAuthKeysResponse, error)
	CreatePreAuthKey(ctx context.Context, input CreatePreAuthKeyInput) (*models.V1CreatePreAuthKeyResponse, error)
	ExpirePreAuthKey(ctx context.Context, user string, key string) error
	ListRoutes(ctx context.Context) (*models.V1GetRoutesResponse, error)
	DeleteRoute(ctx context.Context, routeId string) error
	DisableRoute(ctx context.Context, routeId string) error
	EnableRoute(ctx context.Context, routeId string) error
	GetUser(ctx context.Context, userId string) (*models.V1GetUserResponse, error)
	ListUsers(ctx context.Context) (*models.V1ListUsersResponse, error)
	CreateUser(ctx context.Context, name string) (*models.V1CreateUserResponse, error)
	DeleteUser(ctx context.Context, userId string) error
	RenameUser(ctx context.Context, oldName string, newName string) error
}

type HeadscaleService struct {
	client *client.Headscale
}

type ClientConfig struct {
	Endpoint string
	Token    string
}

func New(c ClientConfig) (Headscale, error) {
	u, err := url.Parse(c.Endpoint)
	if err != nil {
		return nil, err
	}
	transport := httptransport.New(u.Host, u.Path, []string{u.Scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(c.Token)

	client := client.New(transport, strfmt.Default)

	return &HeadscaleService{
		client: client,
	}, nil
}

func (h *HeadscaleService) ListAPIKeys(ctx context.Context) (*models.V1ListAPIKeysResponse, error) {
	request := headscale_service.NewHeadscaleServiceListAPIKeysParams()
	request.SetContext(ctx)

	resp, err := h.client.HeadscaleService.HeadscaleServiceListAPIKeys(request)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *HeadscaleService) CreateAPIKey(ctx context.Context, expiration *strfmt.DateTime) (*models.V1CreateAPIKeyResponse, error) {
	request := headscale_service.NewHeadscaleServiceCreateAPIKeyParams()
	request.SetContext(ctx)
	if expiration != nil {
		request.SetBody(&models.V1CreateAPIKeyRequest{
			Expiration: *expiration,
		})
	}

	resp, err := h.client.HeadscaleService.HeadscaleServiceCreateAPIKey(request)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *HeadscaleService) ExpireAPIKey(ctx context.Context, key string) error {
	request := headscale_service.NewHeadscaleServiceExpireAPIKeyParams()
	request.SetContext(ctx)
	request.SetBody(&models.V1ExpireAPIKeyRequest{
		Prefix: key,
	})
	_, err := h.client.HeadscaleService.HeadscaleServiceExpireAPIKey(request)
	if err != nil {
		return err
	}
	return nil
}

func (h *HeadscaleService) ListDevices(ctx context.Context, user *string) ([]*models.V1Machine, error) {
	request := headscale_service.NewHeadscaleServiceListMachinesParams()
	request.SetContext(ctx)
	request.SetUser(user)

	resp, err := h.client.HeadscaleService.HeadscaleServiceListMachines(request)
	if err != nil {
		return nil, err
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Machines, nil
}

func (h *HeadscaleService) GetDevice(ctx context.Context, deviceId string) (*models.V1Machine, error) {
	request := headscale_service.NewHeadscaleServiceGetMachineParams()
	request.SetContext(ctx)
	request.SetMachineID(deviceId)

	resp, err := h.client.HeadscaleService.HeadscaleServiceGetMachine(request)
	if err != nil {
		return nil, err
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Machine, nil
}

type CreateDeviceInput struct {
	User *string
	Key  *string
}

func (h *HeadscaleService) CreateDevice(ctx context.Context, input CreateDeviceInput) (*models.V1RegisterMachineResponse, error) {
	request := headscale_service.NewHeadscaleServiceRegisterMachineParams()
	request.SetContext(ctx)
	request.SetKey(input.Key)
	request.SetUser(input.User)

	resp, err := h.client.HeadscaleService.HeadscaleServiceRegisterMachine(request)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *HeadscaleService) ExpireDevice(ctx context.Context, deviceId string) (*models.V1ExpireMachineResponse, error) {
	request := headscale_service.NewHeadscaleServiceExpireMachineParams()
	request.SetContext(ctx)
	request.SetMachineID(deviceId)

	resp, err := h.client.HeadscaleService.HeadscaleServiceExpireMachine(request)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *HeadscaleService) DeleteDevice(ctx context.Context, deviceId string) error {
	request := headscale_service.NewHeadscaleServiceDeleteMachineParams()
	request.SetContext(ctx)
	request.SetMachineID(deviceId)

	_, err := h.client.HeadscaleService.HeadscaleServiceDeleteMachine(request)
	if err != nil {
		return err
	}
	return nil
}

func (h *HeadscaleService) RenameDevice(ctx context.Context, deviceId string, name string) (*models.V1RenameMachineResponse, error) {
	request := headscale_service.NewHeadscaleServiceRenameMachineParams()
	request.SetContext(ctx)
	request.SetMachineID(deviceId)
	request.SetNewName(name)

	resp, err := h.client.HeadscaleService.HeadscaleServiceRenameMachine(request)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *HeadscaleService) GetDeviceRoutes(ctx context.Context, deviceId string) (*models.V1GetMachineRoutesResponse, error) {
	request := headscale_service.NewHeadscaleServiceGetMachineRoutesParams()
	request.SetContext(ctx)
	request.SetMachineID(deviceId)

	resp, err := h.client.HeadscaleService.HeadscaleServiceGetMachineRoutes(request)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *HeadscaleService) TagDevice(ctx context.Context, deviceId string, tags []string) (*models.V1SetTagsResponse, error) {
	request := headscale_service.NewHeadscaleServiceSetTagsParams()
	request.SetContext(ctx)
	request.SetMachineID(deviceId)
	request.SetBody(headscale_service.HeadscaleServiceSetTagsBody{
		Tags: tags,
	})

	resp, err := h.client.HeadscaleService.HeadscaleServiceSetTags(request)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *HeadscaleService) MoveDevice(ctx context.Context, deviceId string, user string) (*models.V1MoveMachineResponse, error) {
	request := headscale_service.NewHeadscaleServiceMoveMachineParams()
	request.SetContext(ctx)
	request.SetMachineID(deviceId)
	request.SetUser(&user)

	resp, err := h.client.HeadscaleService.HeadscaleServiceMoveMachine(request)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *HeadscaleService) ListPreAuthKeys(ctx context.Context, user *string) (*models.V1ListPreAuthKeysResponse, error) {
	request := headscale_service.NewHeadscaleServiceListPreAuthKeysParams()
	request.SetContext(ctx)
	request.SetUser(user)

	resp, err := h.client.HeadscaleService.HeadscaleServiceListPreAuthKeys(request)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

type CreatePreAuthKeyInput struct {
	User       string
	Reusable   bool
	Ephermeral bool
	Expiration *strfmt.DateTime
	aclTags    []string
}

func (h *HeadscaleService) CreatePreAuthKey(ctx context.Context, input CreatePreAuthKeyInput) (*models.V1CreatePreAuthKeyResponse, error) {
	request := headscale_service.NewHeadscaleServiceCreatePreAuthKeyParams()
	request.SetContext(ctx)
	body := &models.V1CreatePreAuthKeyRequest{
		User:      input.User,
		Reusable:  input.Reusable,
		Ephemeral: input.Ephermeral,
		ACLTags:   input.aclTags,
	}
	if input.Expiration != nil {
		body.Expiration = *input.Expiration
	}
	request.SetBody(body)

	resp, err := h.client.HeadscaleService.HeadscaleServiceCreatePreAuthKey(request)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
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
		return err
	}
	return nil
}

func (h *HeadscaleService) ListRoutes(ctx context.Context) (*models.V1GetRoutesResponse, error) {
	request := headscale_service.NewHeadscaleServiceGetRoutesParams()
	request.SetContext(ctx)

	resp, err := h.client.HeadscaleService.HeadscaleServiceGetRoutes(request)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *HeadscaleService) DeleteRoute(ctx context.Context, route string) error {
	request := headscale_service.NewHeadscaleServiceDeleteRouteParams()
	request.SetContext(ctx)
	request.SetRouteID(route)

	_, err := h.client.HeadscaleService.HeadscaleServiceDeleteRoute(request)
	if err != nil {
		return err
	}
	return nil
}

func (h *HeadscaleService) DisableRoute(ctx context.Context, route string) error {
	request := headscale_service.NewHeadscaleServiceDisableRouteParams()
	request.SetContext(ctx)
	request.SetRouteID(route)

	_, err := h.client.HeadscaleService.HeadscaleServiceDisableRoute(request)
	if err != nil {
		return err
	}

	return nil
}

func (h *HeadscaleService) EnableRoute(ctx context.Context, route string) error {
	request := headscale_service.NewHeadscaleServiceEnableRouteParams()
	request.SetContext(ctx)
	request.SetRouteID(route)

	_, err := h.client.HeadscaleService.HeadscaleServiceEnableRoute(request)
	if err != nil {
		return err
	}

	return nil
}

func (h *HeadscaleService) ListUsers(ctx context.Context) (*models.V1ListUsersResponse, error) {
	request := headscale_service.NewHeadscaleServiceListUsersParams()
	request.SetContext(ctx)

	resp, err := h.client.HeadscaleService.HeadscaleServiceListUsers(request)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *HeadscaleService) GetUser(ctx context.Context, name string) (*models.V1GetUserResponse, error) {
	request := headscale_service.NewHeadscaleServiceGetUserParams()
	request.SetContext(ctx)
	request.SetName(name)

	resp, err := h.client.HeadscaleService.HeadscaleServiceGetUser(request)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *HeadscaleService) CreateUser(ctx context.Context, name string) (*models.V1CreateUserResponse, error) {
	request := headscale_service.NewHeadscaleServiceCreateUserParams()
	request.SetContext(ctx)
	request.SetBody(&models.V1CreateUserRequest{
		Name: name,
	})

	resp, err := h.client.HeadscaleService.HeadscaleServiceCreateUser(request)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *HeadscaleService) DeleteUser(ctx context.Context, name string) error {
	request := headscale_service.NewHeadscaleServiceDeleteUserParams()
	request.SetContext(ctx)
	request.SetName(name)

	_, err := h.client.HeadscaleService.HeadscaleServiceDeleteUser(request)
	if err != nil {
		return err
	}
	return nil
}

func (h *HeadscaleService) RenameUser(ctx context.Context, oldName string, newName string) error {
	request := headscale_service.NewHeadscaleServiceRenameUserParams()
	request.SetContext(ctx)
	request.SetNewName(newName)
	request.SetOldName(oldName)

	_, err := h.client.HeadscaleService.HeadscaleServiceRenameUser(request)
	if err != nil {
		return err
	}
	return nil
}
