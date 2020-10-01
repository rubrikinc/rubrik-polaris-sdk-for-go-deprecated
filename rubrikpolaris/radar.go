package rubrikpolaris

import (
	"time"
)

// RadarEvents24Hours returns the number of Radar events that occured in the last 24 hours
func (c *Credentials) RadarEvents24Hours(timeout ...int) (float64, error) {

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
