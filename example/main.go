package main

import "github.com/dixonwille/test/env"

func main() {
	env.UI.Log("Log Statement")
	env.UI.Output("Output Statement")
	env.UI.Running("Running Statement")
	env.UI.Success("Success Statement")
	env.UI.Info("Info Statement")
	env.UI.Warn("Warning Statement")
	env.UI.Error("Error statement")
}
