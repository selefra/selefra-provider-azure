package azure_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}
type Config struct {
	Subscriptions []string `yaml:"subscriptions,omitempty" mapstructure:"subscriptions"`
}
