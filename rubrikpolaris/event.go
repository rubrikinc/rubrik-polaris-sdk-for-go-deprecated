package rubrikpolaris

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

const (
	// successiveEventQueryWaitPeriod is the time period in seconds to wait before
	// making the activitySeriesConnection query again
	successiveEventQueryWaitPeriod = 30
)

func (c *Credentials) GetAllEvents(secondsTimeRange int, timeout ...int) (*AllEvents, error) {

	httpTimeout := httpTimeout(timeout)

	query := c.readQueryFile("AllEventsPerTimePeriod.graphql")

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

func (c *Credentials) GetAllAuditLog(timeAgo string, timeout ...int) (*AllAuditLog, error) {

	httpTimeout := httpTimeout(timeout)

	query := c.readQueryFile("AllAuditLogPerTimePeriod.graphql")

	variables := map[string]interface{}{}
	variables["timeAgo"] = timeAgo

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

	query := c.readQueryFile("EventDetails.graphql")

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

func (c *Credentials) GetAllPolarisEvents(timeAgo string, timeout ...int) (*PolarisEvents, error) {
	return c.GetAllRscEventsForCluster(timeAgo, "00000000-0000-0000-0000-000000000000", timeout...)
}

func (c *Credentials) GetAllRscEventsForCluster(timeAgo string, clusterId string, timeout ...int) (*PolarisEvents, error) {

	httpTimeout := httpTimeout(timeout)

	if httpTimeout == 15 {
		httpTimeout = 300
	}

	query := c.readQueryFile("AllPolarisEventPerTimePeriod.graphql")

	variables := map[string]interface{}{}
	variables["timeAgo"] = timeAgo
	variables["clusterId"] = clusterId

	eventDetail, err := c.QueryWithVariables(query, variables, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse PolarisEvents
	mapErr := mapstructure.Decode(eventDetail, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	var additionalData []PolarisEventsEdge
	if apiResponse.Data.ActivitySeriesConnection.PageInfo.HasNextPage == true {

		variables["after"] = apiResponse.Data.ActivitySeriesConnection.PageInfo.EndCursor

		for {

			eventDetailPagination, err := c.QueryWithVariables(query, variables, httpTimeout)
			if err != nil {
				return nil, err
			}

			// Convert the API Response (map[string]interface{}) to a struct
			var apiResponsePagination PolarisEvents
			mapErr := mapstructure.Decode(eventDetailPagination, &apiResponsePagination)
			if mapErr != nil {
				return nil, mapErr
			}

			for _, data := range apiResponsePagination.Data.ActivitySeriesConnection.Edges {
				additionalData = append(additionalData, data)
			}

			if apiResponsePagination.Data.ActivitySeriesConnection.PageInfo.HasNextPage == false {
				break
			}

			variables["after"] = apiResponsePagination.Data.ActivitySeriesConnection.PageInfo.EndCursor

			// Add some sleep before successive activitySeriesConnection queries to ease load on the database
			time.Sleep(time.Duration(successiveEventQueryWaitPeriod) * time.Second)

		}

		for _, data := range additionalData {
			apiResponse.Data.ActivitySeriesConnection.Edges = append(apiResponse.Data.ActivitySeriesConnection.Edges, data)
		}

	}
	return &apiResponse, nil

}
