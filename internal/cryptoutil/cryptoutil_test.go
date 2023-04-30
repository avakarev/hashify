package cryptoutil_test

import (
	"testing"

	"github.com/avakarev/go-testutil"

	"github.com/avakarev/hashify/internal/cryptoutil"
)

func file() string {
	return testutil.FixturePath("../../test/file")
}

func TestChecksumCRC32(t *testing.T) {
	sum, err := cryptoutil.Checksum(file(), cryptoutil.CRC32)
	testutil.MustNoErr(err, t)
	testutil.Diff("7d14dddd", sum, t)
}

func TestChecksumMD5(t *testing.T) {
	sum, err := cryptoutil.Checksum(file(), cryptoutil.MD5)
	testutil.MustNoErr(err, t)
	testutil.Diff("8ddd8be4b179a529afa5f2ffae4b9858", sum, t)
}

func TestChecksumSHA1(t *testing.T) {
	sum, err := cryptoutil.Checksum(file(), cryptoutil.SHA1)
	testutil.MustNoErr(err, t)
	testutil.Diff("a0b65939670bc2c010f4d5d6a0b3e4e4590fb92b", sum, t)
}

func TestChecksumSHA224(t *testing.T) {
	sum, err := cryptoutil.Checksum(file(), cryptoutil.SHA224)
	testutil.MustNoErr(err, t)
	testutil.Diff("de9d76f0f6a015ab6629138a42835e7b44571995e4abb291c0817261", sum, t)
}

func TestChecksumSHA256(t *testing.T) {
	sum, err := cryptoutil.Checksum(file(), cryptoutil.SHA256)
	testutil.MustNoErr(err, t)
	testutil.Diff("03ba204e50d126e4674c005e04d82e84c21366780af1f43bd54a37816b6ab340", sum, t)
}

func TestChecksumSHA384(t *testing.T) {
	sum, err := cryptoutil.Checksum(file(), cryptoutil.SHA384)
	testutil.MustNoErr(err, t)
	testutil.Diff("07f60df0b95043b3a3717638e7776ab76ebaa4fc705ba659063229cf162980c04a9f7496dcda50de6510d40fde3eba8a", sum, t)
}

func TestChecksumSHA512(t *testing.T) {
	sum, err := cryptoutil.Checksum(file(), cryptoutil.SHA512)
	testutil.MustNoErr(err, t)
	testutil.Diff("830445e86a0cfafac4e1531002356f384847a11a7456fb8ccb81ab36e37bff28f34fa2c5bfdd347e964c5c5df0fc305de6394368219307b2ceeb0ec84b7c2b31", sum, t)
}
