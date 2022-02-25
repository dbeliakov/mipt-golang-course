package timepng

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFillWithMask(t *testing.T) {
	var img = image.NewRGBA(image.Rect(0, 0, 6, 10))
	var mask = []int{
		1, 0, 1,
		0, 1, 0,
		1, 0, 1,
		0, 1, 0,
		1, 0, 1,
	}
	fillWithMask(img, mask, color.Black, 2)
	for i, val := range mask {
		var x, y = i % 3, i / 3
		var c color.RGBA
		if val == 1 {
			c = color.RGBA{
				A: 255,
			}
		}
		assert.Equal(t, c, img.At(2 * x, 2 * y))
		assert.Equal(t, c, img.At(2 * x + 1, 2 * y))
		assert.Equal(t, c, img.At(2 * x, 2 * y + 1))
		assert.Equal(t, c, img.At(2 * x + 1, 2 * y + 1))
	}
}

func TestBuildTimeImage(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			img := buildTimeImage(tc.Time, tc.Color, scale)
			var b bytes.Buffer
			require.NoError(t, png.Encode(&b, img))

			f, err := ioutil.ReadFile(tc.File)
			require.NoError(t, err)
			require.Equal(t, f, b.Bytes())
		})
	}
}

func TestTimePNG(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			var b bytes.Buffer
			TimePNG(&b, tc.Time, tc.Color, scale)

			f, err := ioutil.ReadFile(tc.File)
			require.NoError(t, err)
			require.Equal(t, f, b.Bytes())
		})
	}
}

const scale = 10

var testCases = []struct{
	Name string
	Time time.Time
	Color color.RGBA
	File string
}{
	{
		Name: "01",
		Time: parseTime( "19:34"),
		Color: color.RGBA{
			R: 100,
			G: 100,
			B: 255,
			A: 255,
		},
		File: "testdata/01.png",
	},
	{
		Name: "02",
		Time: parseTime( "23:59"),
		Color: color.RGBA{
			R: 100,
			G: 100,
			B: 255,
			A: 255,
		},
		File: "testdata/02.png",
	},
	{
		Name: "03",
		Time: parseTime( "20:48"),
		Color: color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		},
		File: "testdata/03.png",
	},
}

func parseTime(tStr string) time.Time {
	t, err := time.Parse("15:04", tStr)
	if err != nil {
		panic("invalid tests")
	}
	return t
}