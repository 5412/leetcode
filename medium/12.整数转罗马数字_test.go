package medium

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "3",
			args: args{
				num: 3,
			},
			want: "III",
		},
		{
			name: "4",
			args: args{
				num: 4,
			},
			want: "IV",
		},
		{
			name: "9",
			args: args{
				num: 9,
			},
			want: "IX",
		},
		{
			name: "58",
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			name: "1994",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
		{
			name: "100",
			args: args{
				num: 100,
			},
			want: "C",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
