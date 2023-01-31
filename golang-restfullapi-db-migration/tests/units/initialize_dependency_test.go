package units

import (
	"golang-dependency-injection/sandboxes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeSimpleService(t *testing.T) {
	t.Run("ShouldErorr", func(t *testing.T) {
		simpleService, err := sandboxes.InitializeSimpeService(true)

		assert.Nil(t, simpleService)
		assert.NotNil(t, err)
	})
	t.Run("ShouldSuccess", func(t *testing.T) {
		simpleService, err := sandboxes.InitializeSimpeService(false)

		assert.NotNil(t, simpleService)
		assert.Nil(t, err)
	})
}
