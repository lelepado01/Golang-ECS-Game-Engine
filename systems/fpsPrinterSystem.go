package systems

import (
	"simpleRender/ecs"
	"simpleRender/engine"
	"strconv"
	"time"
)

type FPSPrinterSystem struct {
	entities []ecs.BasicEntity
	Engine   *engine.Engine
	Title    string
	lastTime int
	frames   int64
}

func (s *FPSPrinterSystem) Add(b *ecs.BasicEntity) {}

func (s *FPSPrinterSystem) Remove(b ecs.BasicEntity) {}

func (s *FPSPrinterSystem) Update(delta float32) {
	now := time.Now().Second()
	s.frames += 1

	if now-s.lastTime > 1 {
		s.Engine.SetTitle(s.Title + "	fps: " + strconv.FormatInt(s.frames, 10))
		s.lastTime = now
		s.frames = 0
	}
}
