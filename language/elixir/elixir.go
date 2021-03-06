package elixir

import (
	"github.com/maddyonline/glot-code-runner/cmd"
	"github.com/maddyonline/glot-code-runner/util"
	"path/filepath"
)

func Run(files []string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	sourceFiles := util.FilterByExtension(files, "ex")
	args := append([]string{"elixirc"}, sourceFiles...)
	return cmd.Run(workDir, args...)
}
