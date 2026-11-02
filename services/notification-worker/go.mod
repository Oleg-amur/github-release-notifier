module github.com/Oleg-amur/github-release-notifier/services/notification-worker

go 1.25.3

require (
	github.com/Oleg-amur/github-release-notifier/shared/contracts v0.0.0
	github.com/Oleg-amur/github-release-notifier/shared/messaging v0.0.0
	github.com/ilyakaznacheev/cleanenv v1.5.0
	github.com/rabbitmq/amqp091-go v1.10.0
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace github.com/Oleg-amur/github-release-notifier/shared/contracts => ../../shared/contracts

replace github.com/Oleg-amur/github-release-notifier/shared/messaging => ../../shared/messaging
