package bb

import (
	"fmt"
	"testing"
)

func TestCreatePath(t *testing.T) {
	dir, file := makeFileName("/base")
	fmt.Println(dir, file)
}

func TestCreateWriter(t *testing.T) {
	fw := CreateWriter("/tmp/")
	defer fw.Close()

	fw.Write([]byte("TEST"))
}
