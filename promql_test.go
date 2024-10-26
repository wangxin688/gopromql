package gopromql

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		name     string
		pb       *PromqlBuilder
		expected string
		err      error
	}{
		{
			name: "basic case",
			pb: &PromqlBuilder{
				metricName: "metric",
			},
			expected: "metric",
			err:      nil,
		},
		{
			name: "with labels",
			pb: &PromqlBuilder{
				metricName: "metric",
				labels: []Label{
					{
						Name:    "label1",
						Matcher: "=",
						Value:   "value1",
					},
					{
						Name:    "label2",
						Matcher: "!=",
						Value:   "value2",
					},
				},
			},
			expected: `metric{label1="value1",label2!="value2"}`,
			err:      nil,
		},
		{
			name: "with window",
			pb: &PromqlBuilder{
				metricName: "metric",
				window:     "5m",
			},
			expected: "metric[5m]",
			err:      nil,
		},
		{
			name: "with compare",
			pb: &PromqlBuilder{
				metricName: "metric",
				compOps: []Compare{
					{
						Op:    ">",
						Value: 10,
					},
					{
						Op:    "<",
						Value: 20,
					},
				},
			},
			expected: " metric>10 and metric<20",
			err:      nil,
		},
		{
			name: "with function",
			pb: &PromqlBuilder{
				metricName: "metric",
				funcName:   "sum",
			},
			expected: "sum(metric)",
			err:      nil,
		},
		{
			name: "with offset",
			pb: &PromqlBuilder{
				metricName: "metric",
				offset:     "5m",
			},
			expected: "metric offset 5m",
			err:      nil,
		},
		{
			name: "with aggregation",
			pb: &PromqlBuilder{
				metricName: "metric",
				agg: Aggregation{
					Op:     "sum",
					By:     []string{"label1", "label2"},
					AggWay: "by",
				},
			},
			expected: "sum(metric) by (label1,label2)",
			err:      nil,
		},
		{
			name: "with missing metric name",
			pb:   &PromqlBuilder{},
			err:  fmt.Errorf("metric name is required"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.pb.Build()
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("got error %v, want %v", err, tt.err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
