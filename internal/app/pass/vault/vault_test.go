package vault

import (
	"bytes"
	"crypto/aes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

var (
	filepath      = "./test-vault.vlt"
	masterKey     = []byte("0123456789ABCDEF")
	expectedVault *Vault
)

func init() {
	expectedVault = NewVault(filepath, masterKey)
	iv := bytes.Repeat([]byte{0}, aes.BlockSize)
	expectedVault.encryptionIvReader = bytes.NewBuffer(iv)

	expectedVault.AddEntry(Entry{BinaryPath: "test1", Secrets: map[string]string{"foo1": "bar1", "foo2": "bar2"}})
	expectedVault.AddEntry(Entry{BinaryPath: "test2", Secrets: map[string]string{"foo1": "bar1", "foo2": "bar2"}})
}

func TestAddEntry(t *testing.T) {
	v := NewVault(filepath, masterKey)
	expectedEntry := Entry{BinaryPath: "test1", Secrets: map[string]string{"foo1": "bar1", "foo2": "bar2"}}

	assert.Equal(t, 0, len(v.data.Entries))

	v.AddEntry(expectedEntry)
	assert.Equal(t, expectedEntry, v.data.Entries[0])
}

func TestEditEntry(t *testing.T) {
	v := NewVault(filepath, masterKey)
	e := Entry{BinaryPath: "test1", Secrets: map[string]string{"foo1": "bar1", "foo2": "bar2"}}

	v.AddEntry(e)

	e.BinaryPath = "test2"
	v.EditEntry(0, e)

	assert.Equal(t, e, v.data.Entries[0])
}

func TestDeleteEntry(t *testing.T) {
	v := NewVault(filepath, masterKey)
	e := Entry{BinaryPath: "test1", Secrets: map[string]string{"foo1": "bar1", "foo2": "bar2"}}

	v.AddEntry(e)
	v.DeleteEntry(0)

	assert.Equal(t, 0, len(v.data.Entries))
}

func TestSave(t *testing.T) {
	expectedVault.Save()

	assert.FileExists(t, filepath)

	expectedVaultData, _ := ioutil.ReadFile("testdata/test-vault.vlt.golden")
	actualVaultData, _ := ioutil.ReadFile(filepath)
	assert.Equal(t, expectedVaultData, actualVaultData)

	_ = os.Remove(filepath)
}

func TestOpen(t *testing.T) {
	v := NewVault("testdata/test-vault.vlt.golden", masterKey)

	_ = v.Open()

	assert.Equal(t, expectedVault.data, v.data)
}
