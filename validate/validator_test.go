package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	source := []byte(`
document:
  dsl: 1.0.0-alpha1
  namespace: examples
  name: call-http-shorthand-endpoint
  version: 1.0.0-alpha1
do:
  - test:
      call: http
      with:
        method: get
        endpoint: https://petstore.swagger.io/v2/pet/{petId}
`)
	err := FromYAMLSource(source)
	assert.NoError(t, err)

}
