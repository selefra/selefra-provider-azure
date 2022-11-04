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

func createManagedDatabaseVulnerabilityAssessmentScansMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLManagedDatabaseVulnerabilityAssessmentScansClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ManagedDatabaseVulnerabilityAssessmentScans: mockClient,
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
func createManagedInstanceVulnerabilityAssessmentsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLManagedInstanceVulnerabilityAssessmentsClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ManagedInstanceVulnerabilityAssessments: mockClient,
		},
	}

	data := sql.ManagedInstanceVulnerabilityAssessment{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewManagedInstanceVulnerabilityAssessmentListResultPage(sql.ManagedInstanceVulnerabilityAssessmentListResult{Value: &[]sql.ManagedInstanceVulnerabilityAssessment{data}}, func(ctx context.Context, result sql.ManagedInstanceVulnerabilityAssessmentListResult) (sql.ManagedInstanceVulnerabilityAssessmentListResult, error) {
		return sql.ManagedInstanceVulnerabilityAssessmentListResult{}, nil
	})

	mockClient.EXPECT().ListByInstance(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createManagedInstanceEncryptionProtectorsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLManagedInstanceEncryptionProtectorsClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ManagedInstanceEncryptionProtectors: mockClient,
		},
	}

	data := sql.ManagedInstanceEncryptionProtector{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewManagedInstanceEncryptionProtectorListResultPage(sql.ManagedInstanceEncryptionProtectorListResult{Value: &[]sql.ManagedInstanceEncryptionProtector{data}}, func(ctx context.Context, result sql.ManagedInstanceEncryptionProtectorListResult) (sql.ManagedInstanceEncryptionProtectorListResult, error) {
		return sql.ManagedInstanceEncryptionProtectorListResult{}, nil
	})

	mockClient.EXPECT().ListByInstance(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createManagedDatabasesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLManagedDatabasesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ManagedDatabases: mockClient,
		},
	}

	data := sql.ManagedDatabase{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := sql.NewManagedDatabaseListResultPage(sql.ManagedDatabaseListResult{Value: &[]sql.ManagedDatabase{data}}, func(ctx context.Context, result sql.ManagedDatabaseListResult) (sql.ManagedDatabaseListResult, error) {
		return sql.ManagedDatabaseListResult{}, nil
	})

	mockClient.EXPECT().ListByInstance(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createManagedDatabaseVulnerabilityAssessmentsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLManagedDatabaseVulnerabilityAssessmentsClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ManagedDatabaseVulnerabilityAssessments: mockClient,
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

func createManagedInstancesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLManagedInstancesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ManagedInstances:                            mockClient,
			ManagedDatabases:                            createManagedDatabasesMock(t, ctrl).SQL.ManagedDatabases,
			ManagedDatabaseVulnerabilityAssessments:     createManagedDatabaseVulnerabilityAssessmentsMock(t, ctrl).SQL.ManagedDatabaseVulnerabilityAssessments,
			ManagedDatabaseVulnerabilityAssessmentScans: createManagedDatabaseVulnerabilityAssessmentScansMock(t, ctrl).SQL.ManagedDatabaseVulnerabilityAssessmentScans,
			ManagedInstanceVulnerabilityAssessments:     createManagedInstanceVulnerabilityAssessmentsMock(t, ctrl).SQL.ManagedInstanceVulnerabilityAssessments,
			ManagedInstanceEncryptionProtectors:         createManagedInstanceEncryptionProtectorsMock(t, ctrl).SQL.ManagedInstanceEncryptionProtectors,
		},
	}

	data := sql.ManagedInstance{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := sql.NewManagedInstanceListResultPage(sql.ManagedInstanceListResult{Value: &[]sql.ManagedInstance{data}}, func(ctx context.Context, result sql.ManagedInstanceListResult) (sql.ManagedInstanceListResult, error) {
		return sql.ManagedInstanceListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestSQLManagedInstances(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureSqlManagedInstancesGenerator{}), createManagedInstancesMock, azure_client.TestOptions{})
}
