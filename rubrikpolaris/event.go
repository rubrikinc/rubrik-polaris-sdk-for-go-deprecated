package rubrikpolaris

import "time"

func (c *Credentials) GetAllEvents(secondsTimeRange int, timeout ...int) (interface{}, error) {

	httpTimeout := httpTimeout(timeout)

	query, err := c.readQueryFile("AllEventsPerTimePeriod.graphql")
	if err != nil {
		return 0, err
	}

	variables := map[string]interface{}{}
	variables["timeAgo"] = time.Now().Add(time.Duration(secondsTimeRange*-1) * time.Second).UTC().Format(time.RFC3339)

	events, err := c.QueryWithVariables(query, variables, httpTimeout)
	if err != nil {
		return 0, err
	}
	return events, nil

}

func (c *Credentials) GetAllAuditLogByMinute(minuteTimeRange int, timeout ...int) (interface{}, error) {

	httpTimeout := httpTimeout(timeout)

	query, err := c.readQueryFile("AllAuditLogPerTimePeriod.graphql")
	if err != nil {
		return 0, err
	}

	variables := map[string]interface{}{}
	variables["timeAgo"] = time.Now().Add(time.Duration(minuteTimeRange*-1) * time.Minute).UTC().Format(time.RFC3339)

	eventLog, err := c.QueryWithVariables(query, variables, httpTimeout)
	if err != nil {
		return 0, err
	}
	return eventLog, nil

}
