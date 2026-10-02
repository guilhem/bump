package semver

import "testing"

func TestLatest(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "one tag",
			args: args{
				tags: []string{
					"v1.0.0",
				},
			},
			want:    "v1.0.0",
			wantErr: false,
		},
		{
			name: "two tags",
			args: args{
				tags: []string{
					"v1.0.0",
					"v2.0.0-rc.1",
				},
			},
			want:    "v2.0.0-rc.1",
			wantErr: false,
		},
		{
			name:    "no semantic version tags",
			args:    args{tags: []string{"release", "nightly"}},
			wantErr: true,
		},
		{
			name: "ignore non-version tags",
			args: args{tags: []string{"nightly", "v1.2.3", "release", "v1.3.0"}},
			want: "v1.3.0",
		},
		{
			name: "empty",
			args: args{
				tags: []string{},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Latest(tt.args.tags)
			if (err != nil) != tt.wantErr {
				t.Errorf("Latest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Latest() = %v, want %v", got, tt.want)
			}
		})
	}
}
