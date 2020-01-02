package ping

import "testing"

func TestText(t *testing.T) {
    expected := "Ping"
    if actual := Text(); actual != expected {
        t.Errorf("Text() = Actual:[%q], Expected:[%q]", actual, expected)
    }
}

