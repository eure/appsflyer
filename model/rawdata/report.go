package rawdata

import (
	"strings"
	"time"

	utime "github.com/eure/appsflyer/util/time"
)

type Report struct {
	AttributedTouchType       string `json:"attributed_touch_type" csv:"Attributed Touch Type"`
	AttributedTouchTime       string `json:"attributed_touch_time" csv:"Attributed Touch Time"`
	InstallTime               string `json:"install_time" csv:"Install Time"`
	EventTime                 string `json:"event_time" csv:"Event Time"`
	EventName                 string `json:"event_name" csv:"Event Name"`
	EventValue                string `json:"event_value" csv:"Event Value"`
	EventRevenue              string `json:"event_revenue" csv:"Event Revenue"`
	EventRevenueCurrency      string `json:"event_revenue_currency" csv:"Event Revenue Currency"`
	EventRevenueUSD           string `json:"event_revenue_usd" csv:"Event Revenue USD"`
	EventSource               string `json:"event_source" csv:"Event Source"`
	IsReceiptValidated        string `json:"is_receipt_validated" csv:"Is Receipt Validated"`
	Partner                   string `json:"partner" csv:"Partner"`
	MediaSource               string `json:"media_source" csv:"Media Source"`
	Channel                   string `json:"channel" csv:"Channel"`
	Keywords                  string `json:"keywords" csv:"Keywords"`
	Campaign                  string `json:"campaign" csv:"Campaign"`
	CampaignID                string `json:"campaign_id" csv:"Campaign ID"`
	Adset                     string `json:"adset" csv:"Adset"`
	AdsetID                   string `json:"adset_id" csv:"Adset ID"`
	Ad                        string `json:"ad" csv:"Ad"`
	AdID                      string `json:"ad_id" csv:"Ad ID"`
	AdType                    string `json:"ad_type" csv:"Ad Type"`
	SiteID                    string `json:"site_id" csv:"Site ID"`
	SubSiteID                 string `json:"sub_site_id" csv:"Sub Site ID"`
	SubParam1                 string `json:"sub_param_1" csv:"Sub Param 1"`
	SubParam2                 string `json:"sub_param_2" csv:"Sub Param 2"`
	SubParam3                 string `json:"sub_param_3" csv:"Sub Param 3"`
	SubParam4                 string `json:"sub_param_4" csv:"Sub Param 4"`
	SubParam5                 string `json:"sub_param_5" csv:"Sub Param 5"`
	CostModel                 string `json:"cost_model" csv:"Cost Model"`
	CostValue                 string `json:"cost_value" csv:"Cost Value"`
	CostCurrency              string `json:"cost_currency" csv:"Cost Currency"`
	Contributor1Partner       string `json:"contributor_1_partner" csv:"Contributor 1 Partner"`
	Contributor1MediaSource   string `json:"contributor_1_media_source" csv:"Contributor 1 Media Source"`
	Contributor1Campaign      string `json:"contributor_1_campaign" csv:"Contributor 1 Campaign"`
	Contributor1TouchType     string `json:"contributor_1_touch_type" csv:"Contributor 1 Touch Type"`
	Contributor1TouchTime     string `json:"contributor_1_touch_time" csv:"Contributor 1 Touch Time"`
	Contributor2Partner       string `json:"contributor_2_partner" csv:"Contributor 2 Partner"`
	Contributor2MediaSource   string `json:"contributor_2_media_source" csv:"Contributor 2 Media Source"`
	Contributor2Campaign      string `json:"contributor_2_campaign" csv:"Contributor 2 Campaign"`
	Contributor2TouchType     string `json:"contributor_2_touch_type" csv:"Contributor 2 Touch Type"`
	Contributor2TouchTime     string `json:"contributor_2_touch_time" csv:"Contributor 2 Touch Time"`
	Contributor3Partner       string `json:"contributor_3_partner" csv:"Contributor 3 Partner"`
	Contributor3MediaSource   string `json:"contributor_3_media_source" csv:"Contributor 3 Media Source"`
	Contributor3Campaign      string `json:"contributor_3_campaign" csv:"Contributor 3 Campaign"`
	Contributor3TouchType     string `json:"contributor_3_touch_type" csv:"Contributor 3 Touch Type"`
	Contributor3TouchTime     string `json:"contributor_3_touch_time" csv:"Contributor 3 Touch Time"`
	Region                    string `json:"region" csv:"Region"`
	CountryCode               string `json:"country_code" csv:"Country Code"`
	State                     string `json:"state" csv:"State"`
	City                      string `json:"city" csv:"City"`
	PostalCode                string `json:"postal_code" csv:"Postal Code"`
	DMA                       string `json:"dma" csv:"DMA"`
	AdvertisingID             string `json:"advertising_id" csv:"Advertising ID"`
	IDFA                      string `json:"idfa" csv:"IDFA"`
	CustomerUserID            string `json:"customer_user_id" csv:"Customer User ID"`
	IP                        string `json:"ip" csv:"IP"`
	Wifi                      string `json:"wifi" csv:"WIFI"`
	Operator                  string `json:"operator" csv:"Operator"`
	Carrier                   string `json:"carrier" csv:"Carrier"`
	AppsflyerID               string `json:"appsflyer_id" csv:"AppsFlyer ID"`
	AndroidID                 string `json:"android_id" csv:"Android ID"`
	IMEI                      string `json:"imei" csv:"IMEI"`
	IDFV                      string `json:"idfv" csv:"IDFV"`
	Platform                  string `json:"platform" csv:"Platform"`
	DeviceType                string `json:"device_type" csv:"Device Type"`
	OSVersion                 string `json:"os_version" csv:"OS Version"`
	AppVersion                string `json:"app_version" csv:"App Version"`
	SDKVersion                string `json:"sdk_version" csv:"SDK Version"`
	AppID                     string `json:"app_id" csv:"App ID"`
	AppName                   string `json:"app_name" csv:"App Name"`
	BundleID                  string `json:"bundle_id" csv:"Bundle ID"`
	IsRetargeting             string `json:"is_retargeting" csv:"Is Retargeting"`
	Retargeting               string `json:"retargeting" csv:"Retargeting"`
	RetargetingConversionType string `json:"retargeting_conversion_type" csv:"Retargeting Conversion Type"`
	AttributionLookback       string `json:"attribution_lookback" csv:"Attribution Lookback"`
	ReengagementWindow        string `json:"reengagement_window" csv:"Reengagement Window"`
	IsPrimaryAttribution      string `json:"is_primary_attribution" csv:"Is Primary Attribution"`
	UserAgent                 string `json:"user_agent" csv:"User Agent"`
	HTTPReferrer              string `json:"http_referrer" csv:"HTTP Referrer"`
	OriginalURL               string `json:"original_url" csv:"Original URL"`
	Language                  string `json:"language" csv:"Language"`
}

func (r *Report) GetAttributedTouchTime() (time.Time, error) {
	return utime.ParseDateTimeFormat(r.AttributedTouchTime)
}

func (r *Report) GetInstallTime() (time.Time, error) {
	return utime.ParseDateTimeFormat(r.InstallTime)
}

func (r *Report) GetEventTime() (time.Time, error) {
	return utime.ParseDateTimeFormat(r.EventTime)
}

func (r *Report) GetIsRetargeting() bool {
	switch strings.ToLower(r.IsRetargeting) {
	case "false":
		return false
	case "true":
		return true
	}
	return false
}
