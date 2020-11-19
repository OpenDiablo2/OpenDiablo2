package d2systems

import (
	"github.com/gravestench/akara"
	"image/color"
	"math"
	"testing"
	"time"
)

func Test_RenderSystem(t *testing.T) {
	cfg := akara.NewWorldConfig()

	renderSys := NewRenderSystem()

	cfg.With(renderSys)

	world := akara.NewWorld(cfg)

	gameConfigEntity := world.NewEntity()

	renderSys.AddGameConfig(gameConfigEntity)

	vpEntity := world.NewEntity()
	vp := renderSys.AddViewport(vpEntity)

	vp.Width = 400
	vp.Height = 300

	renderSys.AddMainViewport(vpEntity)

	sfc := renderSys.AddSurface(vpEntity)

	loadAttempts := 10
	for loadAttempts > 0 {
		if renderSys.renderer != nil {
			break
		}

		_ = world.Update(0)

		loadAttempts--
	}

	if loadAttempts < 0 {
		t.Fatal("could not create renderer")
	}

	sfc.Surface = renderSys.renderer.NewSurface(400, 300)
	sfc.Surface.DrawRect(20, 20, color.RGBA{100, 100, 255, 255})

	go func(){
		x, y := 0.0, 0.0

		for {
			ms := float64(world.TimeDelta.Milliseconds())
			x, y = x+ms/1000, y+ms/3000
			vp.Top = int(math.Abs(math.Sin(y) * 300))
			vp.Left = int(math.Abs(math.Cos(x) * 400))
			time.Sleep(time.Second/60)
		}
	}()

	err := renderSys.Loop()
	if err != nil {
		t.Fatal(err)
	}
}
