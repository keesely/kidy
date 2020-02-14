// config_test.go kee > 2020/02/14

package config

import (
	//"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	var bc = "kee"
	var data = map[string]interface{}{
		"a": "hello world",
		"b": map[string]interface{}{
			"c": bc,
		},
	}

	cnf := NewConfig(data)
	t.Log("Current a:", cnf.Get("a"))
	if cnf.Get("a") != data["a"] {
		t.Error("get a fail")
	}

	t.Log("Current b.c:", cnf.Get("b.c"))
	if cnf.Get("b.c") != bc {
		t.Error("get b.c fail")
	}

	t.Log("Is Not Exists c:", cnf.Get("c", "no exists"))
	if cnf.Get("c") != "no exists" {
		t.Error("get no value (c) fail")
	}

	t.Log("=====================================================================")
	cnf.Set("b.a", "this is a append to b.a")
	t.Log("Set b.a:", cnf.Get("b.a"))
	if cnf.Get("b.a") != "this is a append to b.a" {
		t.Error("fail b.a")
	}

	cnf.Set("b.c", "Update B/C")
	t.Log("Set b.c", cnf.Get("b.c"))
	if cnf.Get("b.c") == bc {
		t.Error("fail update b.c")
	}

	cnf.Set("c.d", "cd append")
	t.Log("Set c.d", cnf.Get("c.d"))
	if cnf.Get("c.d") != "cd append" {
		t.Error("fail append c.d")
	}

	t.Log("deep Set")
	cnf.Set("aaa.bbb.ccc.ddd.eee.fff", map[string]interface{}{"echo": "hello world"})

	dv := cnf.Get("aaa.bbb.ccc.ddd").(*Config)
	t.Log("AAA->BBB->CCC->DDD:", dv.Get("eee.fff.echo"))

	cnf.Set("deep", dv)
	t.Log("DEEP VALUE: ", cnf.Get("deep.eee.fff.echo"))

	t.Log(">> ALL >>\n", cnf.GetAll())

	t.Log(">>>>>>> TO JSON: \n", string(cnf.ToJSON()))
	t.Log(">>>>>>> TO YAML: \n", string(cnf.ToYAML()))
}
