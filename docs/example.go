package main

import (
    "fmt"
    "github.com/K-minutti/payload"
  )

var data = `{
  "timestamp": "{{.Timestamp}}",
  "id": "{{.GUID}}",
  "data": {
      "name": "{{.RandomFirstName}} {{.RandomLastName}}",
      "date": "{{.RandomDatePast}}",
  }
}
`

func main() {
    message := payload.Hello("Person")
    fmt.Println(message)

    s := payload.Parse(data)
    fmt.Println(s)

    // api payload test

    // function - struct level test

    // generate w/ target
    // insomnia - {% uuid 'v4' %}
    // postman - {{$guid}}
}
