package csv

import (
	"io/ioutil"
	"testing"

	"github.com/eure/appsflyer/model/rawdata"
)

func TestParseCSV(t *testing.T) {
	res, err := ioutil.ReadFile("v5_sample_installs_demo.csv")
	if err != nil {
		t.Fatal(err)
	}
	if err := Parse(string(res), rawdata.Report{}, func(v interface{}) {
		r := v.(rawdata.Report)
		if r.AttributedTouchType == "" {
			t.Fatal("attribute is empty.")
		}
		if r.AttributedTouchTime == "" {
			t.Fatal("attribute is empty.")
		}
		t.Log(r)
	}); err != nil {
		t.Fatal(err)
	}
}
