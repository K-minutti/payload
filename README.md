# Payload
A module for payload mocking

Generate fake data for fields in json

```json
{
    "timestamp": "{{$timestamp}}",
    "id": "{{$guid}}",
    "data": {
        "name": "{{$randomFirstName}} {{$randomLastName}}",
        "date": "{{$randomDatePast}}",
    }
}
```

```json
{
    "timestamp": "{{.Timestamp}}",
    "id": "{{.GUID}}",
    "data": {
        "name": "{{.RandomFirstName}} {{.RandomLastName}}",
        "date": "{{.RandomDatePast}}",
    }
}
```

use in tests

random generation syntax

json composition syntax