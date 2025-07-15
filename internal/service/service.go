package service

import (
	"context"
	"net/url"

	"github.com/awlsring/terraform-provider-headscale/internal/gen/client"
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

	GetUser(ctx context.Context, input GetUserInput) (*models.V1User, error)
	ListUsers(ctx context.Context) ([]*models.V1User, error)
	CreateUser(ctx context.Context, input CreateUserInput) (*models.V1User, error)
	DeleteUser(ctx context.Context, userId string) error
	RenameUser(ctx context.Context, name string, userId string) (*models.V1User, error)
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
