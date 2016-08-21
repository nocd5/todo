// Copyright c 2016 nocd5
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mattn/go-colorable"
	. "github.com/nocd5/todo/internal/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var lsAll *bool
var lsDone *bool

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls [[~]@tag]",
	Short: "Print todo items",
	Run: func(cmd *cobra.Command, args []string) {
		out := colorable.NewColorable(os.Stdout)
		checked := getAnsiColor(FgGreenBold, BgDefault) + viper.GetString("Checked") + getAnsiColor(FgDefault, BgDefault)
		unchecked := getAnsiColor(FgRedBold, BgDefault) + viper.GetString("Unchecked") + getAnsiColor(FgDefault, BgDefault)
		var mark string

		var filteredList []Todo
		for _, todo := range filter(todoList, args) {
			switch todo.Status {
			case "done":
				if !*lsAll && !*lsDone {
					continue
				}
			case "pending":
				if *lsDone {
					continue
				}
			}
			filteredList = append(filteredList, todo)
		}

		if filteredList == nil || len(filteredList) < 0 {
			if os.Getenv("TODO_FORMAT") == "pretty" {
				fmt.Fprintf(out, "\n  %stodo:%s There are no todo items.\n\n",
					getAnsiColor(FgBlueBold, BgDefault), getAnsiColor(FgDefault, BgDefault))
			}
			return
		}

		maxidlen := 0
		maxelaplen := 0
		if os.Getenv("TODO_FORMAT") != "mini" {
			for _, todo := range filteredList {
				maxelaplen = Max(maxelaplen, len(getElapsedString(todo.Modified)))
				maxidlen = Max(maxidlen, len(fmt.Sprintf("%d", todo.ID)))
			}
		}

		if os.Getenv("TODO_FORMAT") != "mini" {
			out.Write([]byte("\n"))
		}
		for _, todo := range filteredList {
			switch todo.Status {
			case "done":
				mark = checked
			case "pending":
				mark = unchecked
			}
			switch os.Getenv("TODO_FORMAT") {
			case "pretty":
				elap := getAnsiColor(FgBlackBold, BgDefault) + "(" + getElapsedString(todo.Modified) + ")" + getAnsiColor(FgDefault, BgDefault)
				fmt.Fprintf(out, "    |  %s%d.  %s%s[ %s ]  %s%s  %s\n",
					getAnsiColor(FgYellowBold, BgDefault), todo.ID, getAnsiColor(FgDefault, BgDefault), strings.Repeat(" ", maxidlen-len(fmt.Sprintf("%d", todo.ID))),
					mark,
					elap, strings.Repeat(" ", maxelaplen-len(getElapsedString(todo.Modified))),
					todo.Desc)
			case "mini":
				fmt.Fprintf(out, "%d. %s\n", todo.ID, todo.Desc)
			default:
				fmt.Fprintf(out, "     %d.%s  %s   %s\n", todo.ID, strings.Repeat(" ", maxidlen-len(fmt.Sprintf("%d", todo.ID))), mark, todo.Desc)
			}
		}
		if os.Getenv("TODO_FORMAT") != "mini" {
			out.Write([]byte("\n"))
		}
	},
}

func getAnsiColor(fg, bg string) string {
	return "\033[" + fg + ";" + bg + "m"
}

func getElapsedString(modified string) string {
	elap := "NaN days ago"
	if tm, err := time.Parse(time.RFC3339, modified); err == nil {
		diff := time.Since(tm)
		if diff.Hours() < 24 {
			if diff.Hours() < 1 {
				if diff.Minutes() < 1 {
					if diff.Seconds() < 1 {
						elap = fmt.Sprintf("%d ms", int((diff / time.Millisecond).Nanoseconds()))
					} else {
						elap = fmt.Sprintf("%d seconds", int(diff.Seconds())+1)
					}
				} else {
					elap = fmt.Sprintf("%d minutes", int(diff.Minutes())+1)
				}
			} else {
				elap = fmt.Sprintf("%d hours", int(diff.Hours())+1)
			}
		} else {
			elap = fmt.Sprintf("%d days", int(diff.Hours()/24)+1)
		}
		elap += " ago"
	}
	return elap
}

func filter(list []Todo, words []string) []Todo {
	dest := list
	for _, w := range words {
		temp := dest
		dest = nil
		for _, l := range temp {
			if []rune(w)[0] == '~' && !strings.Contains(l.Desc, w[1:]) {
				dest = append(dest, l)
			} else if strings.Contains(l.Desc, w) {
				dest = append(dest, l)
			}
		}
	}
	return dest
}

func init() {
	RootCmd.AddCommand(lsCmd)

	lsAll = lsCmd.Flags().BoolP("all", "a", false, "Print completed and pending todo items")
	lsDone = lsCmd.Flags().BoolP("done", "d", false, "Print completed todo items")
}
