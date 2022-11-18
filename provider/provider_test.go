package provider

import (
	"context"
	"os"
	"testing"

	"github.com/selefra/selefra-provider-sdk/env"
	"github.com/selefra/selefra-provider-sdk/grpc/shard"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/storage/database_storage/postgresql_storage"
	"github.com/selefra/selefra-utils/pkg/json_util"
	"github.com/selefra/selefra-utils/pkg/pointer"
)

func TestProvider_PullTable(t *testing.T) {
	wk := "."
	config := `provider:
  xxx`
	os.Setenv("SELEFRA_DATABASE_DSN", "host=127.0.0.1 user=postgres password=pass port=5432 dbname=postgres sslmode=disable")
	os.Setenv("AZURE_TENANT_ID", "7d31740b-d054-490d-acfb-5622ed779d07")
	os.Setenv("AZURE_CLIENT_ID", "f4e8cd97-cb4e-4965-ba7a-3ea5f6960e67")
	os.Setenv("AZURE_CLIENT_SECRET", "8ps8Q~sZXj_lTALXHCAwDmQW7se4xMrwiDU2ccaO")

	myProvider := GetProvider()
	Pull(myProvider, config, wk, "*")
}

func Pull(myProvider *provider.Provider, config, workspace string, pullTables ...string) {

	diagnostics := schema.NewDiagnostics()

	initProviderRequest := &shard.ProviderInitRequest{
		Storage: &shard.Storage{
			Type:           0,
			StorageOptions: json_util.ToJsonBytes(postgresql_storage.NewPostgresqlStorageOptions(env.GetDatabaseDsn())),
		},
		Workspace:      &workspace,
		IsInstallInit:  pointer.TruePointer(),
		ProviderConfig: &config,
	}

	response, err := myProvider.Init(context.Background(), initProviderRequest)
	if err != nil {
		panic(diagnostics.AddFatal("init provider error: %s", err.Error()).ToString())
	}
	if diagnostics.AddDiagnostics(response.Diagnostics).HasError() {
		panic(diagnostics.ToString())
	}

	err = myProvider.PullTables(context.Background(), &shard.PullTablesRequest{
		Tables:        pullTables,
		MaxGoroutines: 100,
		Timeout:       1000 * 60 * 60,
	}, shard.NewFakeProviderServerSender())
	if err != nil {
		panic(diagnostics.AddFatal("provider pull table error: %s", err.Error()).ToString())
	}
}
