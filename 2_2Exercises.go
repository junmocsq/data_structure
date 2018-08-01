package main

type t1 struct {
	name   string
	length int
	num    int
}

const (
	_       = iota
	Single
	Married
)

type t2 struct {
	t1
	marriage int
}

func main() {

}
