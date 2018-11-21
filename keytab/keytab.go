package keytab

import (
	"fmt"
	"github.com/disaster37/go-keytab/cmd"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"regexp"
	"strings"
)

func (k *Keytab) IsExist() (bool, error) {

	if (k.Path == "") || (k.Principal == "") || (k.Cipher == "") {
		return false, errors.New("You must provide path, principal and cipher")
	}

	// Create regex to extract data from keytab
	r, err := regexp.Compile(`\s*\d+\s+(\S+)\s+\((\S+)\)`)
	if err != nil {
		panic(err)
	}

	// First, check if keytab file already exist
	if _, err := os.Stat(k.Path); !os.IsNotExist(err) {
		command := fmt.Sprintf("klist -ek %s", k.Path)
		result, err := cmd.Run(command)
		if err != nil {
			return false, err
		}
		if result.ExitCode != 0 {
			return false, errors.New(fmt.Sprintf("Command return with exit code %d and stdout: %s", result.ExitCode, result.Stdout))
		}

		//  Extract data
		lines := strings.Split(result.Stdout, "\n")
		for _, line := range lines {
			matchs := r.FindStringSubmatch(line)
			if matchs != nil {
				keytab := &Keytab{
					Principal: matchs[1],
					Cipher:    matchs[2],
				}

				log.Debugf("Extract principal %s", keytab)
				if (keytab.Principal == k.Principal) && (keytab.Cipher == k.Cipher) {
					log.Infof("The principal %s with cipher %s exist on keytab %s", k.Principal, k.Cipher, k.Path)
					return true, nil
				}
			}
		}

	} else {
		log.Debugf("The keytab file %s not found", k.Path)
	}

	return false, nil

}

func (k *Keytab) Create() (*Keytab, error) {

	if (k.Path == "") || (k.Principal == "") || (k.Cipher == "") || (k.Password == "") {
		return nil, errors.New("You must provide path, principal, password and cipher")
	}

	// Check if principal already exist on keytab file with same cipher
	isExist, err := k.IsExist()
	if err != nil {
		return nil, err
	}
	if isExist == true {
		return nil, errors.New(fmt.Sprintf("Principal %s with cipher %s already exist on %s", k.Principal, k.Cipher, k.Path))
	}

	command := fmt.Sprintf("printf \"%%b\" \"addent -password -p %s -k 1 -e %s\n%s\nwkt %s\" | ktutil", k.Principal, k.Cipher, k.Password, k.Path)
	result, err := cmd.Run(command)
	if err != nil {
		return nil, err
	}
	if result.ExitCode != 0 {
		return nil, errors.New(fmt.Sprintf("Command return with exit code %d and stdout: %s", result.ExitCode, result.Stdout))
	}
	log.Infof("The principal %s with cipher %s is successfully created on keytab %s", k.Principal, k.Cipher, k.Path)

	return k, nil
}

func (k *Keytab) Delete() error {

	if k.Path == "" {
		return errors.New("You must provide path")
	}

	if _, err := os.Stat(k.Path); !os.IsNotExist(err) {
		err = os.Remove(k.Path)
		if err != nil {
			return err
		}
		log.Infof("We remove the keytab %s successfully", k.Path)
	} else {
		return errors.New(fmt.Sprintf("Keytab %s not found", k.Path))
	}

	return nil

}
