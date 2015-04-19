package book

import (
	//"fmt"
	"os"
	"os/exec"
	"strings"
)

type CurrentProject struct {
	Project
}

type Project struct {
	Path            string
	DotRulebookFile DotRulebookFile
}

type DotRulebookFile struct {
	path string
}

func NewCurrentProject() (c CurrentProject) {
	c.Path = c.GitDir()
	if c.DotRulebookExists() {
		c.DotRulebookFile = DotRulebookFile{c.DotRulebookPath()}
	}
	return
}

func (p CurrentProject) GitDir() string {
	c, _ := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	return strings.TrimSpace(string(c))
}

func (p Project) DotRulebookPath() string {
	return p.Path + "/.rulebook"
}

func (p Project) DotRulebookExists() bool {
	f := p.DotRulebookPath()
	_, err := os.Stat(f)
	return (err == nil)
}
