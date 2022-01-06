/*
Copyright © 2022 kobayashi <abok.1k@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kobayashi/eol/api"
	"github.com/spf13/cobra"
)

type tableHeader struct {
	table.Row
}

func (r *tableHeader) isHeaderExist(s string) bool {
	for _, v := range r.Row {
		if v == s {
			return true
		}
	}
	return false
}

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "A brief description of your command",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a project name")
		} else if len(args) > 2 {
			return errors.New("not allowed multiple args")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		pn := args[0]
		c := api.NewHTTPClient()
		ctx := context.Background()
		if len(args) == 1 {
			res, err := c.GetProjectCycleList(pn, ctx)
			if err != nil {
				log.Fatal(err)
			}
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			header := tableHeader{}
			rs := []table.Row{}
			for _, p := range res {
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
				if p.EOL.S != nil {
					red := color.New(color.FgRed).SprintFunc()
					green := color.New(color.FgGreen).SprintFunc()
					yellow := color.New(color.FgYellow).SprintFunc()
					now := time.Now()
					t, err := time.Parse("2006-01-02", *p.EOL.S)
					if err != nil {
						fmt.Println(err)
					}
					inSixMonth := now.AddDate(0, 6, 0)
					var eol string
					if t.Before(now) {
						eol = red(*p.EOL.S)
					} else if t.Before(inSixMonth) {
						eol = yellow(*p.EOL.S)
					} else {
						eol = green(*p.EOL.S)
					}
					h := "EOL"
					exist := header.isHeaderExist(h)
					if !exist {
						header.Row = append(header.Row, h)
					}
					row = append(row, fmt.Sprintf("%s", eol))
				} else if p.EOL.B != nil {
					red := color.New(color.FgRed).SprintFunc()
					green := color.New(color.FgGreen).SprintFunc()
					var eol string
					if *p.EOL.B {
						eol = red("yes")
					} else {
						eol = green("no")
					}
					h := "EOL"
					exist := header.isHeaderExist(h)
					if !exist {
						header.Row = append(header.Row, h)
					}
					row = append(row, fmt.Sprintf("%s", eol))
				}
				if p.Release != nil {
					h := "Release"
					exist := header.isHeaderExist(h)
					if !exist {
						header.Row = append(header.Row, h)
					}
					row = append(row, *p.Release)
				}
				if p.Latest != nil {
					h := "Latest"
					exist := header.isHeaderExist(h)
					if !exist {
						header.Row = append(header.Row, h)
					}
					row = append(row, *p.Latest)
				}
				// if p.Link != nil {
				// 	h := "Link"
				// 	exist := header.isHeaderExist(h)
				// 	if !exist {
				// 		header.Row = append(header.Row, h)
				// 	}
				// 	row = append(row, *p.Link)
				// }
				// if p.LTS != nil {
				// 	h := "LTS"
				// 	exist := header.isHeaderExist(h)
				// 	if !exist {
				// 		header.Row = append(header.Row, h)
				// 	}
				// 	row = append(row, *p.LTS)
				// }
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
				// if p.CycleShortHand != nil {
				// 	h := "CycleShortHanda"
				// 	exist := header.isHeaderExist(h)
				// 	if !exist {
				// 		header.Row = append(header.Row, h)
				// 	}
				// 	row = append(row, *p.CycleShortHand)
				// }
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
			t.AppendHeader(header.Row)
			t.AppendRows(rs)
			t.Render()
		} else {
			pv := args[1]
			res, err := c.GetProjectCycle(pn, pv, ctx)
			if err != nil {
				log.Fatal(err)
			}
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			header := table.Row{}
			row := table.Row{}

			if res.Cycle.S != nil {
				header = append(header, "#")
				row = append(row, *res.Cycle.S)
			} else if res.Cycle.I != nil {
				header = append(header, "#")
				row = append(row, *res.Cycle.I)
			}
			if res.Release != nil {
				header = append(header, "Release")
				row = append(row, *res.Release)
			}
			// if res.EOL != nil {
			// 	red := color.New(color.FgRed).SprintFunc()
			// 	green := color.New(color.FgGreen).SprintFunc()
			// 	yellow := color.New(color.FgYellow).SprintFunc()
			// 	now := time.Now()
			// 	t, err := time.Parse("2006-01-02", *res.EOL)
			// 	if err != nil {
			// 		fmt.Println(err)
			// 	}
			// 	inSixMonth := now.AddDate(0, 6, 0)
			// 	var eol string
			// 	if t.Before(now) {
			// 		eol = red(*res.EOL)
			// 	} else if t.Before(inSixMonth) {
			// 		eol = yellow(*res.EOL)
			// 	} else {
			// 		eol = green(*res.EOL)
			// 	}
			// 	header = append(header, "EOL")
			// 	row = append(row, fmt.Sprintf("%s", eol))
			// }
			if res.Latest != nil {
				header = append(header, "Latest")
				row = append(row, *res.Latest)
			}
			if res.Link != nil {
				header = append(header, "Link")
				row = append(row, *res.Link)
			}
			if res.LTS != nil {
				header = append(header, "LTS")
				row = append(row, *res.LTS)
			}
			if res.Support.S != nil {
				header = append(header, "Support")
				row = append(row, *res.Support.S)
			} else if res.Support.B != nil {
				header = append(header, "Support")
				row = append(row, *res.Support.B)
			}
			if res.CycleShortHand.S != nil {
				header = append(header, "CycleShortHand")
				row = append(row, *res.CycleShortHand.S)
			} else if res.CycleShortHand.I != nil {
				header = append(header, "CycleShortHand")
				row = append(row, *res.CycleShortHand.I)
			}
			if res.Discontinued.S != nil {
				header = append(header, "Discontinued")
				row = append(row, *res.Discontinued.S)
			} else if res.Discontinued.B != nil {
				header = append(header, "Discontinued")
				row = append(row, *res.Discontinued.B)
			}
			t.AppendHeader(header)
			t.AppendRow(row)
			t.Render()
		}

	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
