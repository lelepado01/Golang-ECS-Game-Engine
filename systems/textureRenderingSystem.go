package systems

import (
	"simpleRender/components"
	"simpleRender/ecs"
	"simpleRender/engine"
)

type TextureRenderingSystem struct {
	entities []TexturedEntity
	Engine   *engine.Engine
}

type TexturedEntity struct {
	ecs.BasicEntity
	components.RectComponent
	components.TextureComponent
}

func (s *TextureRenderingSystem) Add(b *ecs.BasicEntity, space *components.RectComponent, t *components.TextureComponent) {
	s.entities = append(s.entities, TexturedEntity{*b, *space, *t})
}

func (s *TextureRenderingSystem) Remove(b ecs.BasicEntity) {}

func (s *TextureRenderingSystem) Update(delta float32) {
	for _, entity := range s.entities {
		s.Engine.DrawTexture(
			entity.Texture,
			entity.Position)
	}
}
