package awx

import "github.com/urfave/cli"

var flags = []cli.Flag{
	Host,
	UserName,
	Password,
	TemplateName,
	SSL,
	ExtraVars,
}

var Host = cli.StringFlag{
	Name:  "host",
	Value: "127.0.0.1",
	Usage: "IP or HostName for AWX",
}

var UserName = cli.StringFlag{
	Name:  "username, u",
	Value: "admin",
	Usage: "Login UserName",
}

var Password = cli.StringFlag{
	Name:  "password, p",
	Usage: "Login Password",
}

var TemplateName = cli.StringFlag{
	Name:  "template, t",
	Usage: "Specify the template name to execute",
}

var SSL = cli.BoolFlag{
	Name:  "ssl",
	Usage: "Specified only when using SSL",
}

var ExtraVars = cli.StringFlag{
	Name:  "extra, e",
	Usage: "Specify extra vars",
}
