package textbase;

import(
	// "fmt"
	"log"
	"os"
	"strings"
	"bufio"
)



// Read the config file at file path path and return a map with
// phrases indexed by one-character strings that they paraphrase.
func ReadConfig(path string) map[string]string {
	var secrets = make(map[string]string)
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines) // Setting the Split function
	for scanner.Scan() { // lines; stops by becoming false at EOF
		parts := strings.SplitN(scanner.Text(), ":", 2)
		if len(parts) == 2 {
			clr := strings.TrimSpace(parts[0])
			scr := strings.TrimSpace(parts[1])
			secrets[clr] = scr
		}
	}
	return secrets
}	

