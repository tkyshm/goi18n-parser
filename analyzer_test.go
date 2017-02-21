package goi18np

import (
	"testing"
)

func TestDefaultAnalyzerAnalyzeFromFile(t *testing.T) {
	a := DefaultAnalyzer{}

	rs := a.AnalyzeFromFile("test/sample_file.go")
	expected := []string{
		"sample_uniq_key",
		"sample_uniq_key_2",
		"sample_uniq_key_3",
	}
	if rs[0].ID != expected[0] {
		t.Errorf("mismatch I18NRecord ID:\ngot=%s\nexpected=%s", rs[0].ID, expected[0])
	}

	if rs[1].ID != expected[1] {
		t.Errorf("mismatch I18NRecord ID:\ngot=%s\nexpected=%s", rs[1].ID, expected[1])
	}

	if rs[2].ID != expected[2] {
		t.Errorf("mismatch I18NRecord ID:\ngot=%s\nexpected=%s", rs[2].ID, expected[2])
	}
}

func TestDefaultAnalyzerAnalyzeFromFiles(t *testing.T) {
	a := DefaultAnalyzer{}

	files := []string{"test/sample_file.go", "test/sample_file.go"}
	rs := a.AnalyzeFromFiles(files)
	if len(rs) != 3 {
		t.Errorf("mismatch I18NRecord length:\ngot=%d\nexpected=3", len(rs))
	}
}
