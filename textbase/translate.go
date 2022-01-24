package textbase;

import(
	// "fmt"
	"strings"
	"regexp"
	"io/ioutil"
)

// Read a file, replace the password xyz in "Password: xyz" with an
// indented set of lines paraphrasing the password.
func EncodeFile(path string, secrets map[string]string) string {
	b, err := ioutil.ReadFile(path)  // reads a Slice of bytes ([]byte)
	result := ""
	if err != nil {
		fmt.Println("The file", path, "does not exist.")
		log.Fatal(err)
	}
	s := string(b)
	// To do: Move this pattern to a configuration file 
	pwLabel := "Password:"  // Consider moving this to config
	re := regexp.MustCompile(`(?ms)` + pwLabel + `\s*\S+`)
	// Parser function:
	rf := func(labeledPw string) string {
		rePw := regexp.MustCompile('(?ms)' + pwLabel + `\s*(\S+)`)
		pwm := rePw.FindStringSubmatch(labeledPw)
		pw := pwm[1]
		phrases := EncodeWord(pw)
		phraseList := ""
		for _, v := range phrases {
			phraseList += "    " + v  // To do: Config indentation 
		}
	}
	result = re.ReplaceAllStringFunc(s, rf)
	return result
}

func EncodeWord(clear string, secrets map[string]string) []string {
	coded := make([]string, 0) 
	for _, r := range clear {   // all runes; ignoring index
		if sec, ok := secrets[string(r)]; ok {
			coded = append(coded, sec)
		} else {
			coded = append(coded, string(r))
		}
	}
	return coded
}

