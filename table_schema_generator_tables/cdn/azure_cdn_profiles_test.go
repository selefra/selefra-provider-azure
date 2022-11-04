package cdn

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createRoutesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCDNRoutesClient(ctrl)
	s := services.Services{
		CDN: services.CDNClient{
			Routes: mockClient,
		},
	}

	data := cdn.Route{}
	require.Nil(t, faker.FakeObject(&data))

	result := cdn.NewRouteListResultPage(cdn.RouteListResult{Value: &[]cdn.Route{data}}, func(ctx context.Context, result cdn.RouteListResult) (cdn.RouteListResult, error) {
		return cdn.RouteListResult{}, nil
	})

	mockClient.EXPECT().ListByEndpoint(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}
func createRuleSetsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCDNRuleSetsClient(ctrl)
	s := services.Services{
		CDN: services.CDNClient{
			RuleSets: mockClient,
		},
	}

	data := cdn.RuleSet{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := cdn.NewRuleSetListResultPage(cdn.RuleSetListResult{Value: &[]cdn.RuleSet{data}}, func(ctx context.Context, result cdn.RuleSetListResult) (cdn.RuleSetListResult, error) {
		return cdn.RuleSetListResult{}, nil
	})

	mockClient.EXPECT().ListByProfile(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createRulesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCDNRulesClient(ctrl)
	s := services.Services{
		CDN: services.CDNClient{
			Rules: mockClient,
		},
	}

	data := cdn.Rule{}
	require.Nil(t, faker.FakeObject(&data))

	result := cdn.NewRuleListResultPage(cdn.RuleListResult{Value: &[]cdn.Rule{data}}, func(ctx context.Context, result cdn.RuleListResult) (cdn.RuleListResult, error) {
		return cdn.RuleListResult{}, nil
	})

	data.Actions = &[]cdn.BasicDeliveryRuleAction{}
	data.Conditions = &[]cdn.BasicDeliveryRuleCondition{}
	mockClient.EXPECT().ListByRuleSet(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}
func createSecurityPoliciesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCDNSecurityPoliciesClient(ctrl)
	s := services.Services{
		CDN: services.CDNClient{
			SecurityPolicies: mockClient,
		},
	}

	data := cdn.SecurityPolicy{}
	require.Nil(t, faker.FakeObject(&data))

	result := cdn.NewSecurityPolicyListResultPage(cdn.SecurityPolicyListResult{Value: &[]cdn.SecurityPolicy{data}}, func(ctx context.Context, result cdn.SecurityPolicyListResult) (cdn.SecurityPolicyListResult, error) {
		return cdn.SecurityPolicyListResult{}, nil
	})

	mockClient.EXPECT().ListByProfile(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createEndpointsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCDNEndpointsClient(ctrl)
	s := services.Services{
		CDN: services.CDNClient{
			Endpoints: mockClient,
		},
	}

	data := cdn.Endpoint{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := cdn.NewEndpointListResultPage(cdn.EndpointListResult{Value: &[]cdn.Endpoint{data}}, func(ctx context.Context, result cdn.EndpointListResult) (cdn.EndpointListResult, error) {
		return cdn.EndpointListResult{}, nil
	})

	mockClient.EXPECT().ListByProfile(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createCustomDomainsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCDNCustomDomainsClient(ctrl)
	s := services.Services{
		CDN: services.CDNClient{
			CustomDomains: mockClient,
		},
	}

	data := cdn.CustomDomain{}
	require.Nil(t, faker.FakeObject(&data))

	result := cdn.NewCustomDomainListResultPage(cdn.CustomDomainListResult{Value: &[]cdn.CustomDomain{data}}, func(ctx context.Context, result cdn.CustomDomainListResult) (cdn.CustomDomainListResult, error) {
		return cdn.CustomDomainListResult{}, nil
	})

	mockClient.EXPECT().ListByEndpoint(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}

func createProfilesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCDNProfilesClient(ctrl)
	s := services.Services{
		CDN: services.CDNClient{
			Profiles:         mockClient,
			Endpoints:        createEndpointsMock(t, ctrl).CDN.Endpoints,
			CustomDomains:    createCustomDomainsMock(t, ctrl).CDN.CustomDomains,
			Routes:           createRoutesMock(t, ctrl).CDN.Routes,
			RuleSets:         createRuleSetsMock(t, ctrl).CDN.RuleSets,
			Rules:            createRulesMock(t, ctrl).CDN.Rules,
			SecurityPolicies: createSecurityPoliciesMock(t, ctrl).CDN.SecurityPolicies,
		},
	}

	data := cdn.Profile{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := cdn.NewProfileListResultPage(cdn.ProfileListResult{Value: &[]cdn.Profile{data}}, func(ctx context.Context, result cdn.ProfileListResult) (cdn.ProfileListResult, error) {
		return cdn.ProfileListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestCDNProfiles(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureCdnProfilesGenerator{}), createProfilesMock, azure_client.TestOptions{})
}
