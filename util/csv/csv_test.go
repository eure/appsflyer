package csv

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/eure/appsflyer/model/rawdata"
)

func TestWriteCSV(t *testing.T) {
	const fName = "v5_sample_installs_demo_write_csv_test.csv"
	file, err := os.Create(fName)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(fName)
	res, err := ioutil.ReadFile("v5_sample_installs_demo.csv")
	if err != nil {
		t.Fatal(err)
	}
	if err := Write(string(res), file); err != nil {
		t.Fatal(err)
	}
}

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
