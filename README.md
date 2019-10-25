sqltime
===

A `time.Time` wrapper compatible with databases `timestamp` type.

Issue 
-----
Most database `timestamp`  have a resolution of microseconds while Go `time.Time` has a resolution of nanoseconds.

[postgres](https://www.postgresql.org/docs/9.1/datatype-datetime.html)
[mysql](https://dev.mysql.com/doc/refman/8.0/en/fractional-seconds.html)

The resolution difference causes a data loss, so when a record is inserted into the database and retrieve the timestamp differs.

For testing, it gets quite annoying since you can't use `reflect.DeepEqual` to compare the two record ( original and the one fetch from the database)

Moreover,  most of the time, there is the issue of Location. `sqltime` will set the `timestamp` to the right database location. The default database location is `UTC,` but it can easily be changed with:

```go
sqltime.DatabaseLocation, _ = time.LoadLocation([YOUR_LOCATION])
```

Solution 
---
Wrapping the `time.Time` type to truncate the time to database resolution. By default, it will truncate the nanoseconds

The resolution can be changed with:

```go
sqltime.Truncate = time.Microsecond
```

---

example
----

see a full example [here](/example/postgres.go). The example assumes that the database is set
 to the default timezone of `UTC` if not please update the `sqltime.DatabaseLocation`.



Usage with with GORM
---

It is particularly useful with ORM's like [GORM](https://github.com/jinzhu/gorm)

But instead of extending the `gorm.Model` you declare your base.

```go
package Model

import "github.com/SamuelTissot/sqltime"

// define the model
type BaseModel struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt sqltime.Time  `gorm:"type:timestamp"`
	UpdatedAt sqltime.Time  `gorm:"type:timestamp"`
	DeletedAt *sqltime.Time `gorm:"type:timestamp"`
}

// and use it like this
type MyModel struct {
	BaseModel
	Data string
}

```