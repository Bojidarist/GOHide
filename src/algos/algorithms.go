package algos

import (
	"fmt"
	"image"
	"os"
	_ "image/jpeg"
	_ "image/png"
)

// SimpleAlgo Super simple algorithm :thinking: 
func SimpleAlgo(imgPath string) {
	f, err := os.Open(imgPath)
	if err != nil{
		fmt.Println(err)
	}
	defer f.Close()

	img, fmtName, err := image.Decode(f)
	if err != nil{
		fmt.Println(err)
	}

	bounds := img.Bounds()

	fmt.Println(fmtName)
	fmt.Println(bounds.Max.X)
	fmt.Println(bounds.Max.Y)
}
