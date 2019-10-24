sqltime
---

Particularly useful for testing

This came about to fix the issue that everytime I would test two model in the database with `reflect.DeepEqual` it would
failed all the time since the database truncated the time.Time value.

Integrates well wth [GORM](https://github.com/jinzhu/gorm)

**ATTENTION** : this type will truncate the value of time.Time resulting in a data loss of magniture of the value
of Truncate


**NOTE** 
It should work with mysql but it is untested

example
----

see a full example [here](/example/postgres.go). The example assumes that the database is set
 to the default timezone of ` UTC` if not please update the sqltime.DatabaseLocation.



Usage with with GORM
---
```go
// define the model
type BaseModel struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt sqltime.Time  `gorm:"type:timestamp"`
	UpdatedAt sqltime.Time  `gorm:"type:timestamp"`
	DeletedAt *sqltime.Time `gorm:"type:timestamp"`
}
```