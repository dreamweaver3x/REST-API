package dateMarshaller

import (
	"database/sql/driver"
	"strings"
	"time"
	"fmt"
)
const layout = "2006-01-02"

type CustomDate struct {
	time.Time
}

func (c *CustomDate) UnmarshalJSON(b []byte) (err error) {

	s := strings.Trim(string(b), "\"") // remove quotes
	if s == "null" {
		return
	}
	c.Time, err = time.Parse(layout, s)
	return
}
func (c CustomDate) MarshalJSON() ([]byte, error) {
	if c.Time.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, c.Time.Format(layout))), nil
}

func (c CustomDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if c.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return c.Time, nil
}

func (c *CustomDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*c = CustomDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}