package rubrikpolaris

type RadarEnabledClusters struct {
	Data struct {
		RadarClusterConnection struct {
			Nodes []struct {
				ID           string      `mapstructure:"id"`
				LambdaConfig interface{} `mapstructure:"lambdaConfig"`
				Name         string      `mapstructure:"name"`
			} `mapstructure:"nodes"`
		} `mapstructure:"radarClusterConnection"`
	} `mapstructure:"data"`
}
