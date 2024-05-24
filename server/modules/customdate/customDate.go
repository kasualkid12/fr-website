package customdate

import (
	"encoding/json"
	"errors"
	"time"
)

type CustomDate struct {
	time.Time
}

func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	var err error
	cd.Time, err = time.Parse(`"2006-01-02"`, string(data))
	return err
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(cd.Format("2006-01-02"))
}

func (cd *CustomDate) Scan(value interface{}) error {
	if value == nil {
		*cd = CustomDate{Time: time.Time{}}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*cd = CustomDate{Time: v}
		return nil
	case []byte:
		t, err := time.Parse("2006-01-02", string(v))
		if err != nil {
			return err
		}
		*cd = CustomDate{Time: t}
		return nil
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		*cd = CustomDate{Time: t}
		return nil
	default:
		return errors.New("unsupported scan type for CustomDate")
	}
}
