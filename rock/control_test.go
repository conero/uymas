package rock

import (
	"encoding/json"
	"testing"
)

func TestMust(t *testing.T) {
	bys := Must(json.Marshal(map[string]any{"a": 1}))
	t.Logf("byts: %s", bys)
}
