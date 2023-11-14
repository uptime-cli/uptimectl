package authmanager

import (
	"context"
	"errors"

	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
)

var (
	ErrNoLogin          = errors.New("no valid login")
	ErrSessionNotActive = errors.New("session not active")
)

type authManager struct {
}

func NewAuthManager() *authManager {
	return &authManager{}
}

func (a *authManager) Login(ctx context.Context, token string) error {
	return a.login(ctx, token)
}

func (a *authManager) Logout() error {
	context := contextmanager.CurrentContext()
	context.User.User.BetterUptimeToken = ""
	err := contextmanager.AddOrMergeContext(context)
	if err != nil {
		return err
	}
	return contextmanager.Save()
}

func (a *authManager) login(ctx context.Context, token string) error {
	context := contextmanager.CurrentContext()
	context.User.User.BetterUptimeToken = token
	err := contextmanager.AddOrMergeContext(context)
	if err != nil {
		return err
	}
	return contextmanager.Save()
}
