package config

type Config struct {
	ListenAddr string `mapstructure:"LISTEN_ADDR"`
	MongoURI   string `mapstructure:"MONGODB_URI"`
	MSSQLConn  string `mapstructure:"MSSQL_CONN"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
	JWTExpiry  int    `mapstructure:"JWT_EXPIRY_MIN"` // minutes
	RateLimit  int    `mapstructure:"RATE_LIMIT"`     // Requests-per-minute throttle
}
