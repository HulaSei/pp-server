package subscription

import (
	"context"

	"github.com/perfect-panel/server/internal/module/subscription/internal/inventory"
)

// InventoryStore restricts order-driven inventory changes to subscription
// transactions and the durable inbox used to deduplicate them.
type InventoryStore = inventory.Store

// These persisted consumer identities must remain stable across deployments.
const (
	InventoryReserveConsumer = inventory.InventoryReserveConsumer
	InventoryRestoreConsumer = inventory.InventoryRestoreConsumer
)

var ErrOutOfStock = inventory.ErrOutOfStock

func ReserveInventoryOnce(ctx context.Context, store InventoryStore, orderNo string, subscribeID int64) error {
	return inventory.ReserveInventoryOnce(ctx, store, orderNo, subscribeID)
}

func RestoreInventoryOnce(ctx context.Context, store InventoryStore, orderNo string, subscribeID int64) error {
	return inventory.RestoreInventoryOnce(ctx, store, orderNo, subscribeID)
}
