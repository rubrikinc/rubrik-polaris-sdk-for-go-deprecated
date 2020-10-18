package rubrikpolaris

import "time"

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

type AllEvent struct {
	Data struct {
		ActivitySeriesConnection struct {
			Edges []struct {
				Node struct {
					ID                   int         `mapstructure:"id"`
					Fid                  string      `mapstructure:"fid"`
					ActivitySeriesID     string      `mapstructure:"activitySeriesId"`
					LastUpdated          time.Time   `mapstructure:"lastUpdated"`
					LastActivityType     string      `mapstructure:"lastActivityType"`
					LastActivityStatus   string      `mapstructure:"lastActivityStatus"`
					ObjectID             string      `mapstructure:"objectId"`
					ObjectName           string      `mapstructure:"objectName"`
					ObjectType           string      `mapstructure:"objectType"`
					Severity             string      `mapstructure:"severity"`
					Progress             interface{} `mapstructure:"progress"`
					IsCancelable         interface{} `mapstructure:"isCancelable"`
					IsPolarisEventSeries bool        `mapstructure:"isPolarisEventSeries"`
					Cluster              struct {
						ID   string `mapstructure:"id"`
						Name string `mapstructure:"name"`
					} `mapstructure:"cluster"`
					ActivityConnection struct {
						Nodes []struct {
							ID      string `mapstructure:"id"`
							Message string `mapstructure:"message"`
						} `mapstructure:"nodes"`
					} `mapstructure:"activityConnection"`
				} `mapstructure:"node"`
			} `mapstructure:"edges"`
			PageInfo struct {
				EndCursor       string `mapstructure:"endCursor"`
				HasNextPage     bool   `mapstructure:"hasNextPage"`
				HasPreviousPage bool   `mapstructure:"hasPreviousPage"`
			} `mapstructure:"pageInfo"`
		} `mapstructure:"activitySeriesConnection"`
	} `mapstructure:"data"`
}
