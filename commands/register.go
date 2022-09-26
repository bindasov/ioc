package commands

type register struct {
	container map[string]func(args ...interface{}) interface{}
	key       string
	execute   func(args ...interface{}) interface{}
}

func (r *register) Execute() {
	r.container[r.key] = r.execute
}

func NewRegisterCommand(args ...interface{}) interface{} {
	register := &register{
		container: args[0].(map[string]func(args ...interface{}) interface{}),
		key:       args[1].(string),
		execute:   args[2].(func(args ...interface{}) interface{}),
	}
	return register
}
