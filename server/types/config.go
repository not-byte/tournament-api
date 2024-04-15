package types

type Config struct {
	DbUser      string `json:"dbUser"`
	DbPassword  string `json:"dbPass"`
	DbHost      string `json:"dbHost"`
	Environment string `json:"environment"`
}
