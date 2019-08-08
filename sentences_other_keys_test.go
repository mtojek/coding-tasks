package coding_tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindOtherKeys(t *testing.T) {
	sentences := []string{
		"Abrar hrr hrr ifewjfosv",
		"Fewfj efie euf fue",
		"Ffe fqio qf fhq",
		"Fofeh bnq z wdfw",
		"Poknfa ufa b ai hrr wf",
	}
	keys := []string{
		"bnq",
		"ufa",
		"hrr",
	}

	unknown := FindOtherKeys(sentences, keys)
	assert.Len(t, unknown,9)
	assert.Contains(t, unknown, "Fofeh")
	assert.Contains(t, unknown, "z")
	assert.Contains(t, unknown, "wf")
	assert.Contains(t, unknown, "Poknfa")
	assert.Contains(t, unknown, "b")
	assert.Contains(t, unknown, "ai")
	assert.Contains(t, unknown, "Abrar")
	assert.Contains(t, unknown, "ifewjfosv")
	assert.Contains(t, unknown, "wdfw")
}
