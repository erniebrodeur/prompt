package main // import "github.com/erniebrodeur/prompt"

import (
	"fmt"

	"github.com/erniebrodeur/prompt/internal/segments"
	"golang.org/x/sys/unix"
)

var (
	git   = segments.Git{}
	host  = segments.Host{}
	left  = segments.Bookend{Left: true}
	login = segments.Login{}
	mid   = segments.Mid{}
	pwd   = segments.Pwd{}
	right = segments.Bookend{}
	shell = segments.ShellLevel{}
)

func main() {
	var terminalWidth = 40

	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)

	if err == nil {
		terminalWidth = int(ws.Col)
	}

	mid.Count = terminalWidth -
		(left.Len() * 2) - (right.Len() * 2) -
		git.Len() - pwd.Len() - host.Len() -
		login.Len() - shell.Len() - 3 // space count - 1

	output()
}

func output() {
	l := left.Output()
	r := right.Output()

	fmt.Printf("%v %v%v %v%v%v %v%v %v", l, pwd.Output(), git.Output(), r, mid.Output(), l, login.Output(), host.Output(), r)
}

func status() {
	fmt.Printf("bookend left:  %v\n", left.Output())
	fmt.Printf("bookend right: %v\n", right.Output())
	fmt.Printf("git:           %v\n", git.Output())
	fmt.Printf("host:          %v\n", host.Output())
	fmt.Printf("login:         %v\n", login.Output())
	fmt.Printf("mid:           %v\n", mid.Output())
	fmt.Printf("pwd:           %v\n", pwd.Output())
	fmt.Printf("shell:         %v\n", shell.Output())
}
