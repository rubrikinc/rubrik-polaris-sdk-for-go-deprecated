package rubrikpolaris

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

// RadarEventsLast24Hours returns the number of Radar events that occured in the last 24 hours
func (c *Credentials) GetRadarEventsLast24Hours(timeout ...int) (float64, error) {

	httpTimeout := httpTimeout(timeout)

	query, err := c.readQueryFile("RadarEventsPerTimePeriod.graphql")
	if err != nil {
		return 0, err
	}

	variables := map[string]interface{}{}
	variables["timeAgo"] = time.Now().Add(-24 * time.Hour).UTC().Format(time.RFC3339)

	radar, err := c.QueryWithVariables(query, variables, httpTimeout)
	if err != nil {
		return 0, err
	}
	return radar.(map[string]interface{})["data"].(map[string]interface{})["activitySeriesConnection"].(map[string]interface{})["count"].(float64), nil

}

// RadarEventsLast30Days returns the number of Radar events that occured in the last 30 days
func (c *Credentials) GetRadarEventsLast30Days(timeout ...int) (float64, error) {

	httpTimeout := httpTimeout(timeout)

	query, err := c.readQueryFile("RadarEventsPerTimePeriod.graphql")
	if err != nil {
		return 0, err
	}

	variables := map[string]interface{}{}
	variables["timeAgo"] = time.Now().Add(-720 * time.Hour).UTC().Format(time.RFC3339)

	radar, err := c.QueryWithVariables(query, variables, httpTimeout)
	if err != nil {
		return 0, err
	}
	return radar.(map[string]interface{})["data"].(map[string]interface{})["activitySeriesConnection"].(map[string]interface{})["count"].(float64), nil

}

// RadarEventsLastYear returns the number of Radar events that occured in the last year
func (c *Credentials) GetRadarEventsLastYear(timeout ...int) (float64, error) {

	httpTimeout := httpTimeout(timeout)

	query, err := c.readQueryFile("RadarEventsPerTimePeriod.graphql")
	if err != nil {
		return 0, err
	}

	variables := map[string]interface{}{}
	variables["timeAgo"] = time.Now().Add(-8760 * time.Hour).UTC().Format(time.RFC3339)

	radar, err := c.QueryWithVariables(query, variables, httpTimeout)
	if err != nil {
		return 0, err
	}
	return radar.(map[string]interface{})["data"].(map[string]interface{})["activitySeriesConnection"].(map[string]interface{})["count"].(float64), nil

}

// GetRadarEnabledClusters returns the name of each Rubrik cluster with Radar enabled map to its ID value.
func (c *Credentials) GetRadarEnabledClusters(timeout ...int) (map[string]string, error) {

	httpTimeout := httpTimeout(timeout)

	query, err := c.readQueryFile("RadarEnabledClusters.graphql")
	if err != nil {
		return nil, err
	}

	radarEnabledClustersQuery, err := c.Query(query, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse RadarEnabledClusters
	mapErr := mapstructure.Decode(radarEnabledClustersQuery, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	enabledClusters := make(map[string]string)

	for _, cluster := range apiResponse.Data.RadarClusterConnection.Nodes {
		if cluster.LambdaConfig != nil {
			enabledClusters[cluster.Name] = cluster.LambdaConfig.(map[string]interface{})["clusterId"].(string)
		}

	}

	return enabledClusters, nil

}
