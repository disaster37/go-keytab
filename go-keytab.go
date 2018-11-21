package main

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/urfave/cli.v1"
	"os"
	"github.com/disaster37/go-keytab/cmd"
)

var debug bool
var keytabPath string


func main() {

	// Logger setting
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true
	formatter.ForceFormatting = true
	log.SetFormatter(formatter)
	log.SetOutput(os.Stdout)

	// CLI settings
	app := cli.NewApp()
	app.Usage = "Warpper for keytab management. You need to have ktutil and klist binary."
	app.Version = "develop"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path",
			Usage:       "The keytab full path",
			Destination: &keytabPath,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Display debug output",
			Destination: &debug,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "add-principal",
			Usage: "Add new principal on keytab",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "principal",
					Usage: "The principal to add on keytab",
				},
				cli.StringFlag{
					Name:  "password",
					Usage: "The password associated to the principal",
				},
				cli.StringFlag{
					Name:  "cipher",
					Usage: "The cipher to use when encrypt password",
					Value: "aes256-cts-hmac-sha1-96",
				},
			},
			Action: addPrincipal,
		},
		{
			Name:  "add-principal-ciphers",
			Usage: "Add new principal with multiple ciphers (comma separeted without space) on keytab",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "principal",
					Usage: "The principal to add on keytab",
				},
				cli.StringFlag{
					Name:  "password",
					Usage: "The password associated to the principal",
				},
				cli.StringFlag{
					Name:  "ciphers",
					Usage: "The list of cipher to use when encrypt password (separated by comma without space). It create one entry per cipher.",
					Value: "aes256-cts-hmac-sha1-96",
				},
			},
			Action: addPrincipalWithCipherList,
		},
		{
			Name:  "delete-keytab",
			Usage: "Delete keytab file",
			Action: delKeytab,
		},
		{
			Name:  "check-principal",
			Usage: "Check if principal with cipher exist on keytab. Return 0 if found",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "principal",
					Usage: "The principal to add on keytab",
				},
				cli.StringFlag{
					Name:  "cipher",
					Usage: "The cipher to use when encrypt password",
					Value: "aes256-cts-hmac-sha1-96",
				},
			},
			Action: checkPrincipal,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// Check the global parameter
func manageGlobalParameters() error {
	if debug == true {
		log.SetLevel(log.DebugLevel)
	}

	if keytabPath == "" {
		return errors.New("You must set --path parameter")
	}

	return nil
}

// Check the require binary is present
func checkRequire() error {
    result, err := cmd.Run("command -v ktutil")
    if err != nil {
        return err
    }
    if result.ExitCode != 0 {
        return errors.New("ktutil binaire not found. Maybee you need to install krb5-workstation (redhat) or krb5-user (debian) package.")
    }
    
    result, err = cmd.Run("command -v klist")
    if err != nil {
        return err
    }
    if result.ExitCode != 0 {
        return errors.New("ktutil binaire not found. Maybee you need to install krb5-workstation (redhat) or krb5-user (debian) package.")
    }
    
    return nil
}
