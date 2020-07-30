# Go Struct To SQL 

Main purpose of this tool is to have a type safe SQL table creation.
This must support at least MySQL and SQLite. 


## Usage 

Import the lib and use it to initialize your DataBase

```go
package db

import (
	"github.com/everycheck/gostruct-to-sql/generator"
	"model"
	"fmt"
)

func Connect(url string) (*sql.DB, error) {
	db, err := sql.Open(protocol, url)
	if err != nil {
		return nil, fmt.Errorf("Can open conn to %s : %w",url,err)
	}
	g := generator.Generator{}
	g.RegisterType(model.Type1{})
	g.RegisterType(model.Type2{})
	ok = g.isUpToDate(db)
	if ok {
		return db, nil
	}

	query, err := g.Generate()	
	if err != nil {
		return nil, fmt.Errorf("Cannot generate database schema : %w",err)
	}
	_, err := db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("Error while creating schema : %w",err)
	}

	return db,nil
}
```


## Current Feature

 - Can create one table from anonyme struct
 - Can get the real name of a struct
 - Can add field : only supported type is Int as `integer`

## RoadMap

 - Add more type (all defined in the type enum)
 - Add option nullbale if type is a pointer 
 - Allow multi table creation
 - Add foreign key by referencing another type (one to one)/(many to one)
 - Add foreign key by referencing another type as an array (one to many)/(many to many)
 - Add option for auto increment




