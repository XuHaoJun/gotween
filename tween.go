package tween

import (
	"time"
)

type Tween struct {
	// main
	space       *Space
	object      map[string]float64
	valuesEnd   map[string]float64
	duration    time.Duration
	curDuration time.Duration
	// etc
	isPlaying bool
	startTime time.Duration
	delayTime time.Duration
	// easing function
	easingFunc func(k float64) float64
	// callbacks
	onStartFuncFired bool
	onStartFunc      func(*Tween)
	onStopFunc       func(*Tween)
	onUpdateFunc     func(*Tween)
	onCompleteFunc   func(*Tween)
}

func NewTween(obj map[string]float64) *Tween {
	t := &Tween{
		object:      obj,
		easingFunc:  Easing.Linear.None,
		duration:    1 * time.Second,
		curDuration: 1 * time.Second,
	}
	return t
}

func (t *Tween) To(valuesEnd map[string]float64, dur time.Duration) *Tween {
	t.valuesEnd = valuesEnd
	t.duration = dur
	return t
}

func (t *Tween) Easing(f func(k float64) float64) *Tween {
	t.easingFunc = f
	return t
}

func (t *Tween) Start(d time.Duration) *Tween {
	return t
}

func (t *Tween) Stop() *Tween {
	if !t.isPlaying {
		return t
	}
	t.space.Remove(t)
	t.isPlaying = false
	if t.onStopFunc != nil {
		t.onStopFunc(t)
	}
	return t
}

func (t *Tween) Update(delta time.Duration) bool {
	if t.onStartFuncFired == false {
		if t.onStartFunc != nil {
			t.onStartFunc(t)
		}
		t.onStartFuncFired = true
	}
	value := t.easingFunc(float64(delta) / float64(time.Second))
	value = value * float64(t.duration/delta)
	for key, end := range t.valuesEnd {
		vector := end - t.object[key]
		if vector > 0 {
			t.object[key] += value
		} else {
			t.object[key] -= value
		}
	}
	if t.onUpdateFunc != nil {
		t.onUpdateFunc(t)
	}
	t.curDuration -= delta
	if t.curDuration < 0 {
		return false
	}
	return true
}

func (t *Tween) Delay(delay time.Duration) *Tween {
	t.delayTime = delay
	return t
}

func (t *Tween) OnStart(f func(*Tween)) *Tween {
	t.onStartFunc = f
	return t
}

func (t *Tween) OnStop(f func(*Tween)) *Tween {
	t.onStopFunc = f
	return t
}

func (t *Tween) OnUpdate(f func(*Tween)) *Tween {
	t.onUpdateFunc = f
	return t
}

func (t *Tween) OnComplete(f func(*Tween)) *Tween {
	t.onCompleteFunc = f
	return t
}
