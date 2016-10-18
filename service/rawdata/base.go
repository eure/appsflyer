package rawdata

import (
	"os"
	"path"

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
	return entities, backupIfNeed(client, body)
}

func getEachReport(endpoint string, client *dispatcher.Client, f func(report rawdata.Report)) error {
	body, err := client.DispatchGetRequest(endpoint)
	if err != nil {
		return err
	}
	if err := csv.Parse(string(body), rawdata.Report{}, func(v interface{}) {
		f(v.(rawdata.Report))
	}); err != nil {
		return err
	}
	return backupIfNeed(client, body)
}

func backupIfNeed(client *dispatcher.Client, body []byte) error {
	buckupOption := client.BuckupOption
	if buckupOption == nil {
		return nil
	}
	fileName := path.Join(os.TempDir(), client.GetCSVFileNameByDateRange())
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := csv.Write(string(body), file); err != nil {
		return err
	}
	if err := buckupOption.Do(file); err != nil {
		return err
	}
	return os.Remove(fileName)
}
