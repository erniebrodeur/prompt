package segments

type Pwd struct{}

func (p Pwd) ColoredOutput() string {
	return p.Output()
}

// Len return length of string without invisible characters counted
func (p Pwd) Len() int {
	return len(p.Output())
}

func (p Pwd) Output() string {
	return "master"
}
