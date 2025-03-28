package run

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/sunshinexcode/gotk/vfile"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vshell"
	"github.com/sunshinexcode/gotk/vstr"
)

const (
	CliHttp      = "http"
	CliTcp       = "tcp"
	CliWebsocket = "websocket"
)

func initProject(args []string, projectType string) {
	// Current dir
	wd, err := os.Getwd()
	if err != nil {
		vlog.Error("current directory", "err", err)
		return
	}

	if len(args) > 0 {
		if args[0] != "." {
			wd = fmt.Sprintf("%s/%s", wd, args[0])
		}
	}

	vlog.Info("project directory", "dir", wd)

	// Check if path exists
	_, err = os.Stat(wd)
	// Path not existed
	if err != nil && os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(wd, 0754); err != nil {
			vlog.Error("create dir", "err", err)
			return
		}
	} else {
		// Check path not empty
		d, _ := os.ReadDir(wd)
		if len(d) > 0 {
			vlog.Infof("directory existed and not empty, %s", wd)
			return
		}
	}

	// Download example source
	vlog.Infof("download source")
	if res, err := vshell.Exec("cd /tmp && git clone https://github.com/sunshinexcode/gotk.git"); err != nil {
		if vstr.Trim(res) == "fatal: destination path 'gotk' already exists and is not an empty directory." {
			vlog.Infof("download source, git pull")
			_, _ = vshell.Exec("cd /tmp/gotk && git checkout main && git pull")
		} else {
			vlog.Error("download source", "err", err, "res", res)
			return
		}
	}

	// Copy dir
	vlog.Infof("copy dir")
	dirCode := vstr.S("%s/%s", wd, projectType)
	if err = vfile.Copy(vstr.S("/tmp/gotk/examples/%s", projectType), dirCode); err != nil {
		vlog.Error("copy dir", "err", err)
		return
	}

	// Remove .git
	vlog.Infof("remove .git")
	if res, err := vshell.Exec(vstr.S("rm -rf %s/.git", wd)); err != nil {
		vlog.Error("remove .git", "err", err, "res", res)
		return
	}

	// Replace project name
	vlog.Infof("replace project name")
	projectName := filepath.Base(wd)
	projectNameSearch := vstr.S("gotk-example-%s", projectType)
	replaceCode(projectType, projectName, projectNameSearch, dirCode)

	// Finished
	vlog.Infof("finished")
}

func replaceCode(projectType string, projectName string, projectNameSearch string, currentDir string) {
	_ = vfile.ReplaceFile(projectNameSearch, projectName, vstr.S("%s/build/Dockerfile", currentDir))
	_ = vfile.ReplaceFile(`name = "gotk"`, vstr.S(`name = "%s"`, projectName), vstr.S("%s/configs/config-default.toml", currentDir))

	_ = vfile.ReplaceFile(projectNameSearch, projectName, vstr.S("%s/Makefile", currentDir))
	_ = vfile.ReplaceFile(vstr.S("Gotk example %s demo", projectType), cases.Title(language.Und).String(strings.Replace(projectName, "-", " ", -1)), vstr.S("%s/README.md", currentDir))
}
