# Hotel reservation backend


## Project Outline
- users -> book room from a hotel
- admins -> going to check reservation/bookings
- Authentication and authorization -> JWT tokens
- Hotles -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> database management -> seeding, migrating


## Resources
### Mongodb driver
Doucmentation
```
https://mongodb.com/docs/drivers/go/current/quick-start
````

Installing mongodb client
```
go get go.mongodb.org/mongo-driver/mongo
```

### gofiber
Documentation
``` 
https://gofiber.io
```

Installing gofiber
```
go get github.com/gofiber/fiber/v2
```

## Docker
### Installing mongodb as a Docker container
```
docker run --name mongodb -d mongo:latest -p 27017:27017
```
