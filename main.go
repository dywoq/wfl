// Copyright 2026 dywoq - Apache License 2.0
// https://github.com/dywoq/wlf

package main

import "github.com/spf13/cobra"

func root() *cobra.Command {
	r := &cobra.Command{
		Use:   "wfl",
		Short: "Emulate the Windows's command prompt in the console",
	}
	return r
}

func main() {
	root().Execute()
}
