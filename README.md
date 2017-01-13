# qrest
Create a rest api from mariadb for a database. This project was motivated by [prest](https://github.com/nuveo/prest) by [nuveo](https://github.com/nuveo)

## The Works:

To install run,
```
go install github.com/vyasgiridhar/qrest/qrest
```

To start the rest server,
```
qrest -rport 8000 -host 127.0.0.1 -mport 3306 -user *** -pass *** -database ***
```

## GET:
```
/table?page=2&pagesize=10
/table?field=value
```
Where table is the table name in the database.

Returns the result in json format.

## PUT:
```
/table
```

Request body:
```
{
    "FIELD1": "string value",
    "FIELD2": 1234567890
}
```

Insert the JSON data in the request body into the table

## POST:

Working on Deletion.