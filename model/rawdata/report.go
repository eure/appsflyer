package rawdata

import (
	"strings"
	"time"

	utime "github.com/eure/appsflyer/util/time"
)

type Report struct {
	AttributedTouchType string `json:"attributed_touch_type" csv:"Attributed Touch Type"`
	AttributedTouchTime string `json:"attributed_touch_time" csv:"Attributed Touch Time"`
	InstallTime         string `json:"install_time" csv:"Install Time"`
	EventTime           string `json:"event_time" csv:"Event Time"`
	EventName           string `json:"event_name" csv:"Event Name"`
	MediaSource         string `json:"media_source" csv:"Media Source"`
	Channel             string `json:"channel" csv:"Channel"`
	Campaign            string `json:"campaign" csv:"Campaign"`
	Ad                  string `json:"ad" csv:"Ad"`
	AdvertisingID       string `json:"advertising_id" csv:"Advertising ID"`
	IDFA                string `json:"idfa" csv:"IDFA"`
	CustomerUserID      string `json:"customer_user_id" csv:"Customer User ID"`
	IsRetargeting       string `json:"is_retargeting" csv:"Is Retargeting"`
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
