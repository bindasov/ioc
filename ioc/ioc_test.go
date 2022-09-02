package ioc

import (
	"testing"

	"github.com/bindasov/ioc/commands"
	"github.com/stretchr/testify/require"
)

func TestIoC_Resolve(t *testing.T) {
	type deps struct {
		ioc *IoC
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "success",
			handler: func(t *testing.T, deps *deps) {
				result := deps.ioc.Resolve("test")
				require.Equal(t, commands.NewTestCommand1(), result)
			},
		},
		{
			name: "not registered command",
			handler: func(t *testing.T, deps *deps) {
				result := deps.ioc.Resolve("not registered command")
				require.Equal(t, nil, result)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ioc := NewIoC()

			ioc.Resolve("IoC.Register", "test", func(args ...interface{}) commands.Command {
				return commands.NewTestCommand1()
			}).Execute()

			deps := &deps{
				ioc: ioc,
			}

			tc.handler(t, deps)
		})
	}
}

func TestIoC_ResolveWithScopes(t *testing.T) {
	type deps struct {
		ioc *IoC
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "success switching between scopes",
			handler: func(t *testing.T, deps *deps) {
				result := deps.ioc.Resolve("test")
				require.Equal(t, commands.NewTestCommand1(), result)

				deps.ioc.Resolve("Scopes.Current", "1").Execute()

				result = deps.ioc.Resolve("test")
				require.Equal(t, commands.NewTestCommand2(), result)

				deps.ioc.Resolve("Scopes.Current", "0").Execute()

				result = deps.ioc.Resolve("test")
				require.Equal(t, commands.NewTestCommand1(), result)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ioc := NewIoC()

			ioc.Resolve("IoC.Register", "test", func(args ...interface{}) commands.Command {
				return commands.NewTestCommand1()
			}).Execute()
			ioc.Resolve("Scopes.New", "1").Execute()
			ioc.Resolve("IoC.Register", "test", func(args ...interface{}) commands.Command {
				return commands.NewTestCommand2()
			}).Execute()
			ioc.Resolve("Scopes.Current", "0").Execute()

			deps := &deps{
				ioc: ioc,
			}

			tc.handler(t, deps)
		})
	}
}
