package dispatcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/eure/appsflyer/util"
)

type (
	RequiredParameter struct {
		APIToken string
		AppID    string
		FromDate string
		ToDate   string
	}
	OptionalParameter struct {
		Category    string
		MediaSource string
		Reattr      string
	}
	BuckupS3Option struct {
		AccessKeyID     string
		SecretAccessKey string
		Region          string
		Bucket          string
	}
	Client struct {
		HTTPClient *http.Client

		APIBaseURL           string
		APIRequiredParameter RequiredParameter
		APIOptionalParameter OptionalParameter

		BuckupS3Option *BuckupS3Option
	}
)

const (
	defaultAPIBaseURL = "https://hq.appsflyer.com"
)

func NewClient(appID, fromDate, toDate string) *Client {
	return NewClientWithParam(util.GetAPIToken(), appID, fromDate, toDate)
}

func NewClientWithParam(apiToken, appID, fromDate, toDate string) *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		APIBaseURL: defaultAPIBaseURL,
		APIRequiredParameter: RequiredParameter{
			APIToken: apiToken,
			AppID:    appID,
			FromDate: fromDate,
			ToDate:   toDate,
		},
	}
}

func (c *Client) SetOptionalParameter(p OptionalParameter) {
	c.APIOptionalParameter = p
}

func (c *Client) SetBuckupS3Option(o BuckupS3Option) {
	c.BuckupS3Option = &o
}

func (c *Client) DispatchGetRequest(endpoint string) ([]byte, error) {
	u, err := url.Parse(c.APIBaseURL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join("export", c.APIRequiredParameter.AppID, endpoint)
	urlString := u.String()

	values := url.Values{}

	// Required parameters
	values.Set("api_token", c.APIRequiredParameter.APIToken)
	values.Set("from", c.APIRequiredParameter.FromDate)
	values.Set("to", c.APIRequiredParameter.ToDate)

	// Optional parameters
	if c.APIOptionalParameter.Category != "" {
		values.Set("category", c.APIOptionalParameter.Category)
	}
	if c.APIOptionalParameter.MediaSource != "" {
		values.Set("media_source", c.APIOptionalParameter.MediaSource)
	}
	if c.APIOptionalParameter.Reattr != "" {
		values.Set("reattr", c.APIOptionalParameter.Reattr)
	}

	resp, err := c.HTTPClient.Get(urlString + "?" + values.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Return error when status code less than 200 or equal more than 300
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("StatusCode = %d, Message = %s ", resp.StatusCode, string(body))
	}
	return body, nil
}

func (c *Client) GetCSVFileNameByDateRange() string {
	return fmt.Sprintf("appsflyer[%s~%s].csv", c.APIRequiredParameter.FromDate, c.APIRequiredParameter.ToDate)
}
