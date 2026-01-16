// Package rotate provides image rotation and flipping utilities.
package rotate

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"math"
	"os"
)

// Rotate90 rotates an image 90 degrees clockwise.
func Rotate90(src image.Image) image.Image {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := image.NewRGBA(image.Rect(0, 0, h, w))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dst.Set(h-1-y, x, src.At(x+bounds.Min.X, y+bounds.Min.Y))
		}
	}
	return dst
}

// Rotate180 rotates an image 180 degrees.
func Rotate180(src image.Image) image.Image {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dst.Set(w-1-x, h-1-y, src.At(x+bounds.Min.X, y+bounds.Min.Y))
		}
	}
	return dst
}

// Rotate270 rotates an image 270 degrees clockwise (90 counter-clockwise).
func Rotate270(src image.Image) image.Image {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := image.NewRGBA(image.Rect(0, 0, h, w))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dst.Set(y, w-1-x, src.At(x+bounds.Min.X, y+bounds.Min.Y))
		}
	}
	return dst
}

// FlipHorizontal flips an image horizontally (mirror).
func FlipHorizontal(src image.Image) image.Image {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dst.Set(w-1-x, y, src.At(x+bounds.Min.X, y+bounds.Min.Y))
		}
	}
	return dst
}

// FlipVertical flips an image vertically.
func FlipVertical(src image.Image) image.Image {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dst.Set(x, h-1-y, src.At(x+bounds.Min.X, y+bounds.Min.Y))
		}
	}
	return dst
}

// RotateAngle rotates an image by an arbitrary angle (in degrees).
// The background color fills any exposed areas.
func RotateAngle(src image.Image, angle float64, bg color.Color) image.Image {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	// Convert angle to radians
	rad := angle * math.Pi / 180

	// Calculate new dimensions
	sin, cos := math.Abs(math.Sin(rad)), math.Abs(math.Cos(rad))
	newW := int(float64(w)*cos + float64(h)*sin)
	newH := int(float64(w)*sin + float64(h)*cos)

	dst := image.NewRGBA(image.Rect(0, 0, newW, newH))

	// Fill with background color
	draw.Draw(dst, dst.Bounds(), &image.Uniform{bg}, image.Point{}, draw.Src)

	// Calculate center points
	srcCX, srcCY := float64(w)/2, float64(h)/2
	dstCX, dstCY := float64(newW)/2, float64(newH)/2

	sinA, cosA := math.Sin(-rad), math.Cos(-rad)

	for y := 0; y < newH; y++ {
		for x := 0; x < newW; x++ {
			// Translate to origin, rotate, translate back
			dx := float64(x) - dstCX
			dy := float64(y) - dstCY
			srcX := int(dx*cosA-dy*sinA+srcCX) + bounds.Min.X
			srcY := int(dx*sinA+dy*cosA+srcCY) + bounds.Min.Y

			if srcX >= bounds.Min.X && srcX < bounds.Max.X &&
				srcY >= bounds.Min.Y && srcY < bounds.Max.Y {
				dst.Set(x, y, src.At(srcX, srcY))
			}
		}
	}
	return dst
}

// RotateFromFile reads an image file and applies a rotation.
func RotateFromFile(path string, degrees int) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	switch degrees {
	case 90:
		return Rotate90(src), nil
	case 180:
		return Rotate180(src), nil
	case 270:
		return Rotate270(src), nil
	default:
		return RotateAngle(src, float64(degrees), color.White), nil
	}
}

// SaveJPEG saves the rotated image as JPEG.
func SaveJPEG(img image.Image, w io.Writer, quality int) error {
	if quality <= 0 || quality > 100 {
		quality = 85
	}
	return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
}

// SavePNG saves the rotated image as PNG.
func SavePNG(img image.Image, w io.Writer) error {
	return png.Encode(w, img)
}
