package ioc

import (
	"github.com/bindasov/ioc/commands"
)

type IoC struct {
	scopes        map[string]map[string]func(args ...interface{}) interface{}
	containerName string
}

func (i *IoC) Resolve(key string, args ...interface{}) interface{} {
	if _, ok := i.scopes[i.containerName][key]; ok {
		return i.scopes[i.containerName][key](args...)
	}
	return nil
}

func (i *IoC) register(args ...interface{}) interface{} {
	return commands.NewRegisterCommand(i.scopes[i.containerName], args[0], args[1])
}

func (i *IoC) createNewScope(args ...interface{}) interface{} {
	container := make(map[string]func(args ...interface{}) interface{})
	container["IoC.Register"] = i.register
	container["Scopes.New"] = i.createNewScope
	container["Scopes.Current"] = i.setCurrentScope
	return commands.NewScopeCommand(i.scopes, args[0].(string), container, &i.containerName)
}

func (i *IoC) setCurrentScope(args ...interface{}) interface{} {
	return commands.NewCurrentScopeCommand(&i.containerName, args[0])
}

func NewIoC() *IoC {
	ioc := &IoC{
		scopes: make(map[string]map[string]func(args ...interface{}) interface{}),
	}
	ioc.createNewScope("0").(commands.Command).Execute()
	return ioc
}
