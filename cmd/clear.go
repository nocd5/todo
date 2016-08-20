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

	. "github.com/nocd5/todo/internal/common"
	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Destroy todo items",
	Run: func(cmd *cobra.Command, args []string) {
		if *clearDone {
			todoList = todoList.Remove(func(t Todo) bool { return t.Status == "done" })
		} else if *clearAll {
			todoList = todoList[:0]
		} else {
			fmt.Fprintln(os.Stderr, "todo: Please add flag either \"-d, --done\" or \"-a, --all\"")
			return
		}
		Store(dbFile, todoList)
	},
}

var clearAll *bool
var clearDone *bool

func init() {
	RootCmd.AddCommand(clearCmd)

	clearAll = clearCmd.Flags().BoolP("all", "a", false, "Clear completed and pending todo items")
	clearDone = clearCmd.Flags().BoolP("done", "d", false, "Clear completed todo items")
}
