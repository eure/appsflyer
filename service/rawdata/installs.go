package rawdata

import (
	"github.com/eure/appsflyer/dispatcher"
	"github.com/eure/appsflyer/model/rawdata"
)

const endpointInstallReport = "installs_report/v5"
const endpointOrganicInstallReport = "organic_installs_report/v5"

func GetInstallReports(client *dispatcher.Client) ([]rawdata.Report, error) {
	return getReports(endpointInstallReport, client)
}

func GetEachInstallReport(client *dispatcher.Client, f func(report rawdata.Report)) error {
	return getEachReport(endpointInstallReport, client, f)
}

func GetOrganicInstallReports(client *dispatcher.Client) ([]rawdata.Report, error) {
	return getReports(endpointOrganicInstallReport, client)
}

func GetEachOrganicInstallReport(client *dispatcher.Client, f func(report rawdata.Report)) error {
	return getEachReport(endpointOrganicInstallReport, client, f)
}
