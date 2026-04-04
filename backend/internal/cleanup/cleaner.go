package cleanup

import (
	"context"
	"log/slog"
	"time"

	"github.com/0x2E/fusion/internal/config"
	"github.com/0x2E/fusion/internal/store"
)

const checkInterval = 24 * time.Hour

type Cleaner struct {
	store     *store.Store
	retention time.Duration
	logger    *slog.Logger
}

func New(st *store.Store, cfg *config.Config) *Cleaner {
	if cfg.ItemRetentionDays == 0 {
		return nil
	}
	return &Cleaner{
		store:     st,
		retention: time.Duration(cfg.ItemRetentionDays) * 24 * time.Hour,
		logger:    slog.Default(),
	}
}

func (c *Cleaner) Start(ctx context.Context) error {
	c.logger.Info("cleanup service started", "retention_days", int(c.retention.Hours()/24))

	c.run()

	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			c.logger.Info("cleanup service stopping")
			return ctx.Err()
		case <-ticker.C:
			c.run()
		}
	}
}

func (c *Cleaner) run() {
	cutoff := time.Now().Add(-c.retention).Unix()
	deleted, err := c.store.DeleteOldReadItems(cutoff)
	if err != nil {
		c.logger.Error("failed to delete old read items", "error", err)
		return
	}
	if deleted > 0 {
		c.logger.Info("deleted old read items", "count", deleted)
	}
}
