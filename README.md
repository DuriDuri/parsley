# Parsley's Medical Database API

This application presents a database-driven API that writes and reads simple patient objects.

### Features

This application allows basic CRUD operations on the simple patient object.

```
POST /patient
GET /patient/:id
PATCH /patient/:id
DELETE /patient/:id
```

Lists all patient objects saved in the database
```
GET /patients/list   
```

The API insures that all incoming (inserts/updates) requests are processed if they have the correct data type. In addition, it ensures that fields such as `email`, `zipcode` etc. are in the correct format.

### Patient object

```json
{
  "firstName": "Radric",
  "middleName": "Delantic",
  "lastName": "Davis",
  "phones": [{
    "type": "Mobile",
    "number": "5554443333"
  }],
  "email": "guwop@fakehost.test",
  "dob": "1980-02-12",  
  "age": 37,
  "gender": "male",
  "status": "active",
  "termsAccepted": true,
  "termsAcceptedAt": "2018-01-03T10:00:00Z",
  "address": {
    "line1": "123 Main St",
    "line2": "Apt 10",
    "city": "Atlanta",
    "state": "GA",
    "zip": "30363"
  }
}
```


### Running service

Clone repository and use Docker Compose to build and launch the application containers

### Architecture

Data is written in Go and writes to a MongoDB. Database connections are explicitly cloned and timely closed to minimize the blocking nature Mongo is infamous for. 

The applications processor is split into the following entities:

* API -> Define routes and their associated handlers
* DB -> Define a swappable DB. Mongo is used in this example
* Models -> Defines the simple patient object and the validation contract
* Handlers -> Defines the API logic when interacting with data
