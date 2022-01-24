package main;

import(
	"log"
	"fmt"
	"os"
	"path/filepath"
	"github.com/hermannfass/mempass/textbase"
)

func main() {
	fmt.Println("Here is main.")
	cfn := "mempass.config"
	// To do: Search for the config file or ask for the path to it.
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	confpath := filepath.Join(home, cfn)
	fmt.Println("Reading configuration:", confpath)
	secrets := textbase.ReadConfig(confpath)

	t := textbase.Encode(secrets, "boh")
	fmt.Printf("Umschreibung ist:\n%v\n", t)
}

