package commands

type register struct {
	container map[string]func(args ...interface{}) Command
	key       string
	execute   func(args ...interface{}) Command
}

func (r *register) Execute() {
	r.container[r.key] = r.execute
}

func NewRegisterCommand(args ...interface{}) Command {
	register := &register{
		container: args[0].(map[string]func(args ...interface{}) Command),
		key:       args[1].(string),
		execute:   args[2].(func(args ...interface{}) Command),
	}
	return register
}
