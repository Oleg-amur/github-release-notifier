package usecase

import (
	"context"
	"testing"

	"github.com/Oleg-amur/github-release-notifier/services/subscription-service/internal/apperr"
)

func TestConfirmSubscription_Execute(t *testing.T) {
	tests := []struct {
		name        string
		token       string
		activateErr error
		wantErr     error
	}{
		{
			name:        "returns token not found when repository has no token",
			token:       "test-token",
			activateErr: apperr.ErrNotFound,
			wantErr:     apperr.ErrTokenNotFound,
		},
		{
			name:  "confirms subscription",
			token: "test-token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := NewConfirmSubscription(&mockSubscriptionRepo{activateErr: tt.activateErr})

			err := usecase.Execute(context.Background(), tt.token)

			assertErrorIs(t, err, tt.wantErr)
		})
	}
}
