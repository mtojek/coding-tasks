package coding_tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangeMoney_NoMoney(t *testing.T) {
	quarters, dimes, nickels, pennies := ChangeMoney(0)
	assert.Equal(t, quarters,0)
	assert.Equal(t, dimes,0)
	assert.Equal(t, nickels,0)
	assert.Equal(t, pennies,0)
}

func TestChangeMoney_OnePenny(t *testing.T) {
	quarters, dimes, nickels, pennies := ChangeMoney(1)
	assert.Equal(t, quarters,0)
	assert.Equal(t, dimes,0)
	assert.Equal(t, nickels,0)
	assert.Equal(t, pennies,1)
}

func TestChangeMoney_284pennies(t *testing.T) {
	quarters, dimes, nickels, pennies := ChangeMoney(284)
	assert.Equal(t, quarters,11)
	assert.Equal(t, dimes,0)
	assert.Equal(t, nickels,1)
	assert.Equal(t, pennies,4)
}


func TestChangeMoney_99pennies(t *testing.T) {
	quarters, dimes, nickels, pennies := ChangeMoney(99)
	assert.Equal(t, quarters,3)
	assert.Equal(t, dimes,2)
	assert.Equal(t, nickels,0)
	assert.Equal(t, pennies,4)
}