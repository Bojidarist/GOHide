package algos

import (
	"strings"
	"image"
	"os"
	"image/draw"
	"image/color"
	_ "image/jpeg"
	"image/png"
)

func loadImageRGBA(imgPath string) *image.RGBA {
	f, err := os.Open(imgPath)
	if err != nil{
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil{
		panic(err)
	}

	dimg, ok := img.(*image.RGBA)
	if !ok {
		bounds := img.Bounds()
		dimg = image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
		draw.Draw(dimg, dimg.Bounds(), img, bounds.Min, draw.Src)
	}

	return dimg
}

func saveImagePNG(img *image.RGBA, savePath string) {
	f, err := os.Create(savePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func stringToASCIIInt(str string) []int {
	runes := []rune(str)
	var result []int
	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}

	return result
}

// SimpleAlgo Super simple algorithm :thinking: 
func SimpleAlgoEncode(imgPath, saveImgPath, str string) {
	img := loadImageRGBA(imgPath)
	strValues := stringToASCIIInt(str)
	strIndex := 0

	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			if strIndex < len(strValues) {
				r,g,b,a := img.At(x, y).RGBA()
				a = 255 - uint32(strValues[strIndex])
				img.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
				strIndex = strIndex + 1
			} else {
				y = 9999999
				x = 9999999
			}
		}
	}
	saveImagePNG(img, saveImgPath)
}

func SimpleAlgoDecode(imgPath string) string {
	img := loadImageRGBA(imgPath)
	var buffer strings.Builder

	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			_,_,_,a := img.At(x, y).RGBA()
			if a>>8 == 255 { return buffer.String() }
			buffer.WriteString(string(rune(255 - (a>>8))))
		}
	}

	return buffer.String()
}
