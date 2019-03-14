package main

import (
	"fmt"
	"github.com/disaster37/go-keytab/keytab"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
	"strings"
)

// addPrincipal permit to add new Principal on keytab.
// If keytab file not exist, it will be created.
// If keytab file already exist, the principal is append on it.
// If principal with the same cipher already exist, it return error.
func addPrincipal(c *cli.Context) error {

	err := manageGlobalParameters()
	if err != nil {
		return err
	}
	err = checkRequire()
	if err != nil {
		return err
	}
	if c.String("principal") == "" {
		return errors.New("You must set --principal parameter")
	}
	if c.String("password") == "" {
		return errors.New("You must set --password parameter")
	}
	if c.String("cipher") == "" {
		return errors.New("You must set --cipher parameter")
	}

	keytab := &keytab.Keytab{
		Path:      keytabPath,
		Principal: c.String("principal"),
		Password:  c.String("password"),
		Cipher:    c.String("cipher"),
	}

	_, err = keytab.Create()
	if err != nil {
		return err
	}

	return nil
}

// addPrincipalWithCipherList permit to add new Principal on keytab with multiple cipher.
// If keytab file not exist, it will be created.
// If keytab file already exist, the principal is append on it.
// If principal with the same cipher already exist, it return error.
func addPrincipalWithCipherList(c *cli.Context) error {

	err := manageGlobalParameters()
	if err != nil {
		return err
	}
	err = checkRequire()
	if err != nil {
		return err
	}
	if c.String("principal") == "" {
		return errors.New("You must set --principal parameter")
	}
	if c.String("password") == "" {
		return errors.New("You must set --password parameter")
	}
	if c.String("ciphers") == "" {
		return errors.New("You must set --ciphers parameter")
	}

	ciphers := strings.Split(c.String("ciphers"), ",")
	for _, cipher := range ciphers {
		keytab := &keytab.Keytab{
			Path:      keytabPath,
			Principal: c.String("principal"),
			Password:  c.String("password"),
			Cipher:    cipher,
		}

		_, err = keytab.Create()
		if err != nil {
			return err
		}
	}

	return nil

}

// delKeytab permit to remove keytab file.
// If keytab file not exist, it return error.
func delKeytab(c *cli.Context) error {

	err := manageGlobalParameters()
	if err != nil {
		return err
	}

	keytab := &keytab.Keytab{
		Path: keytabPath,
	}

	err = keytab.Delete()
	if err != nil {
		return err
	}

	return nil
}

// checkPrincipal permit to check if principal already exist on keytab.
// If keytab file not exist, it return 1.
// If not found, it return 1.
// If principal exist with same cipher on keytab, it return 0.
func checkPrincipal(c *cli.Context) error {

	err := manageGlobalParameters()
	if err != nil {
		return err
	}
	err = checkRequire()
	if err != nil {
		return err
	}
	if c.String("principal") == "" {
		return errors.New("You must set --principal parameter")
	}
	if c.String("cipher") == "" {
		return errors.New("You must set --cipher parameter")
	}

	keytab := &keytab.Keytab{
		Path:      keytabPath,
		Principal: c.String("principal"),
		Cipher:    c.String("cipher"),
	}

	isExist, err := keytab.IsExist()
	if err != nil {
		return err
	}

	if isExist != true {
		return cli.NewExitError(fmt.Sprintf("Principal %s with cipher %s not found in %s", keytab.Principal, keytab.Cipher, keytab.Path), 1)
	}

	return nil
}
