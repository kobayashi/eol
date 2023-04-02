package commands

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kobayashi/eol/pkg/api"
	"github.com/spf13/cobra"
)

// ProjectArgs checks args counts
func ProjectArgs(args []string) error {
	if len(args) < 1 {
		return errors.New("requires a project name")
	} else if len(args) > 2 {
		return errors.New("not allowed multiple args")
	}
	return nil
}

var outFormat = []string{
	"markdown",
	"csv",
	"html",
}

func checkFormat(format string) error {
	if format == "" {
		return nil
	}
	for _, v := range outFormat {
		if format == v {
			return nil
		}
	}
	return errors.New("input format is not allow. choose from markdown, csv, or html")
}

type tableFormat struct {
	table  table.Writer
	format string
}

// RunGetProject calls api & show data table
func RunGetProject(cmd *cobra.Command, args []string) error {
	format, err := cmd.Flags().GetString("format")
	if err != nil {
		return err
	}
	err = checkFormat(format)
	if err != nil {
		return err
	}
	tf := tableFormat{
		table:  table.NewWriter(),
		format: format,
	}
	packageName := args[0]
	c := api.NewHTTPClient()
	ctx := context.Background()
	if len(args) == 1 {
		res, err := c.GetProjectCycleList(ctx, packageName)
		if err != nil {
			return err
		}
		tc, err := createContentList(res)
		if err != nil {
			return err
		}
		tf.display(tc)
	} else {
		packageVersion := args[1]
		res, err := c.GetProjectCycle(ctx, packageName, packageVersion)
		if err != nil {
			return err
		}
		tc, err := createContent(res)
		if err != nil {
			return err
		}
		tf.display(tc)
	}
	return nil
}

func createContentList(cycles api.CycleList) (*tableContenet, error) {
	header := tableHeader{}
	rs := []table.Row{}
	for _, p := range cycles {
		row := table.Row{}
		if p.Cycle.S != nil {
			h := "Cycle"
			exist := header.isHeaderExist(h)
			if !exist {
				header.Row = append(header.Row, h)
			}
			row = append(row, *p.Cycle.S)
		} else if p.Cycle.I != nil {
			h := "Cycle"
			exist := header.isHeaderExist(h)
			if !exist {
				header.Row = append(header.Row, h)
			}
			row = append(row, *p.Cycle.I)
		}
		if p.CycleShortHand.S != nil {
			h := "Codename"
			exist := header.isHeaderExist(h)
			if !exist {
				header.Row = append(header.Row, h)
			}
			row = append(row, *p.CycleShortHand.S)
		} else if p.CycleShortHand.I != nil {
			h := "Codename"
			exist := header.isHeaderExist(h)
			if !exist {
				header.Row = append(header.Row, h)
			}
			row = append(row, *p.CycleShortHand.I)
		}
		if p.EOL.S != nil {
			eol, err := setColorDate(*p.EOL.S)
			if err != nil {
				return nil, err
			}
			h := "EOL"
			exist := header.isHeaderExist(h)
			if !exist {
				header.Row = append(header.Row, h)
			}
			row = append(row, *eol)
		} else if p.EOL.B != nil {
			var eol string = setColorBool(*p.EOL.B)
			h := "EOL"
			exist := header.isHeaderExist(h)
			if !exist {
				header.Row = append(header.Row, h)
			}
			row = append(row, eol)
		}
		if p.ReleaseDate != nil {
			h := "Release Date"
			exist := header.isHeaderExist(h)
			if !exist {
				header.Row = append(header.Row, h)
			}
			row = append(row, *p.ReleaseDate)
		}
		if p.Support.S != nil {
			h := "Support"
			exist := header.isHeaderExist(h)
			if !exist {
				header.Row = append(header.Row, h)
			}
			row = append(row, *p.Support.S)
		} else if p.Support.B != nil {
			h := "Support"
			exist := header.isHeaderExist(h)
			if !exist {
				header.Row = append(header.Row, h)
			}
			row = append(row, *p.Support.B)
		}
		if p.Latest != nil {
			h := "Latest"
			exist := header.isHeaderExist(h)
			if !exist {
				header.Row = append(header.Row, h)
			}
			row = append(row, *p.Latest)
		}
		if p.Discontinued.S != nil {
			h := "Discontinued"
			fmt.Println(h)
			exist := header.isHeaderExist(h)
			if !exist {
				header.Row = append(header.Row, "Discontinued")
			}
			row = append(row, *p.Discontinued.S)
		} else if p.Discontinued.B != nil {
			h := "Discontinued"
			fmt.Println(h)
			exist := header.isHeaderExist(h)
			if !exist {
				header.Row = append(header.Row, "Discontinued")
			}
			row = append(row, *p.Discontinued.B)
		}
		rs = append(rs, row)
	}
	return &tableContenet{header: header, rows: rs}, nil
}

