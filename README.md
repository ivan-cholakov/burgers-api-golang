# go-rollercoaster-api


## Requirements

To be able to show the desired features of curl this REST API must match a few
requirements:

* [x] `GET /burgers` returns list of burgers as JSON
* [x] `GET /burgers?burger_name={burger_name}` returns a burger with the specified name
* [x] `GET /burgers/{id}` returns details of specific burger as JSON
* [x] `GET /burgers/random` redirects (Status 302) to a random burger
* [x] `POST /burgers` accepts a new burger to be added

### Data Types

A burger object should look like this:
```json
{
  "id": "someid",
  "name": "name of the burger",
  "description": "description of the burger",
  "restaurant": "the restaurant offering the burger",
  "web": "the restaurant's website",
  "image_url": "link to an image representing the burger",
  "addresses": "array of the restaurant's locations",
  "ingredients": "array of the burger's ingredients"
}
```

### Persistence
Data is persisted inside a MongoDb collection

### How to run the project
```code
  docker-compose up
```
or
```code
  go run main.go
```
