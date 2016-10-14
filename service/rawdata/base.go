package rawdata

import (
	"os"
	"path"

	"github.com/eure/appsflyer/dispatcher"
	"github.com/eure/appsflyer/model/rawdata"
	"github.com/eure/appsflyer/util/csv"
	"github.com/eure/appsflyer/util/s3"
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
	return entities, backupS3IfNeed(client, body)
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
	return backupS3IfNeed(client, body)
}

func backupS3IfNeed(client *dispatcher.Client, body []byte) error {
	s3Option := client.BuckupS3Option
	if s3Option == nil {
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
	s3Uploader := s3.Uploader{
		AccessKeyID:     s3Option.AccessKeyID,
		SecretAccessKey: s3Option.SecretAccessKey,
		Region:          s3Option.Region,
		Bucket:          s3Option.Bucket,
	}
	if err := s3Uploader.Do(file); err != nil {
		return err
	}
	return os.Remove(fileName)
}
