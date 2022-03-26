package commands

type Command interface {
	Run(args []string)
}
