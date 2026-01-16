# imgutils-rotate

[![Go Reference](https://pkg.go.dev/badge/github.com/imgutils-org/imgutils-rotate.svg)](https://pkg.go.dev/github.com/imgutils-org/imgutils-rotate)
[![Go Report Card](https://goreportcard.com/badge/github.com/imgutils-org/imgutils-rotate)](https://goreportcard.com/report/github.com/imgutils-org/imgutils-rotate)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Go library for rotating and flipping images. Part of the [imgutils](https://github.com/imgutils-org) collection.

## Features

- 90, 180, 270 degree rotations
- Arbitrary angle rotation
- Horizontal and vertical flipping
- Configurable background color for angle rotations
- Fast implementations for common rotations

## Installation

```bash
go get github.com/imgutils-org/imgutils-rotate
```

## Quick Start

```go
package main

import (
    "image"
    "os"

    "github.com/imgutils-org/imgutils-rotate"
)

func main() {
    // Open image
    file, _ := os.Open("photo.jpg")
    defer file.Close()
    src, _, _ := image.Decode(file)

    // Rotate 90 degrees clockwise
    rotated := rotate.Rotate90(src)

    // Save result
    out, _ := os.Create("rotated.jpg")
    defer out.Close()
    rotate.SaveJPEG(rotated, out, 85)
}
```

## Usage Examples

### Standard Rotations

```go
// Rotate 90 degrees clockwise
rotated := rotate.Rotate90(src)

// Rotate 180 degrees
rotated := rotate.Rotate180(src)

// Rotate 270 degrees clockwise (90 counter-clockwise)
rotated := rotate.Rotate270(src)
```

### Flipping

```go
// Mirror horizontally (flip left-right)
flipped := rotate.FlipHorizontal(src)

// Flip vertically (flip top-bottom)
flipped := rotate.FlipVertical(src)
```

### Arbitrary Angle Rotation

```go
// Rotate 45 degrees with white background
rotated := rotate.RotateAngle(src, 45, color.White)

// Rotate 30 degrees with transparent background
rotated := rotate.RotateAngle(src, 30, color.Transparent)

// Rotate -15 degrees (counter-clockwise)
rotated := rotate.RotateAngle(src, -15, color.Black)
```

### Rotate from File

```go
// Load and rotate in one step
rotated, err := rotate.RotateFromFile("photo.jpg", 90)
if err != nil {
    log.Fatal(err)
}

// For arbitrary angles, use RotateAngle after loading
```

## API Reference

### Functions

| Function | Description |
|----------|-------------|
| `Rotate90(src)` | Rotate 90 degrees clockwise |
| `Rotate180(src)` | Rotate 180 degrees |
| `Rotate270(src)` | Rotate 270 degrees clockwise (90 degrees CCW) |
| `FlipHorizontal(src)` | Mirror horizontally |
| `FlipVertical(src)` | Flip vertically |
| `RotateAngle(src, angle, bg)` | Rotate by arbitrary angle |
| `RotateFromFile(path, degrees)` | Load and rotate (90/180/270) |
| `SaveJPEG(img, w, quality)` | Save as JPEG |
| `SavePNG(img, w)` | Save as PNG |

## Performance Notes

- `Rotate90`, `Rotate180`, `Rotate270` are optimized and very fast
- `FlipHorizontal` and `FlipVertical` are also optimized
- `RotateAngle` is slower as it performs per-pixel calculations

## Common Use Cases

### Fix Phone Photo Orientation

```go
// EXIF orientation 6 = rotate 90 degrees CW
fixed := rotate.Rotate90(src)

// EXIF orientation 3 = rotate 180 degrees
fixed := rotate.Rotate180(src)

// EXIF orientation 8 = rotate 270 degrees CW
fixed := rotate.Rotate270(src)
```

### Create Mirror Effect

```go
mirrored := rotate.FlipHorizontal(src)
```

### Artistic Rotation

```go
// Slight tilt effect
tilted := rotate.RotateAngle(src, 5, color.White)
```

## Requirements

- Go 1.16 or later

## Related Packages

- [imgutils-crop](https://github.com/imgutils-org/imgutils-crop) - Image cropping
- [imgutils-resize](https://github.com/imgutils-org/imgutils-resize) - Image resizing
- [imgutils-sdk](https://github.com/imgutils-org/imgutils-sdk) - Unified SDK

## License

MIT License - see [LICENSE](LICENSE) for details.
