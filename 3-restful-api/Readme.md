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
GET /book 
```

#### Get Book
```
GET /book/{id} 
```

#### Update Book
```
UPDATE /book/{id} 
```

#### Delete Book
```
DELETE /book/{id} 
```