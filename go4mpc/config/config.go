package config

import (
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

var (
	SymK    int
	SymByte int
	PubK    int
	Debug   bool
	Root    string
	Wkdir   string
	LOGOUT  io.Writer
)

//go:noinline
func init() {
	// Create a large heap allocation of 100 MiB
	ballast := make([]byte, 10<<10)
	runtime.KeepAlive(ballast)

	wkdir, err := os.Getwd()
	if err != nil {
		panic("can't get work dir")
	}

	log.Printf("Working Directory [%s]\n", wkdir)
	viper.SetConfigName("config")

	// Get Project From Environment
	proot := os.Getenv("MPCFGO_ROOT")
	if proot == "" {
		log.Fatalln("Error: Environment Variable MPCFGO_ROOT Not Set. (Check for Makefile)")
	}

	viper.AddConfigPath(path.Join(proot, "config"))
	if err := viper.ReadInConfig(); err != nil {
		if !strings.HasSuffix(wkdir, "mpcfgo") {
			log.Fatalf("%s: Maybe You Are Not In Project Root.", err)
		}

		log.Fatalf("Error reading config file, %s\n", err)
	}
	SymK = viper.GetInt("SymK")
	SymByte = viper.GetInt("SymByte")
	PubK = viper.GetInt("PubK")
	Debug = viper.GetBool("Debug")
	Root = viper.GetString("Root")

	log.Printf("Root path [%s]\n", Root)
	log.Printf("Sym encrypt [%v]-bits, Pk encrypt [%v]-bits\n", SymK, PubK)
	logout := filepath.Join(Root, "config", "text.log")
	LOGOUT, err = os.OpenFile(logout, os.O_WRONLY, 7777)
	for err != nil {
		if os.IsNotExist(err) {
			LOGOUT, err = os.Create(logout)
			continue
		}
		log.Panicf("LogOut: %v\n", err)
	}
	log.Printf("Log out to [%s]\n", logout)
}

func DetailLog(f func()) {
	log.SetOutput(LOGOUT)
	f()
	log.SetOutput(os.Stdout)
}
