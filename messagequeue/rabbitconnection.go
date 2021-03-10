package messagequeue

import (
	"api-gateway/config"
	"strings"
)

var env config.Configuration

func SetConfiguration(configuration config.Configuration) {
	// injection of environment configuration
	env = configuration
}

func getConnectionString() string {
	var sb strings.Builder
	sb.WriteString("amqp://")
	sb.WriteString(env.RabbitMq_Username)
	sb.WriteString(":")
	sb.WriteString(env.RabbitMq_Password)
	sb.WriteString("@")
	sb.WriteString(env.RabbitMq_Host)
	sb.WriteString(":")
	sb.WriteString(env.RabbitMq_Port)
	sb.WriteString("/")
	if len(env.RabbitMq_Vhost) > 0 {
		sb.WriteString("%2F")
		sb.WriteString(env.RabbitMq_Vhost)
	}

	return sb.String()
}

