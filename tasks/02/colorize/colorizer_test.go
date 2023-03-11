package colorize

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"image"
	"image/color"
	"image/png"
	"os"
	"path"
	"testing"
)

func TestDecomposeGRB(t *testing.T) {
	for _, filename := range []string{"testdata/gopher-in-ok.png", "testdata/rgb_model.png"} {
		filename := filename
		t.Run(filename, func(t *testing.T) {
			t.Parallel()

			f, err := os.Open(filename)
			require.NoError(t, err)
			img, err := png.Decode(f)
			require.NoError(t, err)

			dec := DecomposeGRB(img)

			require.Equal(t, color.GrayModel, dec.Red.ColorModel())
			require.Equal(t, color.GrayModel, dec.Green.ColorModel())
			require.Equal(t, color.GrayModel, dec.Blue.ColorModel())

			r, g, b := toColorsFiles(filename)
			checkFileContent(t, r, dec.Red)
			checkFileContent(t, g, dec.Green)
			checkFileContent(t, b, dec.Blue)
		})
	}
}

func TestComposeRGB(t *testing.T) {
	for _, filename := range []string{"testdata/gopher-in-ok.png", "testdata/rgb_model.png"} {
		filename := filename
		t.Run(filename, func(t *testing.T) {
			t.Parallel()

			r, g, b := toColorsFiles(filename)
			var dec DecomposedImage

			f, err := os.Open(r)
			require.NoError(t, err)
			dec.Red, err = png.Decode(f)
			require.NoError(t, err)

			f, err = os.Open(g)
			require.NoError(t, err)
			dec.Green, err = png.Decode(f)
			require.NoError(t, err)

			f, err = os.Open(b)
			require.NoError(t, err)
			dec.Blue, err = png.Decode(f)
			require.NoError(t, err)

			require.Equal(t, color.GrayModel, dec.Red.ColorModel())
			require.Equal(t, color.GrayModel, dec.Green.ColorModel())
			require.Equal(t, color.GrayModel, dec.Blue.ColorModel())

			composed := ComposeRGB(dec)
			require.Equal(t, color.RGBAModel, composed.ColorModel())
			checkFileContent(t, filename, composed)
		})
	}
}

func checkFileContent(t *testing.T, filename string, img image.Image) {
	content, err := os.ReadFile(filename)
	require.NoError(t, err)

	var result = bytes.NewBuffer(nil)
	err = png.Encode(result, img)
	require.NoError(t, err)

	require.Equal(t, content, result.Bytes())
}

func toColorsFiles(filename string) (r string, g string, b string) {
	var (
		ext        = path.Ext(filename)
		withoutExt = filename[:len(filename)-len(ext)]
	)

	return withoutExt + "_red" + ext, withoutExt + "_green" + ext, withoutExt + "_blue" + ext
}
