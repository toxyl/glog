package glog

import "testing"

func TestToken(t *testing.T) {
	type args struct {
		tokens []string
		glue   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Token Test 1",
			args: args{
				glue: " ",
				tokens: []string{
					"123",
					"Hello world, we have 5 more tokens:",
					"123, 43.2464 or 43.2346?",
					"Need 14.23% or '424.14%' or something else?",
					"To be or not to be: true or false?",
					"It's nil again!",
					"This was the last of 6 tokens!",
				},
			},
			want: "[38;5;87m123[0m[38;5;119m [0m[38;5;119m[38;5;125mHello world, we have[0m[38;5;125m [0m[38;5;87m5[0m[38;5;125m [0m[38;5;125mmore tokens:[0m[38;5;125m[0m [38;5;87m123, 43.2464[0m[38;5;162m [0m[38;5;162mor[0m[38;5;162m [0m[38;5;87m43.2346[0m[38;5;162m[0m[38;5;162m?[0m[38;5;162m[0m [38;5;137mNeed[0m[38;5;137m [0m[38;5;85m14.23%[0m[38;5;137m [0m[38;5;137mor '[0m[38;5;137m[0m[38;5;85m424.14%[0m[38;5;137m[0m[38;5;137m' or something else?[0m[38;5;137m[0m [38;5;145mTo be or not to be:[0m[38;5;145m [0m[38;5;82mtrue[0m[38;5;145m [0m[38;5;145mor[0m[38;5;145m [0m[38;5;196mfalse[0m[38;5;145m[0m[38;5;145m?[0m[38;5;145m[0m [38;5;138mIt's[0m[38;5;138m [0m[38;5;160mnil[0m[38;5;138m [0m[38;5;138magain![0m[38;5;138m[0m [38;5;126mThis was the last of[0m[38;5;126m [0m[38;5;87m6[0m[38;5;126m [0m[38;5;126mtokens![0m[38;5;126m[0m[0m[38;5;119m[0m",
		},
		{
			name: "Token Test 2",
			args: args{
				glue: " ",
				tokens: []string{
					"The Protestant church was built in 1230 as a replacement of a 12th century predecessor. It has been enlarged and altered multiple times and restored between 1974 and 1976. [5]",
				},
			},
			want: "[38;5;138m[38;5;127mThe Protestant church was built in[0m[38;5;127m [0m[38;5;87m1230[0m[38;5;127m [0m[38;5;127mas a replacement of a 12th century predecessor. It has been enlarged and altered multiple times and restored between[0m[38;5;127m [0m[38;5;87m1974[0m[38;5;127m [0m[38;5;127mand[0m[38;5;127m [0m[38;5;87m1976[0m[38;5;127m[0m[38;5;127m. [[0m[38;5;127m[0m[38;5;87m5[0m[38;5;127m[0m[38;5;127m][0m[38;5;127m[0m[0m[38;5;138m[0m",
		},
		{
			name: "Token Test 3",
			args: args{
				glue: " ",
				tokens: []string{
					"Several estates were built near Oentsjerk, however only Stania State has remained. Eysinga State has become a retirement home.[6] The stins Stania State was probably built in the 16th century. The current estate dates from 1843. Around 1520, it was turned into a castle-like building. In 1546, it became a property of the van Heemstra family who owned it for two centuries, but let it deteriorate. In 1738, it was demolished. In 1843, a new manor house by Looxma who had made a fortune in oil mills. A large garden designed by Lucas Roodbaard [nl] is laid out. In the 1930s, it became a youth hostel, conference centre and contained an outpost of the Fries Museum. In the 1970s, it was bought by the municipality who later sold it a company as an office building.[7] Oentsjerk was home to 441 people in 1840.[6] In the late-20th century it became a commuter's village.[8]",
				},
			},
			want: "",
		},
		{
			name: "Token Test 4",
			args: args{
				glue: " ",
				tokens: []string{
					"hello world, this is a string that should be split on spaces, it contains a bunch of different tokens like -123, 0, 123, -43.2464, 0, 43.2464, -14.23%, 0%, 14.23%, false, true, nil and this will look different. We made 45% profit and lost 124%!",
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Token(tt.args.glue, tt.args.tokens...); got != tt.want {
				t.Errorf("Token() = %v, want %v", got, tt.want)
				t.FailNow()
			}
		})
	}
}
