package dates

import (
	"testing"
	"time"
)

func TestDayIterate(t *testing.T) {
	begin, e := time.Parse("2006-01-02", "2015-01-01")
	if e != nil {
		t.Error(e)
		return
	}
	days, e := DaysInMonth(begin)
	if e != nil {
		t.Error(e)
		return
	}
	compare := map[string]string{
		"2015-01-01": "",
		"2015-01-02": "",
		"2015-01-03": "",
		"2015-01-04": "",
		"2015-01-05": "",
		"2015-01-06": "",
		"2015-01-07": "",
		"2015-01-08": "",
		"2015-01-09": "",
		"2015-01-10": "",
		"2015-01-11": "",
		"2015-01-12": "",
		"2015-01-13": "",
		"2015-01-14": "",
		"2015-01-15": "",
		"2015-01-16": "",
		"2015-01-17": "",
		"2015-01-18": "",
		"2015-01-19": "",
		"2015-01-20": "",
		"2015-01-21": "",
		"2015-01-22": "",
		"2015-01-23": "",
		"2015-01-24": "",
		"2015-01-25": "",
		"2015-01-26": "",
		"2015-01-27": "",
		"2015-01-28": "",
		"2015-01-29": "",
		"2015-01-30": "",
		"2015-01-31": "",
	}

	if len(days) != len(compare) {
		t.Errorf("Missing dates?")
	}
	for _, day := range days {
		fmt := day.Format("2006-01-02")
		_, ok := compare[fmt]
		if !ok {
			t.Errorf("Date should not exist: %s", fmt)
		}
	}
}

TestBrokenDayIterate(t *testing.T) {
	begin, e := time.Parse("2006-01-02", "2015-01-29")
	if e != nil {
		t.Error(e)
		return
	}
	days, e := DaysInMonth(begin)
	if e != nil {
		t.Error(e)
		return
	}
	compare := map[string]string{
		"2015-01-29": "",
		"2015-01-30": "",
		"2015-01-31": "",
	}

	if len(days) != len(compare) {
		t.Errorf("Missing dates?")
	}
	for _, day := range days {
		fmt := day.Format("2006-01-02")
		_, ok := compare[fmt]
		if !ok {
			t.Errorf("Date should not exist: %s", fmt)
		}
	}
}