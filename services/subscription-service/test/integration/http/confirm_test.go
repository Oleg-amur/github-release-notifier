//go:build integration

package http_test

import (
	"net/http"

	subscriptionmodels "github.com/Oleg-amur/github-release-notifier/services/subscription-service/internal/modules/subscriptions/domain"
	dto "github.com/Oleg-amur/github-release-notifier/services/subscription-service/internal/modules/subscriptions/transport/http"
	"github.com/Oleg-amur/github-release-notifier/services/subscription-service/test/integration/testkit"
)

func (s *HTTPSuite) TestConfirm_ActivatesSubscription() {
	subscribeRequest := dto.SubscribeRequest{
		Email: "user@example.com",
		Repo:  "owner/repo",
	}
	s.postSubscribe(subscribeRequest, http.StatusOK)

	token := s.receiveSubscriptionToken()

	s.getConfirm(token, http.StatusOK)

	testkit.AssertSubscriptionStatus(s.T(), s.app.DB, token, subscriptionmodels.StatusActive)
}

func (s *HTTPSuite) TestConfirm_ReturnsNotFoundForMissingToken() {
	s.getConfirm("missing-token", http.StatusNotFound)
}
