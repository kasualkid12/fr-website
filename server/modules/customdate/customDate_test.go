package customdate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCustomDate_UnmarshalJSON(t *testing.T) {
	jsonDate := `"2024-05-24"`
	var cd CustomDate

	err := cd.UnmarshalJSON([]byte(jsonDate))
	assert.NoError(t, err)
	assert.Equal(t, 2024, cd.Year())
	assert.Equal(t, time.May, cd.Month())
	assert.Equal(t, 24, cd.Day())
}

func TestCustomDate_MarshalJSON(t *testing.T) {
	cd := CustomDate{Time: time.Date(2024, 5, 24, 0, 0, 0, 0, time.UTC)}
	expectedJSON := `"2024-05-24"`

	jsonBytes, err := cd.MarshalJSON()
	assert.NoError(t, err)
	assert.JSONEq(t, expectedJSON, string(jsonBytes))
}

func TestCustomDate_Scan(t *testing.T) {
	var cd CustomDate

	// Test with time.Time
	tm := time.Date(2024, 5, 24, 0, 0, 0, 0, time.UTC)
	err := cd.Scan(tm)
	assert.NoError(t, err)
	assert.Equal(t, 2024, cd.Year())
	assert.Equal(t, time.May, cd.Month())
	assert.Equal(t, 24, cd.Day())

	// Test with string
	err = cd.Scan("2024-05-24")
	assert.NoError(t, err)
	assert.Equal(t, 2024, cd.Year())
	assert.Equal(t, time.May, cd.Month())
	assert.Equal(t, 24, cd.Day())

	// Test with []byte
	err = cd.Scan([]byte("2024-05-24"))
	assert.NoError(t, err)
	assert.Equal(t, 2024, cd.Year())
	assert.Equal(t, time.May, cd.Month())
	assert.Equal(t, 24, cd.Day())

	// Test with nil
	err = cd.Scan(nil)
	assert.NoError(t, err)
	assert.True(t, cd.Time.IsZero())

	// Test with unsupported type
	err = cd.Scan(12345)
	assert.Error(t, err)
	assert.Equal(t, "unsupported scan type for CustomDate", err.Error())
}
