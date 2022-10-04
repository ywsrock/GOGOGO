package utils

import (
	"fmt"
	"log"
	"testing"

	"github.com/go-playground/assert/v2"
)

//var src = `C:\Users\en.b\Desktop\test\src\aa.log`
//var dst = `C:\Users\en.b\Desktop\test\dst\aa.log`
var src = `C:\Users\en.b\Desktop\test\src\`
var dst = `C:\Users\en.b\Desktop\test\dst\`

func TestFileCopy(t *testing.T) {
	f, _ := NewFS()
	err := f.FileCopy(src, dst)
	t.Run("No1", func(t *testing.T) {
		log.Println(err)
		assert.IsEqual(err, nil)
	})
}

func TestDirCopy(t *testing.T) {

	var src = `C:\Users\en.b\Desktop\test\src\`
	var dst = `C:\Users\en.b\Desktop\test\dst\`
	fmt.Printf("--->%s--->%s\n", src, dst)
	f, _ := NewFS()
	err := f.DirCopy(src, dst)
	t.Run("No1", func(t *testing.T) {
		log.Println(err)
		assert.IsEqual(err, nil)
	})

}
