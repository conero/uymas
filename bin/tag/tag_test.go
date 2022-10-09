package tag

import "testing"

func TestParseTag(t *testing.T) {
	var vTg = "option short:N,n require"
	tg := ParseTag(vTg)
	if tg.Type != CmdOption {
		t.Error("tag Type parse failure")
	}
	if !tg.IsRequired() {
		t.Error("option require is parse failure")
	}
	if tg.ValueString("short") != "N,n" {
		t.Error("the k-v pairs parse failure: short")
	}
}
