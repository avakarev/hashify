// Package cli parses command line flags and args
package cli

import (
	"flag"
	"log"
	"os"

	"github.com/avakarev/go-buildmeta"

	"github.com/avakarev/hashify/internal/cryptoutil"
)

var (
	helpFlag    bool
	versionFlag bool

	CRC32Flag  bool // CRC32Flag represents state of crc32 flag
	MD5Flag    bool // MD5Flag represents state of md5 flag
	SHA1Flag   bool // SHA1Flag represents state of sha1 flag
	SHA224Flag bool // SHA224Flag represents state of sha224 flag
	SHA256Flag bool // SHA256Flag represents state of sha256 flag
	SHA384Flag bool // SHA384Flag represents state of sha384 flag
	SHA512Flag bool // SHA512Flag represents state of sha512 flag
)

// HashFlags return list of hash algorithms passed as flags
func HashFlags() []string {
	flags := make([]string, 0)
	if CRC32Flag {
		flags = append(flags, cryptoutil.CRC32)
	}
	if MD5Flag {
		flags = append(flags, cryptoutil.MD5)
	}
	if SHA1Flag {
		flags = append(flags, cryptoutil.SHA1)
	}
	if SHA224Flag {
		flags = append(flags, cryptoutil.SHA224)
	}
	if SHA256Flag {
		flags = append(flags, cryptoutil.SHA256)
	}
	if SHA384Flag {
		flags = append(flags, cryptoutil.SHA384)
	}
	if SHA512Flag {
		flags = append(flags, cryptoutil.SHA512)
	}
	return flags
}

// Files returns list of files passed as arguments
func Files() []string {
	return flag.Args()
}

// Validate validates whether passed cli flags and arguments are enough
func Validate() {
	args := os.Args[1:]
	if len(args) == 1 {
		if helpFlag || args[0] == "help" {
			printHelp()
			os.Exit(0)
		}
		if versionFlag || args[0] == "version" {
			printVersion()
			os.Exit(0)
		}
		printHelp()
		os.Exit(1)
	}

	if len(HashFlags()) == 0 || len(Files()) == 0 {
		printHelp()
		os.Exit(1)
	}
}

func printVersion() {
	log.Println("CLI utility to calculate file checksums")
	log.Println("")
	log.Printf("  %-10s %s\n", "Version", buildmeta.Ref)
	log.Printf("  %-10s %s\n", "BuildDate", buildmeta.BuildTimeUTC)
	log.Printf("  %-10s %s\n", "Compiler", buildmeta.Compiler())
	log.Printf("  %-10s %s/%s\n", "Platform", buildmeta.OS(), buildmeta.Arch())
}

func printHelp() {
	log.Println("CLI utility to calculate file checksums")
	log.Println("\nUsage:")
	log.Println("")
	log.Println("  hashify [flags] [files]")
	log.Println("")
	log.Println("Flags:")
	flag.PrintDefaults()
}

func init() {
	flag.BoolVar(&helpFlag, "help", false, "print help")
	flag.BoolVar(&versionFlag, "version", false, "print version")
	flag.BoolVar(&CRC32Flag, cryptoutil.CRC32, false, "calculate CRC32")
	flag.BoolVar(&MD5Flag, cryptoutil.MD5, false, "calculate MD5")
	flag.BoolVar(&SHA1Flag, cryptoutil.SHA1, false, "calculate SHA1")
	flag.BoolVar(&SHA224Flag, cryptoutil.SHA224, false, "calculate SHA224")
	flag.BoolVar(&SHA256Flag, cryptoutil.SHA256, false, "calculate SHA256")
	flag.BoolVar(&SHA384Flag, cryptoutil.SHA384, false, "calculate SHA384")
	flag.BoolVar(&SHA512Flag, cryptoutil.SHA512, false, "calculate SHA512")
	flag.Parse()
}
