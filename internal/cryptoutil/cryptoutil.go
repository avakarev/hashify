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
	"hash/crc32"
	"io"
	"log"
	"os"
)

const (
	CRC32  = "crc32"  // CRC32 keeps the name of crc32 algorithm
	MD5    = "md5"    // MD5 keeps the name of md5 algorithm
	SHA1   = "sha1"   // SHA1 keeps the name of sha1 algorithm
	SHA224 = "sha224" // SHA224 keeps the name of sha224 algorithm
	SHA256 = "sha256" // SHA256 keeps the name of sha256 algorithm
	SHA384 = "sha384" // SHA384 keeps the name of sha384 algorithm
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
	case CRC32:
		return crc32.NewIEEE(), nil
	case MD5:
		return md5.New(), nil //#nosec G401 -- md5 used on purpose
	case SHA1:
		return sha1.New(), nil //#nosec G401 -- sha1 used on purpose
	case SHA224:
		return sha256.New224(), nil
	case SHA256:
		return sha256.New(), nil
	case SHA384:
		return sha512.New384(), nil
	case SHA512:
		return sha512.New(), nil
	default:
		return nil, fmt.Errorf("%q hash algorithm is not supported", name)
	}
}
