package sql

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createServerVulnerabilityAssessmentsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLServerVulnerabilityAssessmentsClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ServerVulnerabilityAssessments: mockClient,
		},
	}

	data := sql.ServerVulnerabilityAssessment{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewServerVulnerabilityAssessmentListResultPage(sql.ServerVulnerabilityAssessmentListResult{Value: &[]sql.ServerVulnerabilityAssessment{data}}, func(ctx context.Context, result sql.ServerVulnerabilityAssessmentListResult) (sql.ServerVulnerabilityAssessmentListResult, error) {
		return sql.ServerVulnerabilityAssessmentListResult{}, nil
	})

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createBackupLongTermRetentionPoliciesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLBackupLongTermRetentionPoliciesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			BackupLongTermRetentionPolicies: mockClient,
		},
	}

	data := sql.BackupLongTermRetentionPolicy{}
	require.Nil(t, faker.FakeObject(&data))

	result := data

	mockClient.EXPECT().ListByDatabase(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}
func createTransparentDataEncryptionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLTransparentDataEncryptionsClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			TransparentDataEncryptions: mockClient,
		},
	}

	data := sql.TransparentDataEncryption{}
	require.Nil(t, faker.FakeObject(&data))

	result := data

	mockClient.EXPECT().Get(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}
func createDatabaseVulnerabilityAssessmentsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLDatabaseVulnerabilityAssessmentsClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			DatabaseVulnerabilityAssessments: mockClient,
		},
	}

	data := sql.DatabaseVulnerabilityAssessment{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewDatabaseVulnerabilityAssessmentListResultPage(sql.DatabaseVulnerabilityAssessmentListResult{Value: &[]sql.DatabaseVulnerabilityAssessment{data}}, func(ctx context.Context, result sql.DatabaseVulnerabilityAssessmentListResult) (sql.DatabaseVulnerabilityAssessmentListResult, error) {
		return sql.DatabaseVulnerabilityAssessmentListResult{}, nil
	})

	mockClient.EXPECT().ListByDatabase(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}
func createDatabaseThreatDetectionPoliciesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLDatabaseThreatDetectionPoliciesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			DatabaseThreatDetectionPolicies: mockClient,
		},
	}

	data := sql.DatabaseSecurityAlertPolicy{}
	require.Nil(t, faker.FakeObject(&data))

	result := data

	mockClient.EXPECT().Get(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}
func createServerDevOpsAuditingSettingsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLServerDevOpsAuditingSettingsClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ServerDevOpsAuditingSettings: mockClient,
		},
	}

	data := sql.ServerDevOpsAuditingSettings{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewServerDevOpsAuditSettingsListResultPage(sql.ServerDevOpsAuditSettingsListResult{Value: &[]sql.ServerDevOpsAuditingSettings{data}}, func(ctx context.Context, result sql.ServerDevOpsAuditSettingsListResult) (sql.ServerDevOpsAuditSettingsListResult, error) {
		return sql.ServerDevOpsAuditSettingsListResult{}, nil
	})

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createServerSecurityAlertPoliciesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLServerSecurityAlertPoliciesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ServerSecurityAlertPolicies: mockClient,
		},
	}

	data := sql.ServerSecurityAlertPolicy{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewLogicalServerSecurityAlertPolicyListResultPage(sql.LogicalServerSecurityAlertPolicyListResult{Value: &[]sql.ServerSecurityAlertPolicy{data}}, func(ctx context.Context, result sql.LogicalServerSecurityAlertPolicyListResult) (sql.LogicalServerSecurityAlertPolicyListResult, error) {
		return sql.LogicalServerSecurityAlertPolicyListResult{}, nil
	})

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createDatabaseBlobAuditingPoliciesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLDatabaseBlobAuditingPoliciesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			DatabaseBlobAuditingPolicies: mockClient,
		},
	}

	data := sql.DatabaseBlobAuditingPolicy{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewDatabaseBlobAuditingPolicyListResultPage(sql.DatabaseBlobAuditingPolicyListResult{Value: &[]sql.DatabaseBlobAuditingPolicy{data}}, func(ctx context.Context, result sql.DatabaseBlobAuditingPolicyListResult) (sql.DatabaseBlobAuditingPolicyListResult, error) {
		return sql.DatabaseBlobAuditingPolicyListResult{}, nil
	})

	mockClient.EXPECT().ListByDatabase(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}
func createDatabasesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLDatabasesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			Databases: mockClient,
		},
	}

	data := sql.Database{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := sql.NewDatabaseListResultPage(sql.DatabaseListResult{Value: &[]sql.Database{data}}, func(ctx context.Context, result sql.DatabaseListResult) (sql.DatabaseListResult, error) {
		return sql.DatabaseListResult{}, nil
	})

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createDatabaseVulnerabilityAssessmentScansMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLDatabaseVulnerabilityAssessmentScansClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			DatabaseVulnerabilityAssessmentScans: mockClient,
		},
	}

	data := sql.VulnerabilityAssessmentScanRecord{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewVulnerabilityAssessmentScanRecordListResultPage(sql.VulnerabilityAssessmentScanRecordListResult{Value: &[]sql.VulnerabilityAssessmentScanRecord{data}}, func(ctx context.Context, result sql.VulnerabilityAssessmentScanRecordListResult) (sql.VulnerabilityAssessmentScanRecordListResult, error) {
		return sql.VulnerabilityAssessmentScanRecordListResult{}, nil
	})

	mockClient.EXPECT().ListByDatabase(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}
func createEncryptionProtectorsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLEncryptionProtectorsClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			EncryptionProtectors: mockClient,
		},
	}

	data := sql.EncryptionProtector{}
	require.Nil(t, faker.FakeObject(&data))

	result := data

	mockClient.EXPECT().Get(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createVirtualNetworkRulesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLVirtualNetworkRulesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			VirtualNetworkRules: mockClient,
		},
	}

	data := sql.VirtualNetworkRule{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewVirtualNetworkRuleListResultPage(sql.VirtualNetworkRuleListResult{Value: &[]sql.VirtualNetworkRule{data}}, func(ctx context.Context, result sql.VirtualNetworkRuleListResult) (sql.VirtualNetworkRuleListResult, error) {
		return sql.VirtualNetworkRuleListResult{}, nil
	})

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createServerAdminsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLServerAdminsClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ServerAdmins: mockClient,
		},
	}

	data := sql.ServerAzureADAdministrator{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewAdministratorListResultPage(sql.AdministratorListResult{Value: &[]sql.ServerAzureADAdministrator{data}}, func(ctx context.Context, result sql.AdministratorListResult) (sql.AdministratorListResult, error) {
		return sql.AdministratorListResult{}, nil
	})

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createServerBlobAuditingPoliciesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLServerBlobAuditingPoliciesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ServerBlobAuditingPolicies: mockClient,
		},
	}

	data := sql.ServerBlobAuditingPolicy{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewServerBlobAuditingPolicyListResultPage(sql.ServerBlobAuditingPolicyListResult{Value: &[]sql.ServerBlobAuditingPolicy{data}}, func(ctx context.Context, result sql.ServerBlobAuditingPolicyListResult) (sql.ServerBlobAuditingPolicyListResult, error) {
		return sql.ServerBlobAuditingPolicyListResult{}, nil
	})

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createFirewallRulesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLFirewallRulesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			FirewallRules: mockClient,
		},
	}

	data := sql.FirewallRule{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.FirewallRuleListResult{Value: &[]sql.FirewallRule{data}}

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}

func createServersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLServersClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			Servers:                              mockClient,
			FirewallRules:                        createFirewallRulesMock(t, ctrl).SQL.FirewallRules,
			Databases:                            createDatabasesMock(t, ctrl).SQL.Databases,
			DatabaseBlobAuditingPolicies:         createDatabaseBlobAuditingPoliciesMock(t, ctrl).SQL.DatabaseBlobAuditingPolicies,
			DatabaseVulnerabilityAssessments:     createDatabaseVulnerabilityAssessmentsMock(t, ctrl).SQL.DatabaseVulnerabilityAssessments,
			DatabaseVulnerabilityAssessmentScans: createDatabaseVulnerabilityAssessmentScansMock(t, ctrl).SQL.DatabaseVulnerabilityAssessmentScans,
			BackupLongTermRetentionPolicies:      createBackupLongTermRetentionPoliciesMock(t, ctrl).SQL.BackupLongTermRetentionPolicies,
			DatabaseThreatDetectionPolicies:      createDatabaseThreatDetectionPoliciesMock(t, ctrl).SQL.DatabaseThreatDetectionPolicies,
			TransparentDataEncryptions:           createTransparentDataEncryptionsMock(t, ctrl).SQL.TransparentDataEncryptions,
			EncryptionProtectors:                 createEncryptionProtectorsMock(t, ctrl).SQL.EncryptionProtectors,
			VirtualNetworkRules:                  createVirtualNetworkRulesMock(t, ctrl).SQL.VirtualNetworkRules,
			ServerAdmins:                         createServerAdminsMock(t, ctrl).SQL.ServerAdmins,
			ServerBlobAuditingPolicies:           createServerBlobAuditingPoliciesMock(t, ctrl).SQL.ServerBlobAuditingPolicies,
			ServerDevOpsAuditingSettings:         createServerDevOpsAuditingSettingsMock(t, ctrl).SQL.ServerDevOpsAuditingSettings,
			ServerVulnerabilityAssessments:       createServerVulnerabilityAssessmentsMock(t, ctrl).SQL.ServerVulnerabilityAssessments,
			ServerSecurityAlertPolicies:          createServerSecurityAlertPoliciesMock(t, ctrl).SQL.ServerSecurityAlertPolicies,
		},
	}

	data := sql.Server{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := sql.NewServerListResultPage(sql.ServerListResult{Value: &[]sql.Server{data}}, func(ctx context.Context, result sql.ServerListResult) (sql.ServerListResult, error) {
		return sql.ServerListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestSQLServers(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureSqlServersGenerator{}), createServersMock, azure_client.TestOptions{})
}
