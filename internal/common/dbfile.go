package common

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
)

func Load(fname string, tlist *TodoList) error {
	if f, err := ioutil.ReadFile(fname); err == nil {
		if _err := json.Unmarshal(f, tlist); _err != nil {
			return err
		}
	} else {
		return err
	}
	sort.Sort(tlist)
	return nil
}

func Store(fname string, tlist TodoList) error {
	if js, err := json.Marshal(tlist); err != nil {
		return err
	} else {
		ioutil.WriteFile(fname, js, os.ModePerm)
	}
	return nil
}
