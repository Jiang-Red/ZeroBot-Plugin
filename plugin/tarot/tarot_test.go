package tarot

import "testing"

func TestSplitTextChunks(t *testing.T) {
	got := splitTextChunks("甲乙丙丁", 3)
	if len(got) != 2 {
		t.Fatalf("splitTextChunks() chunks = %d, want 2", len(got))
	}
	if got[0] != "甲乙丙" || got[1] != "丁" {
		t.Fatalf("splitTextChunks() = %#v, want []string{\"甲乙丙\", \"丁\"}", got)
	}
}

func TestBuildMessage(t *testing.T) {
	msg := makeNodeMessage("结果", "占卜者", 1)
	if len(msg) != 1 {
		t.Fatalf("buildMessage() message segments = %d, want 1", len(msg))
	}
}

func TestDrawResultsAnalyzeSignature(t *testing.T) {
	var analyze func(drawResults, string, string, float32) (string, error) = drawResults.analyze
	if analyze == nil {
		t.Fatal("drawResults.analyze is nil")
	}
}
