# Restful API

### Running

```
go run 4-restful-api-with-database/main.go
```

### API endpoints

#### Create User
```
POST /user 
{
	"username":"johnsmith",
	"firstname": "John",
	"lastname": "Smith"
}
```

#### Get Users
```
GET /users 
```

#### Get User
```
GET /user/{id} 
```

#### Update Book
```
PUT /user 
{
	"username":"johnsmith",
	"firstname": "John",
	"lastname": "Smith"
}
```

#### Delete Book
```
DELETE /user/{id} 
```