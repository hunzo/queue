package models

type Config struct {
	AmqpUri      string `mapstructure:"amqp_uri"`
	GraylogSrv   string `mapstructure:"graylog_srv"`
	QueueName    string `mapstructure:"queue_name"`
	AmqpUsername string `mapstructure:"amqp_username"`
	AmqpPassword string `mapstructure:"amqp_password"`
}

type Payload struct {
	Count     int16  `json:"count"`
	Message   string `json:"message"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
}

type GraylogPayload struct {
	Version      string `json:"version"`
	Host         string `json:"host"`
	ShortMessage string `json:"short_message"`
	StartTime    string `json:"start_time"`
	Count        int    `json:"count"`
	Check        int    `json:"check,omitempty"`
}
