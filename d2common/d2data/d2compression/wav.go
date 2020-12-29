package d2compression

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
)

// WavDecompress decompresses wav files
//nolint:gomnd // binary decode magic
func WavDecompress(data []byte, channelCount int) []byte { //nolint:funlen,gocognit,gocyclo // can't reduce
	Array1 := []int{0x2c, 0x2c}
	Array2 := make([]int, channelCount)

	var sLookup = []int{
		0x0007, 0x0008, 0x0009, 0x000A, 0x000B, 0x000C, 0x000D, 0x000E,
		0x0010, 0x0011, 0x0013, 0x0015, 0x0017, 0x0019, 0x001C, 0x001F,
		0x0022, 0x0025, 0x0029, 0x002D, 0x0032, 0x0037, 0x003C, 0x0042,
		0x0049, 0x0050, 0x0058, 0x0061, 0x006B, 0x0076, 0x0082, 0x008F,
		0x009D, 0x00AD, 0x00BE, 0x00D1, 0x00E6, 0x00FD, 0x0117, 0x0133,
		0x0151, 0x0173, 0x0198, 0x01C1, 0x01EE, 0x0220, 0x0256, 0x0292,
		0x02D4, 0x031C, 0x036C, 0x03C3, 0x0424, 0x048E, 0x0502, 0x0583,
		0x0610, 0x06AB, 0x0756, 0x0812, 0x08E0, 0x09C3, 0x0ABD, 0x0BD0,
		0x0CFF, 0x0E4C, 0x0FBA, 0x114C, 0x1307, 0x14EE, 0x1706, 0x1954,
		0x1BDC, 0x1EA5, 0x21B6, 0x2515, 0x28CA, 0x2CDF, 0x315B, 0x364B,
		0x3BB9, 0x41B2, 0x4844, 0x4F7E, 0x5771, 0x602F, 0x69CE, 0x7462,
		0x7FFF,
	}

	var sLookup2 = []int{
		-1, 0, -1, 4, -1, 2, -1, 6,
		-1, 1, -1, 5, -1, 3, -1, 7,
		-1, 1, -1, 5, -1, 3, -1, 7,
		-1, 2, -1, 4, -1, 6, -1, 8,
	}

	input := d2datautils.CreateStreamReader(data)
	output := d2datautils.CreateStreamWriter()

	input.GetByte()

	shift := input.GetByte()

	for i := 0; i < channelCount; i++ {
		temp := input.GetInt16()
		Array2[i] = int(temp)
		output.PushInt16(temp)
	}

	channel := channelCount - 1

	for input.GetPosition() < input.GetSize() {
		value := input.GetByte()

		if channelCount == 2 {
			channel = 1 - channel
		}

		if (value & 0x80) != 0 {
			switch value & 0x7f {
			case 0:
				if Array1[channel] != 0 {
					Array1[channel]--
				}

				output.PushInt16(int16(Array2[channel]))
			case 1:
				Array1[channel] += 8
				if Array1[channel] > 0x58 {
					Array1[channel] = 0x58
				}

				if channelCount == 2 {
					channel = 1 - channel
				}
			case 2:
			default:
				Array1[channel] -= 8
				if Array1[channel] < 0 {
					Array1[channel] = 0
				}

				if channelCount == 2 {
					channel = 1 - channel
				}
			}
		} else {
			temp1 := sLookup[Array1[channel]]
			temp2 := temp1 >> shift

			if (value & 1) != 0 {
				temp2 += temp1 >> 0
			}
			if (value & 2) != 0 {
				temp2 += temp1 >> 1
			}
			if (value & 4) != 0 {
				temp2 += temp1 >> 2
			}
			if (value & 8) != 0 {
				temp2 += temp1 >> 3
			}
			if (value & 0x10) != 0 {
				temp2 += temp1 >> 4
			}
			if (value & 0x20) != 0 {
				temp2 += temp1 >> 5
			}

			temp3 := Array2[channel]
			if (value & 0x40) != 0 {
				temp3 -= temp2
				if temp3 <= -32768 {
					temp3 = -32768
				}
			} else {
				temp3 += temp2
				if temp3 >= 32767 {
					temp3 = 32767
				}
			}
			Array2[channel] = temp3
			output.PushInt16(int16(temp3))
			Array1[channel] += sLookup2[value&0x1f]

			if Array1[channel] < 0 {
				Array1[channel] = 0
			} else if Array1[channel] > 0x58 {
				Array1[channel] = 0x58
			}
		}
	}

	return output.GetBytes()
}
