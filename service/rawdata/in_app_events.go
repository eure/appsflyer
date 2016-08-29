package rawdata

import (
	"github.com/eure/appsflyer/dispatcher"
	"github.com/eure/appsflyer/model/rawdata"
)

const endpointInAppEventsReport = "in_app_events_report/v5"

func GetInAppEventsReports(client *dispatcher.Client) ([]rawdata.Report, error) {
	return getReports(endpointInAppEventsReport, client)
}

func GetEachInAppEventsReport(client *dispatcher.Client, f func(report rawdata.Report)) error {
	return getEachReport(endpointInAppEventsReport, client, f)
}
