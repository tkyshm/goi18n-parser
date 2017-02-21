package goi18np

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestDefaultAnalyzerAnalyzeFromFile(t *testing.T) {
	expected := []string{
		"sample_uniq_key",
		"sample_uniq_key_2",
		"sample_uniq_key_3",
		"sample_uniq_key_4",
		"sample_uniq_key_5",
	}

	a := DefaultAnalyzer{}
	rs := a.AnalyzeFromFile("test/sample_file.go")

	if rs[0].ID != expected[0] {
		t.Errorf("mismatch I18NRecord ID:\ngot=%s\nexpected=%s", rs[0].ID, expected[0])
	}

	if rs[1].ID != expected[1] {
		t.Errorf("mismatch I18NRecord ID:\ngot=%s\nexpected=%s", rs[1].ID, expected[1])
	}

	if rs[2].ID != expected[2] {
		t.Errorf("mismatch I18NRecord ID:\ngot=%s\nexpected=%s", rs[2].ID, expected[2])
	}

	if rs[3].ID != expected[3] {
		t.Errorf("mismatch I18NRecord ID:\ngot=%s\nexpected=%s", rs[3].ID, expected[3])
	}

	if rs[4].ID != expected[4] {
		t.Errorf("mismatch I18NRecord ID:\ngot=%s\nexpected=%s", rs[4].ID, expected[4])
	}
}

func TestDefaultAnalyzerAnalyzeFromFiles(t *testing.T) {
	files := []string{"test/sample_file.go", "test/sample_file.go"}

	a := DefaultAnalyzer{}
	rs := a.AnalyzeFromFiles(files)

	if len(rs) != 5 {
		t.Errorf("mismatch I18NRecord length:\ngot=%d\nexpected=3", len(rs))
	}
}

func TestSaveJSON(t *testing.T) {
	tmp, err := ioutil.TempFile("", "save_test")
	if err != nil {
		t.Error(err)
	}
	tmp.Close()
	defer os.Remove(tmp.Name())

	a := DefaultAnalyzer{}
	rs := a.AnalyzeFromFile("test/sample_file.go")

	a.SaveJSON(tmp.Name())

	data, err := ioutil.ReadFile(tmp.Name())
	if err != nil {
		t.Error(err)
	}
	var got []I18NRecord
	if err = json.Unmarshal(data, &got); err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, rs) {
		t.Errorf("Mismatched saved json and analyzed data:\ngot=%#v\nexpected=%#v", got, rs)
	}
}
