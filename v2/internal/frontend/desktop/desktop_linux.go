//go:build linux
// +build linux

package desktop

import (
	"context"
	"github.com/liugangnhm/wails/v2/internal/binding"
	"github.com/liugangnhm/wails/v2/internal/frontend"
	"github.com/liugangnhm/wails/v2/internal/frontend/desktop/linux"
	"github.com/liugangnhm/wails/v2/internal/logger"
	"github.com/liugangnhm/wails/v2/pkg/options"
)

func NewFrontend(ctx context.Context, appoptions *options.App, logger *logger.Logger, appBindings *binding.Bindings, dispatcher frontend.Dispatcher) frontend.Frontend {
	return linux.NewFrontend(ctx, appoptions, logger, appBindings, dispatcher)
}
