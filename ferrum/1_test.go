package ferrum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpreadsheetNotation(t *testing.T) {
	assert.Equal(t, "1A", getSpreadsheetNotation(1))
	assert.Equal(t, "1AA", getSpreadsheetNotation(27))
	assert.Equal(t, "1ZZ", getSpreadsheetNotation(702))
	assert.Equal(t, "2A", getSpreadsheetNotation(703))
	assert.Equal(t, "3AA", getSpreadsheetNotation(1431))
	assert.Equal(t, "91973644572YF", getSpreadsheetNotation(64565498489498))
}
