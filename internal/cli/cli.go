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

	MD5Flag    bool // MD5Flag represents state of md5 flag
	SHA1Flag   bool // SHA1Flag represents state of sha1 flag
	SHA256Flag bool // SHA256Flag represents state of sha256 flag
	SHA512Flag bool // SHA512Flag represents state of sha512 flag

	Files []string // Files represents list of files passed as arguments
)

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

	if !MD5Flag && !SHA1Flag && !SHA256Flag && !SHA512Flag {
		printHelp()
		os.Exit(1)
	}

	if len(Files) == 0 {
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
	flag.BoolVar(&MD5Flag, cryptoutil.MD5, false, "calculate MD5")
	flag.BoolVar(&SHA1Flag, cryptoutil.SHA1, false, "calculate SHA1")
	flag.BoolVar(&SHA256Flag, cryptoutil.SHA256, false, "calculate SHA256")
	flag.BoolVar(&SHA512Flag, cryptoutil.SHA512, false, "calculate SHA512")
	flag.Parse()
	Files = flag.Args()
}
