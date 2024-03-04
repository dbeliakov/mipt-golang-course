package encoder

import (
	"testing"
)

func TestCustomEncoder_EncodeDecode(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		mode       EncodeMode
		wantEncode string
		wantDecode string
	}{
		{
			name:       "Test DashMode with English",
			input:      "Hello",
			mode:       DashMode,
			wantEncode: "72-101-108-108-111",
			wantDecode: "Hello",
		},
		{
			name:       "Test UnderScoreMode with Russian",
			input:      "Привет",
			mode:       UnderScoreMode,
			wantEncode: "1055_1088_1080_1074_1077_1090",
			wantDecode: "Привет",
		},
		{
			name:       "Test DashMode with Numbers",
			input:      "123",
			mode:       DashMode,
			wantEncode: "49-50-51",
			wantDecode: "123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := NewCustomEncoder(tt.mode)
			gotEncoded := encoder.Encode(tt.input)
			if gotEncoded != tt.wantEncode {
				t.Errorf("CustomEncoder.Encode() got = %v, want %v", gotEncoded, tt.wantEncode)
			}

			gotDecoded := encoder.Decode(gotEncoded)
			if gotDecoded != tt.wantDecode {
				t.Errorf("CustomEncoder.Decode() got = %v, want %v", gotDecoded, tt.wantDecode)
			}
		})
	}
}
