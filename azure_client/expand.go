package azure_client

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func ExpandSingleSubscription() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		c := client.(*Client)
		if len(c.subscriptions) == 0 {
			return []*schema.ClientTaskContext{}
		}
		return []*schema.ClientTaskContext{
			&schema.ClientTaskContext{Client: c.withSubscription(c.subscriptions[0])},
		}
	}
}

func ExpandSubscription() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		c := client.(*Client)
		var cs = make([]*schema.ClientTaskContext, len(c.subscriptions))
		for i, subID := range c.subscriptions {
			cs[i] = &schema.ClientTaskContext{Client: c.withSubscription(subID)}
		}
		return cs
	}
}
