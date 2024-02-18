package source

import (
	"dal/process/driver"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestParseDSN(t *testing.T) {
	// mysql://root:123456@tcp://localhost:3306,localhost:3305/schema/test?xx=1&b=2
	// mysql://tcp://localhost:3306,localhost:3305/schema/test?xx=1&b=2
	// mysql://tcp://localhost:3306/schema/test?xx=1&b=2
	// mysql://tcp://localhost:3306/test?xx=1&b=2
	bs, err := os.ReadFile("../../samples/source_1")
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := NewSource(string(bs)).parseDSN()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(cfg)
}

func TestMySQL(t *testing.T) {

	bs, err := os.ReadFile("../../samples/source_1")
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := NewSource(string(bs)).parseDSN()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(cfg)

	list, err := driver.NewDriver(cfg).Exec("select u.name ,dm.name as c ,dm.id as id  from data_monitors dm  left join users u on dm.creator_id =u.id")
	if err != nil {
		fmt.Println(err)
	}
	j, _ := json.Marshal(list)
	fmt.Println(string(j))
}
