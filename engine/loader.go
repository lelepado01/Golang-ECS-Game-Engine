package engine

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func LoadBMP(path string, e *Engine) *sdl.Texture {
	img, err := sdl.LoadBMP(path)
	if err != nil {
		fmt.Println("loading player error")
		fmt.Println(err)
		os.Exit(1)
	}
	defer img.Free()

	playerTexture, err := e.renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Println("texture creation error")
		fmt.Println(err)
		os.Exit(1)
	}

	return playerTexture
}

func LoadPNG(path string, e *Engine) *sdl.Texture {
	surfaceImg, err := img.Load(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		os.Exit(4)
	}

	textureImg, err := e.renderer.CreateTextureFromSurface(surfaceImg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		os.Exit(5)
	}
	surfaceImg.Free()

	return textureImg
}
