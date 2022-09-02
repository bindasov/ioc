package main

import (
	"fmt"
	"github.com/bindasov/ioc/commands"
	container "github.com/bindasov/ioc/ioc"
)

func main() {
	ioc := container.NewIoC()

	ioc.Resolve("IoC.Register", "test", func(args ...interface{}) commands.Command {
		fmt.Println("test command1")
		return commands.NewTestCommand1()
	}).Execute()

	ioc.Resolve("test").Execute()

	ioc.Resolve("Scopes.New", "1").Execute()
	ioc.Resolve("IoC.Register", "test", func(args ...interface{}) commands.Command {
		fmt.Println("test command2")
		return commands.NewTestCommand2()
	}).Execute()

	ioc.Resolve("test").Execute()

	ioc.Resolve("Scopes.Current", "0").Execute()
	ioc.Resolve("test").Execute()

}
