package healthchecksio

import "testing"

func Test_generateSlugFromName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{name: "test"},
			want: "test",
		},
		{
			name: "with spaces",
			args: args{name: "test check"},
			want: "test-check",
		},
		{
			name: "with special characters",
			args: args{name: "test-check!"},
			want: "test-check",
		},
		{
			name: "with multiple spaces",
			args: args{name: "test  check"},
			want: "test-check",
		},
		{
			name: "with multiple special characters",
			args: args{name: "test!-check-123!!"},
			want: "test-check-123",
		},
		{
			name: "with trailing spaces",
			args: args{name: "test check "},
			want: "test-check",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateSlugFromName(tt.args.name); got != tt.want {
				t.Errorf("generateSlugFromName() = %v, want %v", got, tt.want)
			}
		})
	}
}
