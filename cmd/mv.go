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
	"strconv"

	. "github.com/nocd5/todo/internal/common"
	"github.com/spf13/cobra"
)

// mvCmd represents the mv command
var mvCmd = &cobra.Command{
	Use:   "mv 1 42",
	Short: "Change the id of given todo",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Fprintln(os.Stderr, "todo: Number of given id should be two")
			return
		}

		from, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		to, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		if _, err := todoList.FindByID(int(from)); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		if _, err := todoList.FindByID(int(to)); err == nil {
			errmsg := fmt.Sprintf("todo: Destination id \"%d\" is already exists", to)
			fmt.Fprintln(os.Stderr, errmsg)
			return
		} else if to == 0 {
			errmsg := fmt.Sprintf("todo: Destination id \"%d\" is invalid", to)
			fmt.Fprintln(os.Stderr, errmsg)
			return
		}

		if idx, err := todoList.FindByID(int(from)); err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			todoList[idx].ID = to
			Store(dbFile, todoList)
		}
	},
}

func init() {
	RootCmd.AddCommand(mvCmd)
}
