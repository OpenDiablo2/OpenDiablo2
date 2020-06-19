package d2enum

import (
	"fmt"
)

type DescFuncID int

func Format1(value float64, string1 string) string {
	// +[value] [string1]
	return fmt.Sprintf("+%f %s", value, string1)
}

func Format2(value float64, string1 string) string {
	// [value]% [string1]
	return fmt.Sprintf("%f%% %s", value, string1)
}

func Format3(value float64, string1 string) string {
	// [value] [string1]
	return fmt.Sprintf("%f %s", value, string1)
}

func Format4(value float64, string1 string) string {
	// +[value]% [string1]
	return fmt.Sprintf("+%f%% %s", value, string1)
}

func Format5(value float64, string1 string) string {
	// [value*100/128]% [string1]
	return fmt.Sprintf("%f%% %s", (value*100.0)/128.0, string1)
}

func Format6(value float64, string1, string2 string) string {
	// +[value] [string1] [string2]
	return fmt.Sprintf("+%f %s %s", value, string1, string2)
}

func Format7(value float64, string1, string2 string) string {
	// [value]% [string1] [string2]
	return fmt.Sprintf("%f%% %s %s", value, string1, string2)
}

func Format8(value float64, string1, string2 string) string {
	// +[value]% [string1] [string2]
	return fmt.Sprintf("+%f%% %s %s", value, string1, string2)
}

func Format9(value float64, string1, string2 string) string {
	// [value] [string1] [string2]
	return fmt.Sprintf("%f %s %s", value, string1, string2)
}

func Format10(value float64, string1, string2 string) string {
	// [value*100/128]% [string1] [string2]
	return fmt.Sprintf("%f%% %s %s", (value*100.0)/128.0, string1, string2)
}

func Format11(value float64) string {
	// Repairs 1 Durability In [100 / value] Seconds
	return fmt.Sprintf("Repairs 1 Durability In %.0f Seconds", 100.0/value)
}

func Format12(value float64, string1 string) string {
	// +[value] [string1]
	return fmt.Sprintf("+%f %s", value, string1)
}

func Format13(value float64, class string) string {
	// +[value] to [class] Skill Levels
	return fmt.Sprintf("+%.0f to %s Skill Levels", value, class)
}

func Format14(value float64, skilltab, class string) string {
	// +[value] to [skilltab] Skill Levels ([class] Only)
	fmtStr := "+%.0f to %s Skill Levels (%s Only)"
	return fmt.Sprintf(fmtStr, value, skilltab, class)
}

func Format15(value float64, slvl int, skill, event string) string {
	// [value]% chance to cast [slvl] [skill] on [event]
	fmtStr := "%.0f%% chance to cast %d %s on %s"
	return fmt.Sprintf(fmtStr, value, slvl, skill, event)
}

func Format16(slvl int, skill string) string {
	// Level [sLvl] [skill] Aura When Equipped
	return fmt.Sprintf("Level %d %s Aura When Equipped", slvl, skill)
}

func Format17(value float64, string1 string, time int) string {
	// [value] [string1] (Increases near [time])
	return fmt.Sprintf("%f %s (Increases near %d)", value, string1, time)
}

func Format18(value float64, string1 string, time int) string {
	// [value]% [string1] (Increases near [time])
	return fmt.Sprintf("%f%% %s (Increases near %d)", value, string1, time)
}

func Format19(value float64, string1 string) string {
	// this is used by stats that use Blizzard's sprintf implementation
	// (if you don't know what that is, it won't be of interest to you
	// eitherway I guess), look at how prismatic is setup, the string is
	// the format that gets passed to their sprintf spinoff.
	return "" // TODO
}

func Format20(value float64, string1 string) string {
	// [value * -1]% [string1]
	return fmt.Sprintf("%f%% %s", value*-1.0, string1)
}

func Format21(value float64, string1 string) string {
	// [value * -1] [string1]
	return fmt.Sprintf("%f %s", value*-1.0, string1)
}

func Format22(value float64, string1, montype string) string {
	// [value]% [string1] [montype]
	return fmt.Sprintf("%f%% %s %s", value, string1, montype)
}

func Format23(value float64, string1 string) string {
	// (warning: this is bugged in vanilla and doesn't work properly
	// see CE forum)
	return "" // TODO
}

func Format24(value float64, string1, monster string) string {
	// [value]% [string1] [monster]
	return fmt.Sprintf("%f%% %s %s", value, string1, monster)
}

func Format25(slvl float64, skill string, charges, maxCharges int) string {
	// Level [slvl] [skill] ([charges]/[maxCharges] Charges)
	fmtStr := "Level %.0f %s (%d/%d Charges)"
	return fmt.Sprintf(fmtStr, slvl, skill, charges, maxCharges)
}

func Format26(value float64, string1 string) string {
	// not used by vanilla, present in the code but I didn't test it yet
	return "" // TODO
}

func Format27(value float64, string1 string) string {
	// not used by vanilla, present in the code but I didn't test it yet
	return "" // TODO
}

func Format28(value float64, skill, class string) string {
	// +[value] to [skill] ([class] Only)
	return fmt.Sprintf("+%f to %s (%s Only)", value, skill, class)
}

func Format29(value float64, skill string) string {
	// +[value] to [skill]
	return fmt.Sprintf("+%.0f to %s", value, skill)
}

func GetDescFunction(n DescFuncID) interface{} {
	m := map[DescFuncID]interface{}{
		DescFuncID(0):  Format1,
		DescFuncID(1):  Format2,
		DescFuncID(2):  Format3,
		DescFuncID(3):  Format4,
		DescFuncID(4):  Format5,
		DescFuncID(5):  Format6,
		DescFuncID(6):  Format7,
		DescFuncID(7):  Format8,
		DescFuncID(8):  Format9,
		DescFuncID(9):  Format10,
		DescFuncID(10): Format11,
		DescFuncID(11): Format12,
		DescFuncID(12): Format13,
		DescFuncID(13): Format14,
		DescFuncID(14): Format15,
		DescFuncID(15): Format16,
		DescFuncID(16): Format17,
		DescFuncID(17): Format18,
		DescFuncID(18): Format19,
		DescFuncID(19): Format20,
		DescFuncID(20): Format21,
		DescFuncID(21): Format22,
		DescFuncID(22): Format23,
		DescFuncID(23): Format24,
		DescFuncID(24): Format25,
		DescFuncID(25): Format26,
		DescFuncID(26): Format27,
		DescFuncID(27): Format28,
		DescFuncID(28): Format29,
	}
	return m[n]
}
