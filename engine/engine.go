package engine

import (
	"simpleRender/colorlib"
	"simpleRender/ecs"
	"simpleRender/mathlib"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Engine is the object responsible for rendering
type Engine struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	bgColor  colorlib.ColorRGBA
	world    ecs.World
}

const _width = 800
const _height = 800

var lastFrameTime float64

// CreateEngine returns a new engine with default configurations
func CreateEngine() Engine {
	window, err := sdl.CreateWindow(
		"Engine",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		_width, _height,
		sdl.WINDOW_OPENGL)

	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	e := Engine{
		window:   window,
		renderer: renderer,
		bgColor:  colorlib.CreateColor(255, 255, 255, 255),
		world:    ecs.World{},
	}

	img.Init(img.INIT_JPG | img.INIT_PNG)
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	return e
}

// SetTitle sets the title of the window
func (e *Engine) SetTitle(name string) {
	e.window.SetTitle(name)
}

// SetBackground sets the background color of the window
func (e *Engine) SetBackground(r uint8, g uint8, b uint8, a uint8) {
	e.bgColor = colorlib.CreateColor(r, g, b, a)
}

func (e *Engine) updateDelta(lastFrameTime *float64) float64 {
	currentTime := float64(time.Now().Nanosecond()) / 1000000.0
	delta := currentTime - *lastFrameTime
	*lastFrameTime = currentTime

	if delta < 0 {
		delta = 1
	}

	return delta
}

// GameLoop starts the rendering loop
func (e *Engine) GameLoop() {
	defer e.destroy()

	for {
		delta := e.updateDelta(&lastFrameTime)

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		e.renderer.SetDrawColor(e.bgColor.GetValues())
		e.renderer.Clear()

		e.renderer.SetDrawColor(0, 0, 0, 255)
		e.world.Update(float32(delta))

		e.renderer.Present()
	}
}

// AddSystemToWorld adds the system to the world queue
func (e *Engine) AddSystemToWorld(s ecs.System) {
	e.world.AddSystem(s)
}

func (e *Engine) DrawTexture(tex *sdl.Texture, r *mathlib.Rectangle) {
	e.renderer.CopyEx(tex, nil, r.GetSDLRectangle(), 0.0, nil, sdl.FLIP_NONE)
}

func (e *Engine) destroy() {
	e.window.Destroy()
	e.renderer.Destroy()
}
