package main

import (
	"log/slog"

	subscriptionpostgresql "github.com/Oleg-amur/github-release-notifier/services/subscription-service/internal/modules/subscriptions/persistence/postgresql"
	subscriptionusecase "github.com/Oleg-amur/github-release-notifier/services/subscription-service/internal/modules/subscriptions/usecase"
)

func setupSubscriptionsModule(
	log *slog.Logger,
	subscriberStore *subscriptionpostgresql.SubscriberStore,
	subscriptionStore *subscriptionpostgresql.SubscriptionStore,
	subscriptionCreator subscriptionusecase.SubscriptionCreator,
	repositoryTracker subscriptionusecase.RepositoryTracker,
	repositoryMetadata subscriptionusecase.RepositoryMetadataReader,
) subscriptionusecase.SubscriptionUsecases {
	getOrCreateSubscriber := subscriptionusecase.NewGetOrCreateSubscriber(log, subscriberStore)

	return subscriptionusecase.SubscriptionUsecases{
		SubscribeToRepository: subscriptionusecase.NewSubscribeToRepository(
			log,
			getOrCreateSubscriber,
			repositoryTracker,
			subscriptionCreator,
		),
		ConfirmSubscription:       subscriptionusecase.NewConfirmSubscription(subscriptionStore),
		UnsubscribeFromRepository: subscriptionusecase.NewUnsubscribeFromRepository(subscriptionStore),
		ListSubscriptions:         subscriptionusecase.NewListSubscriptions(subscriptionStore, repositoryMetadata),
		ListActiveByRepository:    subscriptionusecase.NewListActiveSubscriptionsByRepository(subscriptionStore),
	}
}
