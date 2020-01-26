package d2video

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type BinkVideoMode uint32

const (
	BinkVideoModeNormal                   BinkVideoMode = 0
	BinkVideoModeHeightDoubled            BinkVideoMode = 1
	BinkVideoModeHeightInterlaced         BinkVideoMode = 2
	BinkVideoModeWidthDoubled             BinkVideoMode = 3
	BinkVideoModeWidthAndHeightDoubled    BinkVideoMode = 4
	BinkVideoModeWidthAndHeightInterlaced BinkVideoMode = 5
)

type BinkAudioAlgorithm uint32

const (
	BinkAudioAlgorithmFFT BinkAudioAlgorithm = 0
	BinkAudioAlgorithmDCT BinkAudioAlgorithm = 1
)

type BinkAudioTrack struct {
	AudioChannels     uint16
	AudioSampleRateHz uint16
	Stereo            bool
	Algorithm         BinkAudioAlgorithm
	AudioTrackId      uint32
}

type BinkDecoder struct {
	videoCodecRevision    byte
	fileSize              uint32
	numberOfFrames        uint32
	largestFrameSizeBytes uint32
	VideoWidth            uint32
	VideoHeight           uint32
	FPS                   uint32
	FrameTimeMS           uint32
	streamReader          *d2common.StreamReader
	VideoMode             BinkVideoMode
	HasAlphaPlane         bool
	Grayscale             bool
	AudioTracks           []BinkAudioTrack
	FrameIndexTable       []uint32 // Mask bit 0, as this is defined as a keyframe
	frameIndex            uint32
}

func CreateBinkDecoder(source []byte) *BinkDecoder {
	result := &BinkDecoder{
		streamReader: d2common.CreateStreamReader(source),
	}
	result.loadHeaderInformation()
	return result
}

func (v *BinkDecoder) GetNextFrame() {
	//v.streamReader.SetPosition(uint64(v.FrameIndexTable[i] & 0xFFFFFFFE))
	lengthOfAudioPackets := v.streamReader.GetUInt32() - 4
	samplesInPacket := v.streamReader.GetUInt32()
	v.streamReader.SkipBytes(int(lengthOfAudioPackets))
	log.Printf("Frame %d:\tSamp: %d", v.frameIndex, samplesInPacket)

	v.frameIndex++
}

func (v *BinkDecoder) loadHeaderInformation() {
	v.streamReader.SetPosition(0)
	headerBytes, _ := v.streamReader.ReadBytes(3)
	if string(headerBytes) != "BIK" {
		log.Fatal("Invalid header for bink video")
	}
	v.videoCodecRevision = v.streamReader.GetByte()
	v.fileSize = v.streamReader.GetUInt32()
	v.numberOfFrames = v.streamReader.GetUInt32()
	v.largestFrameSizeBytes = v.streamReader.GetUInt32()
	v.streamReader.SkipBytes(4) // Number of frames again?
	v.VideoWidth = v.streamReader.GetUInt32()
	v.VideoHeight = v.streamReader.GetUInt32()
	fpsDividend := v.streamReader.GetUInt32()
	fpsDivider := v.streamReader.GetUInt32()
	v.FPS = uint32(float32(fpsDividend) / float32(fpsDivider))
	v.FrameTimeMS = 1000 / v.FPS
	videoFlags := v.streamReader.GetUInt32()
	v.VideoMode = BinkVideoMode((videoFlags >> 28) & 0x0F)
	v.HasAlphaPlane = ((videoFlags >> 20) & 0x1) == 1
	v.Grayscale = ((videoFlags >> 17) & 0x1) == 1
	numberOfAudioTracks := v.streamReader.GetUInt32()
	v.AudioTracks = make([]BinkAudioTrack, numberOfAudioTracks)
	for i := 0; i < int(numberOfAudioTracks); i++ {
		v.streamReader.SkipBytes(2) // Unknown
		v.AudioTracks[i].AudioChannels = v.streamReader.GetUInt16()
	}
	for i := 0; i < int(numberOfAudioTracks); i++ {
		v.AudioTracks[i].AudioSampleRateHz = v.streamReader.GetUInt16()
		flags := v.streamReader.GetUInt16()
		v.AudioTracks[i].Stereo = ((flags >> 13) & 0x1) == 1
		v.AudioTracks[i].Algorithm = BinkAudioAlgorithm((flags >> 12) & 0x1)
	}
	for i := 0; i < int(numberOfAudioTracks); i++ {
		v.AudioTracks[i].AudioTrackId = v.streamReader.GetUInt32()
	}
	v.FrameIndexTable = make([]uint32, v.numberOfFrames+1)
	for i := 0; i < int(v.numberOfFrames+1); i++ {
		v.FrameIndexTable[i] = v.streamReader.GetUInt32()
	}
}
