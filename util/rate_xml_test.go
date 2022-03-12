package util

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRateFromXML(t *testing.T) {
	f, err := os.Open("rate_test.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	body, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	rate, err := ParseRateFromXML(body)
	assert.NoError(t, err)

	assert.Equal(t, "Reference rates", rate.Subject)
	assert.Equal(t, "European Central Bank", rate.Sender.Name)
	assert.Equal(t, 65, len(rate.Cube.Cubes))

	cubes := rate.Cube.Cubes
	assert.Equal(t, "2022-03-11", cubes[0].Date)

	items := cubes[0].Cubes
	assert.Equal(t, 31, len(items))
	assert.Equal(t, "USD", items[0].Currency)
	assert.Equal(t, 1.099, items[0].Rate)
}
