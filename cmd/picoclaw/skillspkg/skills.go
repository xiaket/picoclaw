// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package skillspkg

import (
	"path/filepath"

	"github.com/sipeed/picoclaw/pkg/skills"
)

var (
	workspace        string
	globalSkillsDir  string
	builtinSkillsDir string
)

func SetWorkspace(ws string) {
	workspace = ws
}

func SetGlobalDirs(global, builtin string) {
	globalSkillsDir = global
	builtinSkillsDir = builtin
}

func getInstaller() *skills.SkillInstaller {
	return skills.NewSkillInstaller(workspace)
}

func getLoader() *skills.SkillsLoader {
	return skills.NewSkillsLoader(workspace, globalSkillsDir, builtinSkillsDir)
}

func getWorkspace() string {
	return workspace
}

func getBuiltinSkillsDir() string {
	return filepath.Join(workspace, "../picoclaw/skills")
}
