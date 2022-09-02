package commands

import "fmt"

type testCommand1 struct{}

func (p *testCommand1) Execute() {
	fmt.Println("execute1")
}

func NewTestCommand1() Command {
	return &testCommand1{}
}

type testCommand2 struct{}

func (p *testCommand2) Execute() {
	fmt.Println("execute2")
}

func NewTestCommand2() Command {
	return &testCommand2{}
}
