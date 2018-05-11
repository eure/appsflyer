package rawdata

import (
	"github.com/eure/appsflyer/dispatcher"
	"github.com/eure/appsflyer/model/rawdata"
	"github.com/eure/appsflyer/util/csv"
)

func getReports(endpoint string, client *dispatcher.Client) ([]rawdata.Report, error) {
	body, err := client.DispatchGetRequest(endpoint)
	if err != nil {
		return nil, err
	}
	var entities []rawdata.Report
	if err := csv.Parse(string(body), rawdata.Report{}, func(v interface{}) {
		entities = append(entities, v.(rawdata.Report))
	}); err != nil {
		return nil, err
	}
	return entities, nil
}

func getEachReport(endpoint string, client *dispatcher.Client, f func(report rawdata.Report)) error {
	body, err := client.DispatchGetRequest(endpoint)
	if err != nil {
		return err
	}
	return csv.Parse(string(body), rawdata.Report{}, func(v interface{}) {
		f(v.(rawdata.Report))
	})
}
