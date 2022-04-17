package broker

type Config struct {
}

func GetConfigs() map[string]string {
	return map[string]string{
		"Access-Control-Allow-Origin": "*",
	}
}
