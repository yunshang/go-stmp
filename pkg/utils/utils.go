package utils
import (
	"time"
	"fmt"
	"os"
	"io"
	"path/filepath"
	"io/ioutil"
	"regexp"
)

const MailDir = "./mails"
const MailTemplate = "Sender:\n\nTo:\n\nSubject:\n\nBody:\n\n"

// NewMailTemplate create a new mail file
func NewMailTemplate(name string) error {
	timestamp := time.Now().UTC().Format("20190417150405")
	if name == "" {
		return fmt.Errorf("please specify a name for the new tmplate")
	}

	name = fmt.Sprintf("%s_%s.mail", timestamp, name)

	// ensure email directory exists
	if err := ensureDir(filepath.Dir(MailDir)); err != nil {
		return err
	}

	// check file does not already exist
	path := filepath.Join(MailDir, name)
	fmt.Printf("Creating Mail file: %s\n", path)

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return fmt.Errorf("file already exists")
	}

	// write new mail template
	file, err := os.Create(path)
	if err != nil {
	    fmt.Print(err)
		return err
	}

	defer mustClose(file)
	_, err = file.WriteString(MailTemplate)
	return err
}

// ensureDir creates a directory if it does not already exist
func ensureDir(dir string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("unable to create directory `%s`", dir)
	}

	return nil
}

// mustClose ensures a stream is closed
func mustClose(c io.Closer) {
	if err := c.Close(); err != nil {
		panic(err)
	}
}

// FindFile find mail file
func FindFile(re *regexp.Regexp) ([]string, error) {
	files, err := ioutil.ReadDir(MailDir)
	if err != nil {
		return nil, fmt.Errorf("could not find mail directory `%s`", MailDir)
	}

	matches := []string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if !re.MatchString(name) {
			matches = append(matches, fmt.Sprintf("%s/%s", MailDir,name))
			break
		}

	}

	return matches, nil
}