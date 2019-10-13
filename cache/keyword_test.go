package cache

import (
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/assert"
)

func TestItStoresAKeyword(t *testing.T) {
	kw := randomdata.SillyName()

	err := delKeyword(kw)
	assert.NoError(t, err)

	StoreKeyword(kw)

	count, err := getKeyword(kw)
	assert.NoError(t, err)

	assert.Equal(t, int64(1), count)
}
