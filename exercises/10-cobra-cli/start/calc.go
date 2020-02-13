package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"strconv"
)

type calcAddCmd struct {
	a   int
	b   int
	out io.Writer
}

func newAddCmd(out io.Writer) *cobra.Command {
	calcAdd := &calcAddCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "add",
		Short: "adds two number, prints result and exits",
		RunE: func(cmd *cobra.Command, args []string) error {
			a, _ := strconv.Atoi(args[0])
			b, _ := strconv.Atoi(args[1])
			calcAdd.a = a
			calcAdd.b = b
			return calcAdd.run()
		},
	}
	return cmd
}

func (c *calcAddCmd) run() error {
	_, err := fmt.Fprintf(c.out, "Result: %d\n", c.a + c.b)
	if err != nil {
		return err
	}
	return nil
}
