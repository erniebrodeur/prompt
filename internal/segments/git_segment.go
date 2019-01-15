package segments

import (
	"bufio"
	"regexp"
	"strings"
)

type Git struct {
	branch, remoteBranch, direction, dirty, gitString string
}

// type Git struct {
// 	branch       string
// 	remoteBranch string
// 	direction    string
// 	directionBy  int
// 	dirty        bool
// }

var gitHeaderRegexp = regexp.MustCompile(`## (?P<local_branch>\w*)...(?P<remote_branch>\S*)(..(?P<direction>ahead|behind) (?P<direction_count>\d)]|)`)

func (g Git) ColoredOutput() string {
	return g.Output()
}

// Len return length of string without invisible characters counted
func (g Git) Len() int {
	return len(g.Output())
}

func (g Git) Output() string {
	// g = Git.parse(g)
	// output := fmt.Sprintf(":%v", g.branch)

	return ""
}

func (g Git) dirtyOutput() string {
	// if g.dirty {
	// 	return "*"
	// }

	return ""
}

func (g Git) directionOutput() string {
	// if len(g.direction) != 0 {
	// 	if g.direction == "ahead" {
	// 		return " (push)"
	// 	}
	// 	return " (pull)"
	// }

	return ""
}
func (g Git) getGitString() {

}

func (g *Git) parse() Git {
	// out, err := exec.Command("git", "status", "--porcelain", "--ahead-behind", "-b").Output()

	// if err != nil {
	// 	return g
	// }

	if g.gitString == "" {
		// g.getGitOutput()
	}

	lines := strings.Split(string(g.gitString), "\n")

	gitHeaderRegexp.MatchString(lines[0])

	parts := gitHeaderRegexp.FindAllStringSubmatch(lines[0], -1)
	g.branch = parts[0][1]
	g.remoteBranch = parts[0][2]
	// g.direction = parts[0][4]
	// g.directionBy, _ = strconv.Atoi(parts[0][5])

	if len(lines) > 2 {
		g.dirty = "*"
	}

	scanner := bufio.NewScanner(strings.NewReader(string(g.gitString)))
	scanner.Split(bufio.ScanLines)

	return *g
}