func createContent(cycle *api.Cycle) (*tableContenet, error) {
	header := table.Row{}
	row := table.Row{}
	if cycle.Cycle.S != nil {
		header = append(header, "#")
		row = append(row, *cycle.Cycle.S)
	} else if cycle.Cycle.I != nil {
		header = append(header, "#")
		row = append(row, *cycle.Cycle.I)
	}
	if cycle.ReleaseDate != nil {
		header = append(header, "Release")
		row = append(row, *cycle.ReleaseDate)
	}
	if cycle.EOL.S != nil {
		eol, err := setColorDate(*cycle.EOL.S)
		if err != nil {
			return nil, err
		}
		header = append(header, "EOL")
		row = append(row, *eol)
	} else if cycle.EOL.B != nil {
		var eol string = setColorBool(*cycle.EOL.B)
		header = append(header, "EOL")
		row = append(row, eol)
	}
	if cycle.Latest != nil {
		header = append(header, "Latest")
		row = append(row, *cycle.Latest)
	}
	if cycle.Link != nil {
		header = append(header, "Link")
		row = append(row, *cycle.Link)
	}
	if cycle.LTS != nil {
		header = append(header, "LTS")
		row = append(row, *cycle.LTS)
	}
	if cycle.Support.S != nil {
		header = append(header, "Support")
		row = append(row, *cycle.Support.S)
	} else if cycle.Support.B != nil {
		header = append(header, "Support")
		row = append(row, *cycle.Support.B)
	}
	if cycle.CycleShortHand.S != nil {
		header = append(header, "CycleShortHand")
		row = append(row, *cycle.CycleShortHand.S)
	} else if cycle.CycleShortHand.I != nil {
		header = append(header, "CycleShortHand")
		row = append(row, *cycle.CycleShortHand.I)
	}
	if cycle.Discontinued.S != nil {
		header = append(header, "Discontinued")
		row = append(row, *cycle.Discontinued.S)
	} else if cycle.Discontinued.B != nil {
		header = append(header, "Discontinued")
		row = append(row, *cycle.Discontinued.B)
	}
	th := tableHeader{header}
	tcl := tableContenet{
		header: th,
		rows:   []table.Row{row},
	}
	return &tcl, nil
}

type tableHeader struct {
	table.Row
}

type tableContenet struct {
	header tableHeader
	rows   []table.Row
}

func (h *tableHeader) isHeaderExist(s string) bool {
	for _, v := range h.Row {
		if v == s {
			return true
		}
	}
	return false
}

func (tcl *tableContenet) display(format string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(tcl.header.Row)
	t.AppendRows(tcl.rows)
	t.Render()
}

func (tf *tableFormat) display(tcl *tableContenet) {
	tf.table.SetOutputMirror(os.Stdout)
	tf.table.AppendHeader(tcl.header.Row)
	tf.table.AppendRows(tcl.rows)

	switch tf.format {
	case "":
		tf.table.Render()
	case "markdown":
		tf.table.RenderMarkdown()
	case "csv":
		tf.table.RenderCSV()
	case "html":
		tf.table.RenderHTML()
	}
}

var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()

func setColorDate(eolDate string) (*string, error) {
	now := time.Now()
	t, err := time.Parse("2006-01-02", eolDate)
	if err != nil {
		return nil, err
	}
	inSixMonth := now.AddDate(0, 6, 0)
	var res string
	if t.Before(now) {
		res = red(eolDate)
	} else if t.Before(inSixMonth) {
		res = yellow(eolDate)
	} else {
		res = green(eolDate)
	}
	return &res, nil
}

func setColorBool(isEOL bool) string {
	var res string
	if isEOL {
		res = red("yes")
	} else {
		res = green("no")
	}
	return res
}
