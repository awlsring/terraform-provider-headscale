package service

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/awlsring/terraform-provider-headscale/internal/gen/client"
	"github.com/awlsring/terraform-provider-headscale/internal/gen/client/headscale_service"
	"github.com/awlsring/terraform-provider-headscale/internal/gen/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Headscale interface {
	ListAPIKeys(ctx context.Context) ([]*models.V1APIKey, error)
	CreateAPIKey(ctx context.Context, expiration *strfmt.DateTime) (string, error)
	ExpireAPIKey(ctx context.Context, key string) error
	ListDevices(ctx context.Context, user *string) ([]*models.V1Node, error)
	GetDevice(ctx context.Context, deviceId string) (*models.V1Node, error)
	CreateDevice(ctx context.Context, user string, key string) (*models.V1Node, error)
	ExpireDevice(ctx context.Context, deviceId string) (*models.V1Node, error)
	DeleteDevice(ctx context.Context, deviceId string) error
	RenameDevice(ctx context.Context, deviceId string, newName string) (*models.V1Node, error)
	GetDeviceRoutes(ctx context.Context, deviceId string) ([]*models.V1Route, error)
	TagDevice(ctx context.Context, deviceId string, tags []string) (*models.V1Node, error)
	MoveDevice(ctx context.Context, deviceId string, newOwner string) (*models.V1Node, error)
	ListPreAuthKeys(ctx context.Context, user string) ([]*models.V1PreAuthKey, error)
	CreatePreAuthKey(ctx context.Context, input CreatePreAuthKeyInput) (*models.V1PreAuthKey, error)
	ExpirePreAuthKey(ctx context.Context, user string, key string) error
	ListRoutes(ctx context.Context) ([]*models.V1Route, error)
	DeleteRoute(ctx context.Context, routeId string) error
	DisableRoute(ctx context.Context, routeId string) error
	EnableRoute(ctx context.Context, routeId string) error
	GetUserById(ctx context.Context, userId string) (*models.V1User, error)
	GetUserByName(ctx context.Context, userId string) (*models.V1User, error)
	ListUsers(ctx context.Context) ([]*models.V1User, error)
	CreateUser(ctx context.Context, name string) (*models.V1User, error)
	DeleteUser(ctx context.Context, userId string) error
	RenameUser(ctx context.Context, oldName string, newName string) (*models.V1User, error)
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

func (h *HeadscaleService) ListDevices(ctx context.Context, user *string) ([]*models.V1Node, error) {
	request := headscale_service.NewHeadscaleServiceListNodesParams()
	request.SetContext(ctx)
	request.SetUser(user)

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

type CreateDeviceInput struct {
	User *string
	Key  *string
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

func (h *HeadscaleService) GetDeviceRoutes(ctx context.Context, deviceId string) ([]*models.V1Route, error) {
	request := headscale_service.NewHeadscaleServiceGetNodeRoutesParams()
	request.SetContext(ctx)
	request.SetNodeID(deviceId)

	resp, err := h.client.HeadscaleService.HeadscaleServiceGetNodeRoutes(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Routes, nil
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
	body := models.HeadscaleServiceMoveNodeBody{
		User: user,
	}
	request.SetBody(&body)

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

func (h *HeadscaleService) ListUsers(ctx context.Context) ([]*models.V1User, error) {
	request := headscale_service.NewHeadscaleServiceListUsersParams()
	request.SetContext(ctx)

	resp, err := h.client.HeadscaleService.HeadscaleServiceListUsers(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Users, nil
}
func (h *HeadscaleService) GetUserById(ctx context.Context, userId string) (*models.V1User, error) {
	request := headscale_service.NewHeadscaleServiceListUsersParams()
	request.SetContext(ctx)
	request.SetID(&userId)

	resp, err := h.client.HeadscaleService.HeadscaleServiceListUsers(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	// Check if any users were returned
	if len(resp.Payload.Users) == 0 {
		return nil, fmt.Errorf("user %q not found", userId)
	}

	// Return the first (and presumably only) user
	return resp.Payload.Users[0], nil
}

func (h *HeadscaleService) GetUserByName(ctx context.Context, name string) (*models.V1User, error) {
	request := headscale_service.NewHeadscaleServiceListUsersParams()
	request.SetContext(ctx)
	request.SetName(&name)

	resp, err := h.client.HeadscaleService.HeadscaleServiceListUsers(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	// Check if any users were returned
	if len(resp.Payload.Users) == 0 {
		return nil, fmt.Errorf("user %q not found", name)
	}

	// Return the first (and presumably only) user
	return resp.Payload.Users[0], nil
}

func (h *HeadscaleService) CreateUser(ctx context.Context, name string) (*models.V1User, error) {
	request := headscale_service.NewHeadscaleServiceCreateUserParams()
	request.SetContext(ctx)
	request.SetBody(&models.V1CreateUserRequest{
		Name: name,
	})

	resp, err := h.client.HeadscaleService.HeadscaleServiceCreateUser(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.User, nil
}

func (h *HeadscaleService) DeleteUser(ctx context.Context, userId string) error {
	request := headscale_service.NewHeadscaleServiceDeleteUserParams()
	request.SetContext(ctx)
	request.SetID(userId)

	_, err := h.client.HeadscaleService.HeadscaleServiceDeleteUser(request)
	if err != nil {
		return handleRequestError(err)
	}
	return nil
}

func (h *HeadscaleService) RenameUser(ctx context.Context, oldId string, newName string) (*models.V1User, error) {
	request := headscale_service.NewHeadscaleServiceRenameUserParams()
	request.SetContext(ctx)
	request.SetNewName(newName)
	request.SetOldID(oldId)

	resp, err := h.client.HeadscaleService.HeadscaleServiceRenameUser(request)
	if err != nil {
		return nil, handleRequestError(err)
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return resp.Payload.User, nil
}
