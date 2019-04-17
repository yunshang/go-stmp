package stmp
import (
	"log"
	"bufio"
    "fmt"
    "os"
	"regexp"
	"go-stmp/pkg/utils"
)

func ReadFile(name string) {
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

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
