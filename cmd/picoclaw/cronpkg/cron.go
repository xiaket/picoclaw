// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package cronpkg

var cronStorePath string

func SetCronStorePath(workspace string) {
	cronStorePath = workspace + "/cron/jobs.json"
}

func GetCronStorePath() string {
	return cronStorePath
}
