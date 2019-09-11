package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/LakshmiMekala/rules-1/ruleapi/tests"
	"github.com/project-flogo/core/engine"
)

func BenchmarkTestSimpleJSON(b *testing.B) {

	data, err := ioutil.ReadFile(filepath.FromSlash("./flogo.json"))
	if err != nil {
		b.Fatal(err)
	}
	cfg, err := engine.LoadAppConfig(string(data), false)
	if err != nil {
		b.Fatal(err)
	}
	e, err := engine.New(cfg)
	if err != nil {
		b.Fatal(err)
	}
	if e == nil {
		b.Fatal("failed to create app engine")
	}

	e.Start()
	defer func() {
		e.Stop()
		e = nil
	}()

	// b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tests.Command("curl", "http://localhost:7777/test/n1?name=Bob")
		fmt.Println(n, err)

	}

}
