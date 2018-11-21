package keytab

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKeytab(t *testing.T) {

	logrus.SetLevel(logrus.DebugLevel)

	os.Remove("/tmp/test.keytab")

	// Add principal on new keytab file
	keytab := &Keytab{
		Path:      "/tmp/test.keytab",
		Principal: "test@TEST.TEST",
		Password:  "test",
		Cipher:    "aes128-cts-hmac-sha1-96",
	}
	keytab, err := keytab.Create()
	assert.NoError(t, err)
	assert.NotNil(t, keytab)
	isExist, err := keytab.IsExist()
	assert.NoError(t, err)
	assert.Equal(t, true, isExist)

	// Create existing principal exist with error
	keytab = &Keytab{
		Path:      "/tmp/test.keytab",
		Principal: "test@TEST.TEST",
		Password:  "test",
		Cipher:    "aes128-cts-hmac-sha1-96",
	}
	keytab, err = keytab.Create()
	assert.Error(t, err)

	// Create new principal on existing keytab
	keytab = &Keytab{
		Path:      "/tmp/test.keytab",
		Principal: "test2@TEST.TEST",
		Password:  "test",
		Cipher:    "aes128-cts-hmac-sha1-96",
	}
	keytab, err = keytab.Create()
	assert.NoError(t, err)
	assert.NotNil(t, keytab)
	isExist, err = keytab.IsExist()
	assert.NoError(t, err)
	assert.Equal(t, true, isExist)

	// Create same principal with new cipher
	keytab.Cipher = "aes256-cts-hmac-sha1-96"
	keytab, err = keytab.Create()
	assert.NoError(t, err)
	assert.NotNil(t, keytab)
	isExist, err = keytab.IsExist()
	assert.NoError(t, err)
	assert.Equal(t, true, isExist)

	// Delete existing keytab file
	err = keytab.Delete()
	assert.NoError(t, err)
	_, err = os.Stat(keytab.Path)
	assert.Equal(t, true, os.IsNotExist(err))

	// Delete not existing keytab file
	err = keytab.Delete()
	assert.Error(t, err)
}
