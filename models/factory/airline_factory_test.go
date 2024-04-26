package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructAirline(t *testing.T) {
	airline := ConstructAirline()
	name := airline.Name
	updateCount := 2
	airline.SetCount(updateCount)
	assert.Equal(t, name, airline.Name, "Name before and after setting count are not the same.")
	assert.Equal(t, updateCount, airline.Count, "Count is not updated.")
}
