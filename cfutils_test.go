package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseInt(t *testing.T) {
	assert.Equal(t, 2147483647, ParseInt("2147483647"))
	assert.Panics(t, func() {
		ParseInt("2147483648")
	})
}

func TestParseInt64(t *testing.T) {
	assert.Equal(t, 2147483647, ParseInt64("2147483647"))
	assert.Equal(t, 2147483648, ParseInt64("2147483648"))
}

func TestParseCfTime(t *testing.T) {
	assert.Equal(t, time.Date(2012, 10, 15, 19, 53, 53, 0, time.UTC).Unix(),
		ParseCfTime("2012/10/15 19:53:53 +0000").Unix())
}

func TestParseSlackTime(t *testing.T) {
	now := time.Now()
	slackNow := fmt.Sprintf("%d.0", now.Unix())
	assert.Equal(t, now.Unix(), ParseSlackTime(slackNow).Unix())
}

func TestParseSimpleDate(t *testing.T) {
	assert.Equal(t, SimpleDate{2014, 1, 18}, ParseSimpleDate("2014-01-18"))
}

func TestSimpleDateToTime(t *testing.T) {
	someDate := time.Date(2014, 1, 12, 0, 0, 0, 0, time.UTC)
	simpleSomeDate := SimpleDate{someDate.Year(), int(someDate.Month()), someDate.Day()}
	assert.Equal(t, someDate.Unix(), SimpleDateToTime(simpleSomeDate).Unix())
}
