package space

var (
	newFlags = BaseNewFlags()

	SubCmdNew = &cli.Command{
		Name:        "new",
		Usage:       "Create a new space within the program",
		Description: "Create a new space configuration within the infrastructure program",
		Flags:       newFlags,
		Action:      newSpace,
	}
)

func newSpace(cCtx *cli.Context) error {
	space := &base.Space{
		Name:      flagSpace,
		Region:    flagRegion,
		AccountId: flagAccountId,
		Operator:  flagOperator,
	}
	root := base.RootAccount()
	err := createNewSpace(root, space)
	if err != nil {
		return err
	}
	return nil
}

func createNewSpace(root *base.Root, space *base.Space) error {
	return nil
}
