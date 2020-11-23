package builder

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGoBuildWasm(t *testing.T) {
	prjDir := "/Users/tschinke/git/github.com/golangee/forms-example/www/"
	hash, err := HashFileTree(prjDir)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", hex.EncodeToString(hash))

	root, err := GoEnv("GOROOT")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", root)

	err = GoBuildWasm(prjDir, "bla.wasm")
	if err != nil {
		t.Fatal(err)
	}
}
