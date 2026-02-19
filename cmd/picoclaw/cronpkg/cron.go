// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package cronpkg

import "path/filepath"

var cronStorePath string

func SetCronStorePath(workspace string) {
	cronStorePath = filepath.Join(workspace, "cron", "jobs.json")
}

func GetCronStorePath() string {
	return cronStorePath
}
