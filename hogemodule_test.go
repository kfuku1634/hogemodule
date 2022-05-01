package hogemodule 

import (
    "testing"
)

func TestHoge(t *testing.T) {
    want := "Hoge"
    if got := Hoge(); got != want {
        t.Errorf("Hoge() = %q, want %q", got, want)
    }
}
