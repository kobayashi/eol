package commands

import (
	"encoding/json"
	"testing"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kobayashi/eol/pkg/api"
)

func TestTableHeader(t *testing.T) {
	th := tableHeader{}
	th.Row = table.Row{"a", "b", "c", "d"}
	if !th.isHeaderExist("a") {
		t.Error("No header")
	}
	if th.isHeaderExist("e") {
		t.Error("Unexpected")
	}
}

func TestCreateContentList(t *testing.T) {
	s := `[
  {
    "cycle": "21.04",
    "cycleShortHand": "HirsuteHippo",
    "lts": false,
    "release": "2021-04-22",
    "support": "2022-01-01",
    "eol": "2022-01-01",
    "latest": "21.04",
    "link": "https://wiki.ubuntu.com/HirsuteHippo/ReleaseNotes/"
  },
  {
    "cycle": 10,
    "cycleShortHand": 10,
    "lts": false,
    "release": "2020-10-22",
    "support": true,
    "eol": true,
    "latest": "20.10",
    "link": "https://wiki.ubuntu.com/GroovyGorilla/ReleaseNotes/"
  }
]`
	cl := api.CycleList{}
	json.Unmarshal([]byte(s), &cl)
	tl, err := createContentList(cl)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if len(tl.header.Row) != 5 {
		t.Errorf("Unexpected header count: %d", len(tl.header.Row))
	}
	if len(tl.rows) != 2 {
		t.Errorf("Unexpected row count: %d", len(tl.rows))
	}
}

func TestCreateContent(t *testing.T) {
	testCases := []string{
		`{
			"release": "2020-10-05",
			"eol": "2025-10-05",
			"latest": "3.9.5",
			"link": "https://www.python.org/downloads/release/python-395/"
		}`,
		`{
			"cycleShortHand": "HirsuteHippo",
			"lts": false,
			"release": "2021-04-22",
			"support": "2022-01-01",
			"eol": "2022-01-01",
			"latest": "21.04",
			"link": "https://wiki.ubuntu.com/HirsuteHippo/ReleaseNotes/"
		}`,
	}
	for i, s := range testCases {
		c := api.Cycle{}
		json.Unmarshal([]byte(s), &c)
		tc, err := createContent(&c)
		if err != nil {
			t.Errorf("error: %s", err)
		}
		rowCount := []int{4, 7}

		if len(tc.rows[0]) != rowCount[i] {
			t.Errorf("Unexpected count: %d", len(tc.rows[0]))
		}
	}
}
