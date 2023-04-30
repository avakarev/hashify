// Package cryptoutil implements file checksum helpers
package cryptoutil

import (
	"crypto/md5"  //#nosec G501 -- md5 used on purpose
	"crypto/sha1" //#nosec G505 -- sha1 used on purpose
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

const (
	MD5    = "md5"    // MD5 keeps the name of md5 algorithm
	SHA1   = "sha1"   // SHA1 keeps the name of sha1 algorithm
	SHA256 = "sha256" // SHA256 keeps the name of sha256 algorithm
	SHA512 = "sha512" // SHA512 keeps the name of sha512 algorithm
)

// Checksum calculates checksum for given file by given hash algorithm
func Checksum(file string, hashName string) (string, error) {
	h, err := NewHash(hashName)
	if err != nil {
		return "", err
	}

	f, err := os.Open(file) //#nosec G304 -- user file provided by user on their machine
	if err != nil {
		return "", err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	copyBuf := make([]byte, 1024*1024)
	if _, err := io.CopyBuffer(h, f, copyBuf); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// NewHash returns hash algorithm implementation by name
func NewHash(name string) (hash.Hash, error) {
	switch name {
	case MD5:
		return md5.New(), nil //#nosec G401 -- md5 used on purpose
	case SHA1:
		return sha1.New(), nil //#nosec G401 -- sha1 used on purpose
	case SHA256:
		return sha256.New(), nil
	case SHA512:
		return sha512.New(), nil
	default:
		return nil, fmt.Errorf("%q hash algorithm is not supported", name)
	}
}
