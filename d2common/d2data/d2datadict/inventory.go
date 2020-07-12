package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"log"
)

type box struct {
	Left   int
	Right  int
	Top    int
	Bottom int
	Width  int
	Height int
}

type grid struct {
	Box     *box
	Rows    int
	Columns int
	CellWidth int
	CellHeight int
}

type InventoryRecord struct {
	Name  string
	Panel *box
	Grid  *grid
	Slots map[d2enum.EquippedSlot]*box
}

var Inventory map[string]*InventoryRecord

func LoadInventory(file []byte) {
	d := d2common.LoadDataDictionary(file)
	Inventory = make(map[string]*InventoryRecord, 0)
	for d.Next() {
		// we need to calc the width/height for the box as it isn't
		// specified in the txt file
		pBox := &box{}
		pBox.Left = d.Number("invLeft")
		pBox.Right = d.Number("invRight")
		pBox.Top = d.Number("invTop")
		pBox.Bottom = d.Number("invBottom")
		pBox.Width = pBox.Right - pBox.Left
		pBox.Height = pBox.Bottom - pBox.Top

		gBox := &box{
			Left:   d.Number("gridLeft"),
			Right:  d.Number("gridRight"),
			Top:    d.Number("gridTop"),
			Bottom: d.Number("gridBottom"),
		}
		gBox.Width = gBox.Right - gBox.Left
		gBox.Height = gBox.Bottom - gBox.Top

		record := &InventoryRecord{
			Name:  d.String("class"),
			Panel: pBox,
			Grid: &grid{
				Box:     gBox,
				Rows:    d.Number("gridY"),
				Columns: d.Number("gridX"),
				CellWidth:  d.Number("gridBoxWidth"),
				CellHeight: d.Number("gridBoxHeight"),
			},
			Slots: map[d2enum.EquippedSlot]*box{
				d2enum.EquippedSlotHead: {
					d.Number("headLeft"),
					d.Number("headRight"),
					d.Number("headTop"),
					d.Number("headBottom"),
					d.Number("headWidth"),
					d.Number("headHeight"),
				},
				d2enum.EquippedSlotNeck: {
					d.Number("neckLeft"),
					d.Number("neckRight"),
					d.Number("neckTop"),
					d.Number("neckBottom"),
					d.Number("neckWidth"),
					d.Number("neckHeight"),
				},
				d2enum.EquippedSlotTorso: {
					d.Number("torsoLeft"),
					d.Number("torsoRight"),
					d.Number("torsoTop"),
					d.Number("torsoBottom"),
					d.Number("torsoWidth"),
					d.Number("torsoHeight"),
				},
				d2enum.EquippedSlotLeftArm: {
					d.Number("lArmLeft"),
					d.Number("lArmRight"),
					d.Number("lArmTop"),
					d.Number("lArmBottom"),
					d.Number("lArmWidth"),
					d.Number("lArmHeight"),
				},
				d2enum.EquippedSlotRightArm: {
					d.Number("rArmLeft"),
					d.Number("rArmRight"),
					d.Number("rArmTop"),
					d.Number("rArmBottom"),
					d.Number("rArmWidth"),
					d.Number("rArmHeight"),
				},
				d2enum.EquippedSlotLeftHand: {
					d.Number("lHandLeft"),
					d.Number("lHandRight"),
					d.Number("lHandTop"),
					d.Number("lHandBottom"),
					d.Number("lHandWidth"),
					d.Number("lHandHeight"),
				},
				d2enum.EquippedSlotRightHand: {
					d.Number("rHandLeft"),
					d.Number("rHandRight"),
					d.Number("rHandTop"),
					d.Number("rHandBottom"),
					d.Number("rHandWidth"),
					d.Number("rHandHeight"),
				},
				d2enum.EquippedSlotGloves: {
					d.Number("glovesLeft"),
					d.Number("glovesRight"),
					d.Number("glovesTop"),
					d.Number("glovesBottom"),
					d.Number("glovesWidth"),
					d.Number("glovesHeight"),
				},
				d2enum.EquippedSlotBelt: {
					d.Number("beltLeft"),
					d.Number("beltRight"),
					d.Number("beltTop"),
					d.Number("beltBottom"),
					d.Number("beltWidth"),
					d.Number("beltHeight"),
				},
				d2enum.EquippedSlotLegs: {
					d.Number("feetLeft"),
					d.Number("feetRight"),
					d.Number("feetTop"),
					d.Number("feetBottom"),
					d.Number("feetWidth"),
					d.Number("feetHeight"),
				},
			},
		}

		Inventory[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d Inventory Panel records", len(Inventory))
}
