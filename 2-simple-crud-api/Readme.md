# A simple Book App CRUD


### Running

```
go run 2-simple-crud-api/main.go
```

### API call examples

#### Create Book
```
POST /book/create 
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
GET /book/list 
```

#### Get Book
```
GET /book/show?id={uuid} 
```

#### Update Book
```
UPDATE /book/update?id={uuid} 
```

#### Delete Book
```
DELETE /book/delete?id={uuid} 
```