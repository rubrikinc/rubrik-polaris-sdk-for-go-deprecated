package rubrikpolaris

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

func (c *Credentials) GetAllEvents(secondsTimeRange int, timeout ...int) (*AllEvents, error) {

	httpTimeout := httpTimeout(timeout)

	query, err := c.readQueryFile("AllEventsPerTimePeriod.graphql")
	if err != nil {
		return nil, err
	}

	variables := map[string]interface{}{}
	variables["timeAgo"] = time.Now().Add(time.Duration(secondsTimeRange*-1) * time.Second).UTC().Format(time.RFC3339)

	events, err := c.QueryWithVariables(query, variables, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse AllEvents
	mapErr := mapstructure.Decode(events, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}
	return &apiResponse, nil

}

func (c *Credentials) GetAllAuditLogByMinute(minuteTimeRange int, timeout ...int) (*AllAuditLog, error) {

	httpTimeout := httpTimeout(timeout)

	query, err := c.readQueryFile("AllAuditLogPerTimePeriod.graphql")
	if err != nil {
		return nil, err
	}

	variables := map[string]interface{}{}
	variables["timeAgo"] = time.Now().Add(time.Duration(minuteTimeRange*-1) * time.Minute).UTC().Format(time.RFC3339)

	eventLog, err := c.QueryWithVariables(query, variables, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse AllAuditLog
	mapErr := mapstructure.Decode(eventLog, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}
	return &apiResponse, nil

}

func (c *Credentials) GetEventDetails(activitySeriesID, clusterUUID string, timeout ...int) (*EventSeriesDetail, error) {

	httpTimeout := httpTimeout(timeout)

	query, err := c.readQueryFile("AllEventDetails.graphql")
	if err != nil {
		return nil, err
	}

	variables := map[string]interface{}{}
	variables["activitySeriesId"] = activitySeriesID
	variables["clusterUuid"] = clusterUUID

	eventDetail, err := c.QueryWithVariables(query, variables, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse EventSeriesDetail
	mapErr := mapstructure.Decode(eventDetail, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}
	return &apiResponse, nil

}
