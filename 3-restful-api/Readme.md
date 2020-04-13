# Restful API
Project inspired in this tutorial https://www.youtube.com/watch?v=SonwZ6MF5BE

### Running

```
go run 3-restful-api/main.go
```

### API call examples

#### Create Book
```
POST /book 
{
	"isbn":"382938",
	"title": "My Beautiful Book",
	"author": {
		"firstName": "John",
		"lastName": "Black"
	}
}
```

#### Get Books
```
GET /books
```

#### Get Book
```
GET /book/{id} 
```

#### Update Book
```
PUT /book/{id} 
{
	"isbn":"382938",
	"title": "My Beautiful Book",
	"author": {
		"firstName": "John",
		"lastName": "Black"
	}
}
```

#### Delete Book
```
DELETE /book/{id} 
```