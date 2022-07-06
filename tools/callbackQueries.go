package tbtools

import (
	"strings"
)

// Compiles callback data in command-args type.
// Concatenates command and args with \x00 symbol
func CompileCbQryData(command, args string) string {
	if args == "" {
		return command
	}

	return command + "\x00" + args
}

// Decompiles callback data in command-args type.
// Use to decompile output from CompileCbQryData.
func DecompileCbQryData(cbQryData string) (command, args string) {
	data := strings.SplitN(cbQryData, "\x00", 2)
	command = data[0]
	if len(data) == 2 {
		args = data[1]
	}
	return
}
