package property_v1

import (
	"context"
	"github.com/antinvestor/apis"
	"testing"
)

func TestNewPropertyClient(t *testing.T) {

	ctx := context.Background()

	propertyCli, err := NewPropertyClient(ctx, apis.WithEndpoint("127.0.0.1:7020"))
	if err != nil {
		t.Errorf("Could not setup notification client : %v", err)
	}

	if propertyCli == nil {
		t.Error("Could not work with nill notification client ")
	}

}

func TestNotificationClient_Send(t *testing.T) {

	ctx := context.Background()

	notificationCli, err := NewPropertyClient(ctx, apis.WithEndpoint("127.0.0.1:7020"))
	if err != nil {
		t.Errorf("Could not setup notification client : %v", err)
	}

	payload := map[string]string{}
	payload["testing"] = "test"

	_, err = notificationCli.Send(ctx, "testing", "contact", "en", "client.test", payload)
	if err != nil {
		t.Errorf("Could not send message from notification client : %v", err)
	}

}
