package Auth

//
//func GetLoginCmd() *cli.Command {
//	c := &cli.Command{
//		Name:        "login",
//		Action:      loginAction,
//		Description: "Authorize to access the OneSky API with access token",
//		Usage:       "Authorize to access the OneSky API with access token",
//		UsageText:   "`onesky auth login --access-token=ACCESS_TOKEN`",
//		Flags: []cli.Flag{
//			&cli.StringFlag{
//				Name:     "access-token",
//				Aliases:  []string{"t"},
//				Usage:    "Set `ACCESS_TOKEN`",
//				Required: true,
//			},
//		},
//	}
//
//	return c
//}
//
//func loginAction(c *cli.Context) (e error) {
//
//	fmt.Println(c.String("access-token"))
//	return e
//}
