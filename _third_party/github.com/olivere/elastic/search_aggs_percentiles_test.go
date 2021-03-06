package elastic

import (
	"encoding/json"
	"testing"
)

func TestPercentilesAggregation(t *testing.T) {
	agg := NewPercentilesAggregation().Field("price")
	data, err := json.Marshal(agg.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"percentiles":{"field":"price"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestPercentilesAggregationWithCustomPercents(t *testing.T) {
	agg := NewPercentilesAggregation().Field("price").Percentiles(0.2, 0.5, 0.9)
	data, err := json.Marshal(agg.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"percentiles":{"field":"price","percents":[0.2,0.5,0.9]}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
