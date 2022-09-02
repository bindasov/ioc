package commands

type createNewScope struct {
	scopes       map[string]map[string]func(args ...interface{}) Command
	id           string
	newContainer map[string]func(args ...interface{}) Command
	container    *string
}

func (s *createNewScope) Execute() {
	s.scopes[s.id] = s.newContainer
	*s.container = s.id
}

func NewScopeCommand(args ...interface{}) Command {
	newScope := &createNewScope{
		scopes:       args[0].(map[string]map[string]func(args ...interface{}) Command),
		id:           args[1].(string),
		newContainer: args[2].(map[string]func(args ...interface{}) Command),
		container:    args[3].(*string),
	}
	return newScope
}

type setCurrentScope struct {
	currentScope *string
	scopeName    string
}

func (cs *setCurrentScope) Execute() {
	*cs.currentScope = cs.scopeName
}

func NewCurrentScopeCommand(args ...interface{}) Command {
	currentScope := &setCurrentScope{
		currentScope: args[0].(*string),
		scopeName:    args[1].(string),
	}
	return currentScope
}
