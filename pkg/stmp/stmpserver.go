package stmp
import (
	"log"
	"bufio"
    "fmt"
    "os"
	"regexp"
	"go-stmp/pkg/utils"
	"go-stmp/pkg/config"
	// "crypto/tls"
	// "net"
	// "net/mail"
	"net/smtp"
)

// ReadFile set config
func ReadFile(name string) (map[string]string, error) {
	name = regexp.QuoteMeta(name)
	ver := regexp.MustCompile(fmt.Sprintf(`^\d.*\$s.sql$`, name))
    matches := make(map[string]string)
	files, err := utils.FindFile(ver)
	if len(files) == 0 {
		return matches,fmt.Errorf("can't find file: %s*.mail", ver)
	}
    file, err := os.Open(files[0])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
    scan := ""
	_match := regexp.MustCompile(`:`)
    for scanner.Scan() {
		scan = regexp.QuoteMeta(scanner.Text())
		if scan == "" {
			continue
		}
		s := _match.Split(scan, 2)
		matches[s[0]] = s[1]
	}

    if err := scanner.Err(); err != nil {
		return matches, err
	}

	return matches, nil
}

// SendMail  send mail
func SendMail(name string)  {
	config := config.New()
	mail_config, err := ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	config.Sender = mail_config["Sender"]
	config.To = mail_config["To"]
	config.Subject = mail_config["Subject"]
	config.Body = mail_config["Body"]
	result := Data(&config)
	fmt.Print(result)
}

func Data(_config *config.Config) error {
	msg := "From:" + _config.Sender + "\n" + "To:" + _config.To + "\n" + "Subject:" + _config.Subject + _config.Body 
    err := smtp.SendMail(fmt.Sprintf("%s:%s", _config.Host,_config.Port), 
		smtp.PlainAuth("", _config.Sender, _config.Password, _config.Host), 
		_config.Sender, []string{_config.To}, []byte(msg))

	fmt.Print(err)
	if err != nil {
		return err
		// log.Printf("smtp error: %s", err)
	}
	return err
}
