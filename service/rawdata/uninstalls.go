package rawdata

import (
	"github.com/eure/appsflyer/dispatcher"
	"github.com/eure/appsflyer/model/rawdata"
)

const endpointUninstallReport = "uninstall_events_report/v5"

func GetUninstallReports(client *dispatcher.Client) ([]rawdata.Report, error) {
	return getReports(endpointInstallReport, client)
}

func GetEachUninstallReport(client *dispatcher.Client, f func(report rawdata.Report)) error {
	return getEachReport(endpointInstallReport, client, f)
}