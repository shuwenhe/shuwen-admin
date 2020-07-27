package zhihu

import "testing"

func TestRun(t *testing.T) {
	tests := []struct {
		url string
	}{
		{url: "https://www.zhihu.com/special/all"},
	}
	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			Run(tt.url)
		})
	}
}
