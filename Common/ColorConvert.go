package Common

import (
	"image/color"
	"math"
	"sync"

	"github.com/hajimehoshi/ebiten"
)

type colorMCacheKey uint32

type colorMCacheEntry struct {
	m     ebiten.ColorM
	atime int64
}

var (
	textM          sync.Mutex
	colorMCache    = map[colorMCacheKey]*colorMCacheEntry{}
	emptyColorM    ebiten.ColorM
	monotonicClock int64
	cacheLimit     = 512
)

func init() {
	emptyColorM.Scale(0, 0, 0, 0)
}

func now() int64 {
	monotonicClock++
	return monotonicClock
}

// ColorToColorM converts a normal color to a color matrix
func ColorToColorM(clr color.Color) ebiten.ColorM {
	// RGBA() is in [0 - 0xffff]. Adjust them in [0 - 0xff].
	cr, cg, cb, ca := clr.RGBA()
	cr >>= 8
	cg >>= 8
	cb >>= 8
	ca >>= 8
	if ca == 0 {
		return emptyColorM
	}
	key := colorMCacheKey(uint32(cr) | (uint32(cg) << 8) | (uint32(cb) << 16) | (uint32(ca) << 24))
	e, ok := colorMCache[key]
	if ok {
		e.atime = now()
		return e.m
	}
	if len(colorMCache) > cacheLimit {
		oldest := int64(math.MaxInt64)
		oldestKey := colorMCacheKey(0)
		for key, c := range colorMCache {
			if c.atime < oldest {
				oldestKey = key
				oldest = c.atime
			}
		}
		delete(colorMCache, oldestKey)
	}

	cm := ebiten.ColorM{}
	rf := float64(cr) / float64(ca)
	gf := float64(cg) / float64(ca)
	bf := float64(cb) / float64(ca)
	af := float64(ca) / 0xff
	cm.Scale(rf, gf, bf, af)
	e = &colorMCacheEntry{
		m:     cm,
		atime: now(),
	}
	colorMCache[key] = e

	return e.m
}
