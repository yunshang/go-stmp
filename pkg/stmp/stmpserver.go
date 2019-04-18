package stmp
import (
	"log"
	"bufio"
    "fmt"
    "os"
	"regexp"
	"go-stmp/pkg/utils"
	"go-stmp/pkg/config"
	"text/template"
	"bytes"
	"strconv"
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
	mailConfig, err := ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	config.Sender = strconv.Quote(mailConfig["Sender"])
	config.To = strconv.Quote(mailConfig["To"])
	config.Subject = strconv.Quote(mailConfig["Subject"])
	config.Body = strconv.Quote(mailConfig["Body"])
	result := Data(&config)
	fmt.Print(result)
}

// Data fill send mail data
func Data(_config *config.Config) error {
	buffer := new(bytes.Buffer)
	template := template.Must(template.New("emailTemplate").Parse(emailScript()))
    template.Execute(buffer, &_config)
	fmt.Print(_config)

    err := smtp.SendMail(fmt.Sprintf("%s:%s", _config.Host, _config.Port), 
		smtp.PlainAuth("", _config.Sender, _config.Password, _config.Host),_config.Sender, []string{_config.To},  buffer.Bytes())

	fmt.Print(err)
	if err != nil {
		return err
	}
	return err
}

// emailScript email template parse
func emailScript() (script string) {
    return "Sender: {{.Sender}}<br /> To: {{.To}}<br /> Subject: {{.Subject}}<br /> MIME-version: 1.0<br /> Content-Type: text/html; charset=&quot;UTF-8&quot;<br /> <br /> {{.Body}}"
}