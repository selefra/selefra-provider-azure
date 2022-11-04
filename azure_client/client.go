package azure_client

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	_ "github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"

	"github.com/selefra/selefra-provider-azure/azure_client/services"
)

type Client struct {
	subscriptions  []string
	SubscriptionId string
	services       map[string]*services.Services
}

func (c *Client) SetSubscriptionServices(subscriptionId string, servicesSet services.Services) {
	c.services[subscriptionId] = &servicesSet
}

func (c *Client) withSubscription(subscriptionId string) *Client {
	return &Client{
		subscriptions:  c.subscriptions,
		services:       c.services,
		SubscriptionId: subscriptionId,
	}
}

func (c *Client) AzureServices() *services.Services {
	return c.services[c.SubscriptionId]
}

func NewClients(configs Configs) ([]*Client, error) {
	var clients []*Client

	for _, c := range configs.Providers {
		cls, err := newClient(c)
		if err != nil {
			return nil, err
		}
		clients = append(clients, cls)
	}
	return clients, nil
}

func newClient(config Config) (*Client, error) {
	azureAuth, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		azureAuth, err = auth.NewAuthorizerFromCLI()
		if err != nil {
			return nil, err
		}
	}

	var azCred azcore.TokenCredential
	azCred, err = azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		azCred, err = azidentity.NewEnvironmentCredential(nil)
		if err != nil {
			return nil, err
		}
	}
	client := NewAzureClient(config.Subscriptions)
	if len(config.Subscriptions) == 0 {
		ctx := context.Background()
		svc := subscription.NewSubscriptionsClient()
		svc.Authorizer = azureAuth
		res, err := svc.List(ctx)
		if err != nil {
			return nil, err
		}
		subscriptions := make([]string, 0)
		for res.NotDone() {
			for _, sub := range res.Values() {
				switch sub.State {
				case subscription.Disabled:
					fmt.Println("Not fetching from subscription because it is disabled subscription", *sub.SubscriptionID)
				case subscription.Deleted:
					fmt.Println("Not fetching from subscription because it is deleted subscription", *sub.SubscriptionID)
				default:
					subscriptions = append(subscriptions, *sub.SubscriptionID)
				}
			}
			err := res.NextWithContext(ctx)
			if err != nil {
				return nil, err
			}
			client.subscriptions = subscriptions
		}
		if len(client.subscriptions) == 0 {
			return nil, errors.New("could not find any subscription")
		}
	}
	for _, sub := range client.subscriptions {
		svcs, err := services.InitServices(sub, azureAuth, azCred)
		if err != nil {
			return nil, err
		}
		client.SetSubscriptionServices(sub, svcs)
	}
	return client, nil
}

func NewAzureClient(subscriptions []string) *Client {
	return &Client{
		subscriptions: subscriptions,
		services:      make(map[string]*services.Services),
	}
}

func (c *Client) ID() string {
	return c.SubscriptionId
}

func (c Client) Services() *services.Services {
	return c.services[c.SubscriptionId]
}

func ScopeSubscription(subscriptionID string) string {
	return "subscriptions/" + subscriptionID
}

type ResourceDetails struct {
	Subscription  string
	ResourceGroup string
	Provider      string
	ResourceType  string
	ResourceName  string
}

const resourceIDPatternText = `(?i)subscriptions/(.+)/resourceGroups/(.+)/providers/(.+?)/(.+?)/(.+)`

var resourceIDPattern = regexp.MustCompile(resourceIDPatternText)

func ParseResourceID(resourceID string) (ResourceDetails, error) {
	match := resourceIDPattern.FindStringSubmatch(resourceID)

	if len(match) == 0 {
		return ResourceDetails{}, fmt.Errorf("parsing failed for %s. Invalid resource Id format", resourceID)
	}

	v := strings.Split(match[5], "/")
	resourceName := v[len(v)-1]

	result := ResourceDetails{
		Subscription:  match[1],
		ResourceGroup: match[2],
		Provider:      match[3],
		ResourceType:  match[4],
		ResourceName:  resourceName,
	}

	return result, nil
}
