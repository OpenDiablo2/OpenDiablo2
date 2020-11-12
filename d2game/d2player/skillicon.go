package d2player

import (
	"log"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const (
	skillLabelXOffset = 49
	skillLabelYOffset = -4

	skillIconXOff  = 346
	skillIconYOff  = 59
	skillIconDistX = 69
	skillIconDistY = 68
)

type skillIcon struct {
	*d2ui.BaseWidget
	lvlLabel *d2ui.Label
	sprite   *d2ui.Sprite
	skill    *d2hero.HeroSkill
}

func newSkillIcon(ui *d2ui.UIManager, baseSprite *d2ui.Sprite, skill *d2hero.HeroSkill) *skillIcon {
	base := d2ui.NewBaseWidget(ui)
	label := ui.NewLabel(d2resource.Font16, d2resource.PaletteSky)

	x := skillIconXOff + skill.SkillColumn*skillIconDistX
	y := skillIconYOff + skill.SkillRow*skillIconDistY

	res := &skillIcon{
		BaseWidget: base,
		sprite:     baseSprite,
		skill:      skill,
		lvlLabel:   label,
	}

	res.SetPosition(x, y)

	return res
}

func (si *skillIcon) SetVisible(visible bool) {
	si.BaseWidget.SetVisible(visible)
	si.lvlLabel.SetVisible(visible)
}

func (si *skillIcon) renderSprite(target d2interface.Surface) {
	x, y := si.GetPosition()

	if err := si.sprite.SetCurrentFrame(si.skill.IconCel); err != nil {
		log.Printf("Cannot set Frame %e", err)
		return
	}

	if si.skill.SkillPoints == 0 {
		target.PushSaturation(skillIconGreySat)
		defer target.Pop()

		target.PushBrightness(skillIconGreyBright)
		defer target.Pop()
	}

	si.sprite.SetPosition(x, y)
	si.sprite.Render(target)
}

func (si *skillIcon) renderSpriteLabel(target d2interface.Surface) {
	if si.skill.SkillPoints == 0 {
		return
	}

	x, y := si.GetPosition()
	si.lvlLabel.SetText(strconv.Itoa(si.skill.SkillPoints))
	si.lvlLabel.SetPosition(x+skillLabelXOffset, y+skillLabelYOffset)
	si.lvlLabel.Render(target)
}

func (si *skillIcon) Render(target d2interface.Surface) {
	si.renderSprite(target)
	si.renderSpriteLabel(target)
}

func (si *skillIcon) Advance(elapsed float64) error {
	return nil
}
