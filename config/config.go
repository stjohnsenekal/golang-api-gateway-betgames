package config

type Configuration struct {
	Port              			int
	Connection_String 			string
	Secret						string
	Log_file_path				string
	Error_file_path				string

	HTTPS						bool
	CERT_file_path				string
	KEY_file_path				string

	RabbitMq_Username				string
	RabbitMq_Password				string
	RabbitMq_Host					string
	RabbitMq_Port					string
	RabbitMq_Vhost					string
	RabbitMq_Exchange_Name			string
	RabbitMq_Exchange_Type			string

	Redis_UserId_Expire		    int
	Redis_UserToken_Expire		int

	RejectInvalidSignatures 	bool
	RejectExpiredRequests		bool

	HealthCheckPoint			int

	RequestTimeoutLength	int



}

