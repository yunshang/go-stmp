package stmp
import (
	"log"
	"bufio"
    "fmt"
    "os"
	"regexp"
	"go-stmp/pkg/utils"
	"go-stmp/pkg/config"
)

// ReadFile set config
func ReadFile(name string, config *config.Config) {
	name = regexp.QuoteMeta(name)
	ver := regexp.MustCompile(fmt.Sprintf(`^\d.*\$s.sql$`, name))
	files, err := utils.FindFile(ver)
	if len(files) == 0 {
		fmt.Errorf("can't find file: %s*.mail", ver)
	}
    file, err := os.Open(files[0])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	matches := map[string]string{}

	scanner := bufio.NewScanner(file)
    scan := ""
    for scanner.Scan() {
		scan = regexp.QuoteMeta(scanner.Text())
		scan = regexp.MustCompile(fmt.Sprintf(`^\d.*\$s.sql$`, scan))
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

// SendMail  send mail
func SendMail(name string)  {
	config := config.New()
	ReadFile(name, &config)
}
