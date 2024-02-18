package process

import (
	"data-api/process/source"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"testing"
)

func initSource() {
	fsSource, err := antlr.NewFileStream("../samples/source_1")
	if err != nil {
		panic(err)
	}
	source.DefaultStore.Store(fsSource.String())
}

func TestTemplate(t *testing.T) {
	da := NewDal("")

	template := "select * from xxxx where a=? and b=? and c=? and d=? and e in ? and f in ?"
	values := []interface{}{
		"33",
		true,
		1,
		2.1,
		[]int{1, 2, 3},
		[]string{"xxx", "ff"},
	}
	fmt.Println(da.buildSql(template, values))
}

func TestStageExecution(t *testing.T) {
	initSource()
	fsApi, err := antlr.NewFileStream("../samples/api_1")
	if err != nil {
		t.Fatal(err)
	}

	content := fsApi.String()

	da := NewDal(content)

	da.addProcessVar("a", 1, "stage.test")

	da.addProcessVar("c", 3, "stage.test")

	if err := da.stageExecution(); err != nil {
		t.Fatal(err)
	}
	fmt.Println(da.getVarString())
}

func TestReturn(t *testing.T) {

	initSource()

	fs, err := antlr.NewFileStream("../samples/api_1")
	if err != nil {
		t.Fatal(err)
	}

	content := fs.String()

	da := NewDal(content)

	da.addProcessVar("a", 1, "request")

	fmt.Println(da.getVarString())

	c, err := da.getContent()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(c)
}
