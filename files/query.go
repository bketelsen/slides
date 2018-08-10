package files

import (
	"io/ioutil"
	"time"
)

var epoch = time.Unix(1494505756, 0)

func LatestFileIn(path string) (latest string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return ""
	}
	latestTime := epoch
	for _, f := range files {
		path := f.Name()
		pathModifiedAt := f.ModTime()
		if pathModifiedAt.After(latestTime) {
			latestTime = pathModifiedAt
			latest = path
		}
	}
	return
}
