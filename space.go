package tween

import (
	"time"
)

type Space struct {
	tweens map[*Tween]struct{}
}

func NewSpace() *Space {
	return &Space{
		tweens: make(map[*Tween]struct{}),
	}
}

func (s *Space) Tween(obj map[string]float64) *Tween {
	t := NewTween(obj)
	s.Add(t)
	return t
}

func (s *Space) Add(t *Tween) {
	s.tweens[t] = struct{}{}
	t.space = s
}

func (s *Space) Remove(t *Tween) {
	delete(s.tweens, t)
}

func (s *Space) Update(delta time.Duration) bool {
	if len(s.tweens) == 0 {
		return false
	}
	for tween, _ := range s.tweens {
		updateAble := tween.Update(delta)
		if updateAble == false {
			delete(s.tweens, tween)
		}
	}
	return true
}
