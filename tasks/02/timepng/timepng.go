package timepng

import (
	"image"
	"image/color"
	"io"
	"time"
)

// TimePNG записывает в `out` картинку в формате png с текущим временем
func TimePNG(out io.Writer, t time.Time, c color.Color, scale int) {
	// TODO: Implement me
}

// buildTimeImage создает новое изображение с временем `t`
func buildTimeImage(t time.Time, c color.Color, scale int) *image.RGBA {
	// TODO: Implement me
	return nil
}

// fillWithMask заполняет изображение `img` цветом `c` по маске `mask`. Маска `mask`
// должна иметь пропорциональные размеры `img` с учетом фактора `scale`
// NOTE: Так как это вспомогательная функция, можно считать, что mask имеет размер (3x5)
func fillWithMask(img *image.RGBA, mask []int, c color.Color, scale int) {
	// TODO: implement me
}

var nums = map[rune][]int{
	'0': {
		1, 1, 1,
		1, 0, 1,
		1, 0, 1,
		1, 0, 1,
		1, 1, 1,
	},
	'1': {
		0, 1, 1,
		0, 0, 1,
		0, 0, 1,
		0, 0, 1,
		0, 0, 1,
	},
	'2': {
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
	},
	'3': {
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
	},
	'4': {
		1, 0, 1,
		1, 0, 1,
		1, 1, 1,
		0, 0, 1,
		0, 0, 1,
	},
	'5': {
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
	},
	'6': {
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
	},
	'7': {
		1, 1, 1,
		0, 0, 1,
		0, 1, 0,
		0, 1, 0,
		0, 1, 0,
	},
	'8': {
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
	},
	'9': {
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
	},
	':': {
		0, 0, 0,
		0, 1, 0,
		0, 0, 0,
		0, 1, 0,
		0, 0, 0,
		0, 0, 0,
	},
}