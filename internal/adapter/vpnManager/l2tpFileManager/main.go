package l2tpFileManager

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"vpn-guest-bot/internal/core/interfaces"

	"github.com/pkg/errors"
)

type manager struct {
	path              string
	passwordGenerator interfaces.PasswordGenerator
}

func New(path string, passwordGenerator interfaces.PasswordGenerator) interfaces.VpnManager {
	return &manager{path: path, passwordGenerator: passwordGenerator}
}

func (m *manager) Create(name string) (string, error) {
	password := m.passwordGenerator.Generate()

	file, err := os.OpenFile(m.path, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return "", errors.Wrap(err, "open file")
	}
	defer file.Close()

	_, err = file.Write([]byte(fmt.Sprintf("%q l2tpd %q *\n", name, password)))
	if err != nil {
		return "", errors.Wrap(err, "write file")
	}

	return password, nil
}

func (m *manager) Delete(name string) error {
	file, err := os.OpenFile(m.path, os.O_RDWR, 0644)
	if err != nil {
		return errors.Wrap(err, "open file")
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.Wrap(err, "read file")
	}

	rows := bytes.Split(data, []byte("\n"))
	prefix := []byte(`"` + name + `"`)
	for _, row := range rows {
		if bytes.HasPrefix(row, prefix) {
			data = bytes.Replace(data, append(row, '\n'), []byte{}, 1)
			break
		}
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return errors.Wrap(err, "seek file")
	}

	_, err = file.Write(data)
	if err != nil {
		return errors.Wrap(err, "write file")
	}

	err = file.Truncate(int64(len(data)))
	if err != nil {
		return errors.Wrap(err, "clear file")
	}

	return nil
}
