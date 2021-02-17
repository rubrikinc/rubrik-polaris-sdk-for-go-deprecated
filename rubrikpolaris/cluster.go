package rubrikpolaris

import (
	"github.com/mitchellh/mapstructure"
)

func (c *Credentials) GetCDMClusterIdByName(clusterNames []string, timeout ...int) ([]string, error) {

	httpTimeout := httpTimeout(timeout)

	query := c.readQueryFile("CDMClusterIdByName.graphql")

	variables := map[string]interface{}{}
	variables["clusterNames"] = clusterNames

	cdmClusters, err := c.QueryWithVariables(query, variables, httpTimeout)
	if err != nil {
		return nil, err
	}

	// Convert the API Response (map[string]interface{}) to a struct
	var apiResponse ClusterIdByName
	mapErr := mapstructure.Decode(cdmClusters, &apiResponse)
	if mapErr != nil {
		return nil, mapErr
	}

	var clusterIds []string

	for _, value := range apiResponse.Data.ClusterConnection.Edges {
		if contains(clusterNames, value.Node.NAME) {
			clusterIds = append(clusterIds, value.Node.ID)
		}

	}

	if apiResponse.Data.ClusterConnection.PageInfo.HasNextPage == true {
		variables["after"] = apiResponse.Data.ClusterConnection.PageInfo.EndCursor

		for {

			cdmClustersPagination, err := c.QueryWithVariables(query, variables, httpTimeout)
			if err != nil {
				return nil, err
			}

			// Convert the API Response (map[string]interface{}) to a struct
			var apiResponsePagination ClusterIdByName
			mapErr := mapstructure.Decode(cdmClustersPagination, &apiResponsePagination)
			if mapErr != nil {
				return nil, mapErr
			}

			for _, value := range apiResponsePagination.Data.ClusterConnection.Edges {

				if contains(clusterNames, value.Node.NAME) {
					clusterIds = append(clusterIds, value.Node.ID)
				}

			}

			if apiResponsePagination.Data.ClusterConnection.PageInfo.HasNextPage == false {
				break
			}

			variables["after"] = apiResponsePagination.Data.ClusterConnection.PageInfo.EndCursor

		}

	}

	return clusterIds, nil

}
