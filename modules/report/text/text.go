package text

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/netm4ul/netm4ul/tests"

	"github.com/netm4ul/netm4ul/core/config"
	"github.com/netm4ul/netm4ul/core/database"
	"github.com/netm4ul/netm4ul/core/database/models"
)

type Text struct {
	Filename string
	Width    int
	DB       models.Database
	cfg      config.ConfigToml
	funcs    template.FuncMap
}

//NewReport returns a new initialized report struct
func NewReport() *Text {
	t := Text{}
	t.Width = 80

	err := config.LoadConfig("")
	if err != nil {
		panic("Couldn't load config file !")
	}

	t.cfg = config.Config
	t.DB = database.NewDatabase(&t.cfg)

	t.funcs = template.FuncMap{
		"Center": func(text string) string {
			return fmt.Sprintf(fmt.Sprintf("%%%ds", (len(text)+t.Width)/2), text)
		},
		"CenterWithFix": func(prefix, suffix, text string) string {
			return fmt.Sprintf(fmt.Sprintf("%%%ds", (len(prefix+text+suffix)+t.Width)/2), prefix+text+suffix)
		},
		"Pad": func(char string, padlen int) string {
			return strings.Repeat(char, padlen)
		},
		"FormatDate": func(ti time.Time) string {
			return ti.Format(time.Stamp)
		},
		"Suffix": func(text string, char string) string {
			return text + strings.Repeat(char, t.Width-len(text))
		},
	}

	return &t
}

func (t *Text) Name() string {
	return "Text"
}

const reportPath = "./reports"
const templatesPath = "modules/report/text/templates/"

//Generate a new report in text format
func (t *Text) Generate(name string) error {

	var buff bytes.Buffer
	data, err := t.getData()

	templates := []string{
		templatesPath + "toc.tmpl",
		templatesPath + "ips.tmpl",
		templatesPath + "ports.tmpl",
		templatesPath + "vulns.tmpl",
		templatesPath + "domains.tmpl",
		templatesPath + "index.tmpl",
	}

	if err != nil {
		return err
	}

	tmpl, err := template.New("Index").Funcs(t.funcs).ParseFiles(templates...)
	if err != nil {
		return err
	}
	err = tmpl.ExecuteTemplate(&buff, "index", data)
	if err != nil {
		return err
	}

	return WriteReport(name, buff.Bytes())
}

func (t *Text) getData() (map[string]interface{}, error) {

	var data map[string]interface{}
	data = make(map[string]interface{})

	data["Name"] = t.cfg.Project.Name
	data["Date"] = time.Now()
	data["Description"] = t.cfg.Project.Description
	//TOFIX Replace with real value !
	data["Domains"] = tests.NormalProject.Domains

	ips, err := t.DB.GetIPs(t.cfg.Project.Name)
	if err != nil {
		return nil, errors.New("Couldn't retrieve IPs from the database : " + err.Error())
	}

	data["IPs"] = ips
	data["Ports"] = make([]models.Port, 0)

	for _, ip := range ips {
		data["Ports"] = append(data["Ports"].([]models.Port), ip.Ports...)
	}

	return data, nil
}

//WriteReport will create a new report file. If the "reports" folder does not exist, it will create it.
func WriteReport(name string, data []byte) error {

	if _, err := os.Stat(reportPath); os.IsNotExist(err) {
		os.Mkdir(reportPath, 0755)
	}

	fullPath := path.Join(reportPath, name)
	err := ioutil.WriteFile(fullPath, data, 0600)

	if err != nil {
		return err
	}

	return nil
}
