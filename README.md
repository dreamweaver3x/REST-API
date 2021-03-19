# REST API
Hi. This is my REST API server. I'm using Echo as Go web framework and GORM as ORM with PostgreSQL database.
 So, this server is made to keep the info about your ADs. 
 
 Client can POST date, number of views, clicks and cost, information stores in database. Also it counts cost per click(cpc) and cost per mille (cost*1000/views, cpm).
 
 Client can GET info sorted by date or other parametr and DELETE everything stored in database.
 
In this README i'll tell you
- what every package does
- how to communicate with this server
- how to run it on your pc

_____
## Packages
### models
Here is stored struct which GORM uses to create table in database. 
For `date` i made my own type `CutsomDate`, because default format for `Time.Time` in GO is not, what i want, so `CustomDate` looks like `YYYY-MM-DD`.

For `Cost`,`Cpm` and `cpc` i used `decimal` type, because float is unreliable, when we want to work with money.

`func InitModels(db *gorm.DB) error` - this func migrates DB, when we run our server.
_____
### dateMarshaller
This package is mode for new type i made. There are methods to make proper work with JSON
- `func (c *CustomDate) UnmarshalJSON(b []byte) (err error)`
- `func (c CustomDate) MarshalJSON() ([]byte, error)`
- 
Also methods for proper work with GORM
- `func (c CustomDate) Value() (driver.Value, error)`
- `func (c *CustomDate) Scan(v interface{}) error`
_____
### repository
All communication with database stored in this package. There are 3 methods: 
- `Create(stats *models.Stats)` which adds new info in database.
- `GetStats(from, to dateMarshaller.CustomDate, order string)` - bring to client info he needs between date `from` and `to`, ordered by whatever client asks.
- `DeleteFromDB()` - deletes everything in DB.
_____
### app
Here is our server and all 3 handlers.
```golang
e.POST("/save", a.AddNewStats)
``` 
Uses `func (a *Application) AddNewStats(c echo.Context) error`, it Unmarshall request (works with JSON) and sends it to DataBase.
```golang
e.GET("/stats", a.GetStats)
``` 
Uses `func (a *Application) GetStats(c echo.Context) error` it Unmarshall request, and calls `GetStats` from repository, then sends to client info he needs in JSON.
```golang
e.DELETE("/delete", a.DeleteStats)
``` 
Calls ` func (a *Application) DeleteStats(c echo.Context) error`, which deletes all information from database.
### main
Here we initialize our Databasea and server.
### config
Our configuration for database + port.
___
## How to communicate with server
### POST
if you want to send info to database, use POST method 
example: 
`curl -X POST http://0.0.0.0:8080/save \
-H 'Content-Type: application/json' \
-d '{"date": "2100-07-19","views": 7000, "cost": 6111, "clicks": 400}'`

(or `http://localhost:8080/save`, depends on how you run the server)
You'll get relpy with 202 code and message "Stats added"
### GET
If you want to get info from database, use GET method, with date `from` and `to` + order.
example: `curl -X GET http://0.0.0.0:8080/stats \
-H 'Content-Type: application/json' \
-d '{"from":"2000-01-01", "to":"3000-01-01", "order":"clicks"}'` 

(or `http://localhost:8080/save`, depends on how you run the server)
You'll get reply with 200 code and info from database in JSON format (order works improperly with `cost`, `cpm` and `cpc`, i'll fix it later)
### DELETE
To clear database, use method DELETE
example: `curl -X DELETE http://0.0.0.0:8080/delete`
You'll get reply with 202 code and message "Stats have been deleted."
_____
## How to RUN
- Clone this repository
- Use `docker-compose up`, enjoy
    - to communicate with your server use URL `http://0.0.0.0:8080`
*  If you want to run it without docker, add `.env` file with your environment configuration, for example: `APP_NAME = api_server
DB_DSN = "host=localhost user=db_user password=pwd123 dbname=stats port=54320 sslmode=disable"
PORT = 8080` and type in terminal `go run cmd/app/main.go -dev=true`.

Flag must be `true`, so server will take environment config from `.env`
