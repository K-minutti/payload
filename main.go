package payload

import (
	"bytes"
	"fmt"
	"github.com/jaswdr/faker"
	"encoding/json"
	"os"
	"text/template"
	"time"
)

type Data struct {
	GUID            string
	Timestamp       string
	RandomFirstName string
	RandomLastName  string
	RandomDatePast  string
}

var f = faker.New()
var d = Data{
	GUID:            f.UUID().V4(),
	Timestamp:       f.Time().UnixDate(time.Now()),
	RandomFirstName: f.Person().FirstName(),
	RandomLastName:  f.Person().LastName(),
	RandomDatePast:  "Thu Jun 13 2019 03:08:43 GMT+0530",
}

func Hello(name string) string {
	msg := fmt.Sprintf("Hi, %v. Hello how are you?", name)
	return msg
}

func ParseWith(body, values string) string {
	t, err := template.New("x").Parse(body)
	if err != nil {
		panic(err)
	}

	m := map[string]interface{}{}
	if err := json.Unmarshal([]byte(values), &m); err != nil {
		panic(err)
	}

	var r bytes.Buffer

	err = t.Execute(&r, m)
	if err != nil {
		panic(err)
	}

	s := r.String()
	return s
}

func Parse(body string) string {
	t, err := template.New("x").Parse(body)
	if err != nil {
		panic(err)
	}

	var r bytes.Buffer

	err = t.Execute(&r, d)
	if err != nil {
		panic(err)
	}

	s := r.String()
	return s
}

func ParseDisplay(body string) {
	d := Data{
		GUID:            "611c2e81-2ccb-42d8-9ddc-2d0bfa65c1b4",
		Timestamp:       "1562757109",
		RandomFirstName: "Heaven",
		RandomLastName:  "Earth",
		RandomDatePast:  "Thu Jun 13 2019 03:08:43 GMT+0530",
	}

	t, err := template.New("x").Parse(body)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, d)
	if err != nil {
		panic(err)
	}
}
