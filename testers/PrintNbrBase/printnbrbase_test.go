package student

import (
	"io/ioutil"
	"os"
	"testing"

)

func TestPrintNbrBase(t *testing.T) {
	tables := []struct {
		nbr  int
		base string
		res  string
	}{
		{125, "0123456789", "125"},
		{125, "01234567890", "NV"},
		{-125, "01", "-1111101"},
		{125, "-01", "NV"},
		{125, "0123456789ABCDEF", "7D"},
		{-125, "choumi", "-uoi"},
		{123, "asdf", "sfdf"},
		{125, "aa", "NV"},
		{125, "a", "NV"},
		//values from audit
		{919617, "01", "11100000100001000001"},
		{753639, "CHOUMIisDAcat!", "HIDAHI"},
		{-174336, "CHOUMIisDAcat!", "-MssiD"},
		{-661165, "1", "NV"},
		{-861737, "Zone01Zone01", "NV"},
		{125, "0123456789ABCDEF", "7D"},
		{-125, "choumi", "-uoi"},
		{125, "-ab", "NV"},
		{-9223372036854775808, "0123456789", "-9223372036854775808"},
	}
	for _, table := range tables {
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		PrintNbrBase(table.nbr, table.base)

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout

		total := string(out)
		if total != table.res {
			t.Errorf("PrintNbrBase of (%d, %s) was incorrect, got: %s, want: %s.", table.nbr, table.base, total, table.res)
		}
	}
}
