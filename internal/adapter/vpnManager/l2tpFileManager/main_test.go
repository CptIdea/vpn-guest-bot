package l2tpFileManager

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManager_Create(t *testing.T) {
	// create a temporary file for testing
	file, err := ioutil.TempFile("", "test.*.conf")
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	// create a manager with the temporary file and a mock password generator
	manager := New(file.Name(), mockPasswordGenerator{})

	// create a new L2TP VPN connection
	password, err := manager.Create("Test Connection")
	assert.NoError(t, err)

	// verify that the file was written correctly
	data, err := ioutil.ReadFile(file.Name())
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"Test Connection" l2tpd "`+password+`" *`)
}

func TestManager_Delete(t *testing.T) {
	// create a temporary file with some test data
	file, err := ioutil.TempFile("", "test.*.conf")
	assert.NoError(t, err)
	defer os.Remove(file.Name())
	_, err = file.WriteString(
		`"Connection 1" l2tpd "password1" *
"Connection 2" l2tpd "password2" *
"Connection 3" l2tpd "password3" *
`,
	)
	assert.NoError(t, err)

	// create a manager with the temporary file and a mock password generator
	manager := New(file.Name(), mockPasswordGenerator{})

	// delete a connection
	err = manager.Delete("Connection 2")
	assert.NoError(t, err)

	// verify that the file was updated correctly
	data, err := ioutil.ReadFile(file.Name())
	assert.NoError(t, err)
	assert.Contains(
		t, string(data), `"Connection 1" l2tpd "password1" *
"Connection 3" l2tpd "password3" *
`,
	)
}

// mockPasswordGenerator is a mock
type mockPasswordGenerator struct {
}

func (m mockPasswordGenerator) Generate() string {
	return "some password"
}
