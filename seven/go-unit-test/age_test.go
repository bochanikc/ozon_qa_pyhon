package unit

import "testing"

func Test_age(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{
			n:    0,
			want: "",
		},
		{
			n:    1,
			want: "baby",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := age(tt.n)

			if got != tt.want {
				t.Errorf("age() = %v, want %v", got, tt.want)
			}
		})
	}
}
