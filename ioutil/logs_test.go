package ioutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimeParse(t *testing.T) {
	datetime := TimeParse("01/Nov/2021:00:00:00 +0800")
	assert.NotNil(t, datetime)
	assert.Equal(t, int64(1635696000000), datetime.UnixMilli())
}