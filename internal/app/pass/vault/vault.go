package vault

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"encoding/gob"
	"github.com/QuentinBrosse/pass/internal/app/pass/utils"
	"io"
	"io/ioutil"
)

const (
	vaultFileMagic         = "PASS"
	vaultFileFormatVersion = 1
)

var DevEncryptionKey = []byte("0123456789ABCDEF") // TODO: remove this once we can access the masterkey properly

// Hold all information about an entry in the vault
type Entry struct {
	BinaryPath string            // Path of the program
	Secrets    map[string]string // Map of secret names/secrets for this Entry's program
}

// Hold all information about a vault which will be written to the vault file
type Data struct {
	Magic         string  // Magic ID identifying the vault file format
	FormatVersion int     // Format version of the vault file
	Entries       []Entry // Entries of the vault
}

// Hold all information about a vault needed at runtime
type Vault struct {
	filepath           string    // Full file path of the vault on the filesystem
	encryptionKey      []byte    // Encryption key used to encrypt the vault
	encryptionIvReader io.Reader // io.Reader used to generate an IV when encrypting the vault
	data               Data      // Vault actual data
}

// Save the Vault v to the file specified by v.filepath: compress it, encrypt it and write it to the file
func (v *Vault) Save() {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	gobEncoder := gob.NewEncoder(gzipWriter)

	if err := gobEncoder.Encode(v.data); err != nil {
		panic(err.Error())
	}

	if err := gzipWriter.Flush(); err != nil {
		panic(err.Error())
	}

	encryptedVaultData, err := aescbc.Encrypt(buf.Bytes(), v.encryptionKey, v.encryptionIvReader)
	if err != nil {
		panic(err.Error())
	}

	if err := ioutil.WriteFile(v.filepath, encryptedVaultData, 0644); err != nil {
		panic(err.Error())
	}
}

// Open the Vault v: read the file specified by v.filepath, decrypt it, uncompress it and load it into v.data
func (v *Vault) Open() (err error) {
	data, err := ioutil.ReadFile(v.filepath)

	if err != nil {
		return
	}

	data, err = aescbc.Decrypt(data, v.encryptionKey)

	if err != nil {
		return
	}

	gzipReader, err := gzip.NewReader(bytes.NewBuffer(data))

	if err != nil {
		return
	}

	gobDecoder := gob.NewDecoder(gzipReader)
	err = gobDecoder.Decode(&v.data)
	if err != nil {
		return
	}

	return nil
}

// Add the Entry entry to the Vault v
func (v *Vault) AddEntry(entry Entry) {
	v.data.Entries = append(v.data.Entries, entry)
}

// Replace the Entry at idx by Entry entry in the Vault v
func (v *Vault) EditEntry(idx int, entry Entry) {
	v.data.Entries[idx] = entry
}

// Delete Entry at idx from the Vault v
func (v *Vault) DeleteEntry(idx int) {
	v.data.Entries = append(v.data.Entries[:idx], v.data.Entries[idx+1:]...)
}

// Construct a new Vault with the specified filepath and encryptionKey
func NewVault(filepath string, encryptionKey []byte) *Vault {
	return &Vault{filepath: filepath, encryptionKey: encryptionKey, data: Data{Magic: vaultFileMagic, FormatVersion: vaultFileFormatVersion}, encryptionIvReader: rand.Reader}
}
