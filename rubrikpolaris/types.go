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

type AllEvent struct {
	Data struct {
		ActivitySeriesConnection struct {
			Edges []struct {
				Node struct {
					ID                   int         `mapstructure:"id"`
					Fid                  string      `mapstructure:"fid"`
					ActivitySeriesID     string      `mapstructure:"activitySeriesId"`
					LastUpdated          string      `mapstructure:"lastUpdated"`
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

type AllAuditLog struct {
	Data struct {
		UserAuditConnection struct {
			Edges []struct {
				Node struct {
					ID       string `mapstructure:"id"`
					Message  string `mapstructure:"message"`
					Time     string `mapstructure:"time"`
					Severity string `mapstructure:"severity"`
					Status   string `mapstructure:"status"`
					Cluster  struct {
						ID   string `mapstructure:"id"`
						Name string `mapstructure:"name"`
					} `mapstructure:"cluster"`
				} `mapstructure:"node"`
			} `mapstructure:"edges"`
		} `mapstructure:"userAuditConnection"`
	} `mapstructure:"data"`
}

type AllEvents struct {
	Data struct {
		ActivitySeriesConnection struct {
			Edges []struct {
				Node struct {
					ID                   int         `mapstructure:"id"`
					Fid                  string      `mapstructure:"fid"`
					ActivitySeriesID     string      `mapstructure:"activitySeriesId"`
					LastUpdated          string      `mapstructure:"lastUpdated"`
					LastActivityType     string      `mapstructure:"lastActivityType"`
					LastActivityStatus   string      `mapstructure:"lastActivityStatus"`
					ObjectID             string      `mapstructure:"objectId"`
					ObjectName           string      `mapstructure:"objectName"`
					ObjectType           string      `mapstructure:"objectType"`
					Severity             string      `mapstructure:"severity"`
					Progress             string      `mapstructure:"progress"`
					IsCancelable         interface{} `mapstructure:"isCancelable"`
					IsPolarisEventSeries bool        `mapstructure:"isPolarisEventSeries"`
					Typename             string      `mapstructure:"__typename"`
					Cluster              struct {
						ID       string `mapstructure:"id"`
						Name     string `mapstructure:"name"`
						Typename string `mapstructure:"__typename"`
					} `mapstructure:"cluster"`
					ActivityConnection struct {
						Nodes []struct {
							ID       string `mapstructure:"id"`
							Message  string `mapstructure:"message"`
							Typename string `mapstructure:"__typename"`
						} `mapstructure:"nodes"`
						Typename string `mapstructure:"__typename"`
					} `mapstructure:"activityConnection"`
				} `mapstructure:"node"`
				Typename string `mapstructure:"__typename"`
			} `mapstructure:"edges"`
			PageInfo struct {
				EndCursor       string `mapstructure:"endCursor"`
				HasNextPage     bool   `mapstructure:"hasNextPage"`
				HasPreviousPage bool   `mapstructure:"hasPreviousPage"`
				Typename        string `mapstructure:"__typename"`
			} `mapstructure:"pageInfo"`
			Typename string `mapstructure:"__typename"`
		} `mapstructure:"activitySeriesConnection"`
	} `mapstructure:"data"`
}

type EventSeriesDetail struct {
	Data struct {
		ActivitySeries struct {
			ActivityConnection struct {
				Nodes []struct {
					Message  string `mapstructure:"message"`
					Status   string `mapstructure:"status"`
					Time     string `mapstructure:"time"`
					Severity string `mapstructure:"severity"`
				} `mapstructure:"nodes"`
			} `mapstructure:"activityConnection"`
			ID               int    `mapstructure:"id"`
			Fid              string `mapstructure:"fid"`
			ActivitySeriesID string `mapstructure:"activitySeriesId"`
			ObjectID         string `mapstructure:"objectId"`
			ObjectName       string `mapstructure:"objectName"`
			ObjectType       string `mapstructure:"objectType"`
			Cluster          struct {
				ID   string `mapstructure:"id"`
				Name string `mapstructure:"name"`
			} `mapstructure:"cluster"`
			LastActivityStatus string `mapstructure:"lastActivityStatus"`
		} `mapstructure:"activitySeries"`
	} `mapstructure:"data"`
}

type EventSeriesDetailMessage struct {
	Message          string `mapstructure:"message"`
	Status           string `mapstructure:"status"`
	Time             string `mapstructure:"time"`
	Severity         string `mapstructure:"severity"`
	ID               int    `mapstructure:"id"`
	Fid              string `mapstructure:"fid"`
	ActivitySeriesID string `mapstructure:"activitySeriesId"`
	ObjectID         string `mapstructure:"objectId"`
	ObjectName       string `mapstructure:"objectName"`
	ObjectType       string `mapstructure:"objectType"`
	Cluster          struct {
		ID   string `mapstructure:"id"`
		Name string `mapstructure:"name"`
	} `mapstructure:"cluster"`
}

type RadarEvent struct {
	Data struct {
		ActivitySeriesConnection struct {
			Edges []struct {
				Node struct {
					ID                 int    `mapstructure:"id"`
					Fid                string `mapstructure:"fid"`
					ActivitySeriesID   string `mapstructure:"activitySeriesId"`
					LastUpdated        string `mapstructure:"lastUpdated"`
					LastActivityType   string `mapstructure:"lastActivityType"`
					LastActivityStatus string `mapstructure:"lastActivityStatus"`
					ObjectID           string `mapstructure:"objectId"`
					ObjectName         string `mapstructure:"objectName"`
					ObjectType         string `mapstructure:"objectType"`
					Severity           string `mapstructure:"severity"`
					Progress           string `mapstructure:"progress"`
					Cluster            struct {
						ID   string `mapstructure:"id"`
						Name string `mapstructure:"name"`
					} `mapstructure:"cluster"`
					ActivityConnection struct {
						Nodes []struct {
							ID      string `mapstructure:"id"`
							Message string `mapstructure:"message"`
							Time    string `mapstructure:"time"`
						} `mapstructure:"nodes"`
					} `mapstructure:"activityConnection"`
				} `mapstructure:"node"`
			} `mapstructure:"edges"`
		} `mapstructure:"activitySeriesConnection"`
	} `mapstructure:"data"`
}

type PolarisEvents struct {
	Data struct {
		ActivitySeriesConnection struct {
			Edges []struct {
				Node struct {
					ID                 int    `mapstructure:"id"`
					Fid                string `mapstructure:"fid"`
					ActivitySeriesID   string `mapstructure:"activitySeriesId"`
					LastUpdated        string `mapstructure:"lastUpdated"`
					LastActivityType   string `mapstructure:"lastActivityType"`
					LastActivityStatus string `mapstructure:"lastActivityStatus"`
					ObjectID           string `mapstructure:"objectId"`
					ObjectName         string `mapstructure:"objectName"`
					ObjectType         string `mapstructure:"objectType"`
					Severity           string `mapstructure:"severity"`
					Progress           string `mapstructure:"progress"`
					Cluster            struct {
						ID   string `mapstructure:"id"`
						Name string `mapstructure:"name"`
					} `mapstructure:"cluster"`
					ActivityConnection struct {
						Nodes []struct {
							ID      string `mapstructure:"id"`
							Message string `mapstructure:"message"`
							Time    string `mapstructure:"time"`
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

type PolarisEventsEdge struct {
	Node struct {
		ID                 int    `mapstructure:"id"`
		Fid                string `mapstructure:"fid"`
		ActivitySeriesID   string `mapstructure:"activitySeriesId"`
		LastUpdated        string `mapstructure:"lastUpdated"`
		LastActivityType   string `mapstructure:"lastActivityType"`
		LastActivityStatus string `mapstructure:"lastActivityStatus"`
		ObjectID           string `mapstructure:"objectId"`
		ObjectName         string `mapstructure:"objectName"`
		ObjectType         string `mapstructure:"objectType"`
		Severity           string `mapstructure:"severity"`
		Progress           string `mapstructure:"progress"`
		Cluster            struct {
			ID   string `mapstructure:"id"`
			Name string `mapstructure:"name"`
		} `mapstructure:"cluster"`
		ActivityConnection struct {
			Nodes []struct {
				ID      string `mapstructure:"id"`
				Message string `mapstructure:"message"`
				Time    string `mapstructure:"time"`
			} `mapstructure:"nodes"`
		} `mapstructure:"activityConnection"`
	} `mapstructure:"node"`
}

type ClusterIdByName struct {
	Data struct {
		ClusterConnection struct {
			Edges []struct {
				Node struct {
					ID   string `mapstructure:"id"`
					NAME string `mapstructure:"name"`
				} `mapstructure:"node"`
			} `mapstructure:"edges"`
			PageInfo struct {
				EndCursor   string `mapstructure:"endCursor"`
				HasNextPage bool   `mapstructure:"hasNextPage"`
			} `mapstructure:"pageInfo"`
		} `mapstructure:"clusterConnection"`
	} `mapstructure:"data"`
}

type EnableRadar struct {
	Data struct {
		EnableAutomaticFmdUpload struct {
			ClusterID string `mapstructure:"clusterId"`
			Enabled   bool   `mapstructure:"enabled"`
		} `mapstructure:"enableAutomaticFmdUpload"`
	} `mapstructure:"data"`
}
