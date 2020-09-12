package d2video

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
)

// BinkVideoMode is the video mode type
type BinkVideoMode uint32

const (
	// BinkVideoModeNormal is a normal video
	BinkVideoModeNormal BinkVideoMode = iota

	// BinkVideoModeHeightDoubled is a height-doubled video
	BinkVideoModeHeightDoubled

	// BinkVideoModeHeightInterlaced is a height-interlaced video
	BinkVideoModeHeightInterlaced

	// BinkVideoModeWidthDoubled is a width-doubled video
	BinkVideoModeWidthDoubled

	// BinkVideoModeWidthAndHeightDoubled is a width and height-doubled video
	BinkVideoModeWidthAndHeightDoubled

	// BinkVideoModeWidthAndHeightInterlaced is a width and height interlaced video
	BinkVideoModeWidthAndHeightInterlaced
)

// BinkAudioAlgorithm represents the type of bink audio algorithm
type BinkAudioAlgorithm uint32

const (
	// BinkAudioAlgorithmFFT is the FTT audio algorithm
	BinkAudioAlgorithmFFT BinkAudioAlgorithm = iota

	// BinkAudioAlgorithmDCT is the DCT audio algorithm
	BinkAudioAlgorithmDCT
)

// BinkAudioTrack represents an audio track
type BinkAudioTrack struct {
	AudioChannels     uint16
	AudioSampleRateHz uint16
	Stereo            bool
	Algorithm         BinkAudioAlgorithm
	AudioTrackID      uint32
}

// BinkDecoder represents the bink decoder
type BinkDecoder struct {
	AudioTracks           []BinkAudioTrack
	FrameIndexTable       []uint32
	streamReader          *d2datautils.StreamReader
	fileSize              uint32
	numberOfFrames        uint32
	largestFrameSizeBytes uint32
	VideoWidth            uint32
	VideoHeight           uint32
	FPS                   uint32
	FrameTimeMS           uint32
	VideoMode             BinkVideoMode
	frameIndex            uint32
	videoCodecRevision    byte
	HasAlphaPlane         bool
	Grayscale             bool

	// Mask bit 0, as this is defined as a keyframe

}

// CreateBinkDecoder returns a new instance of the bink decoder
func CreateBinkDecoder(source []byte) *BinkDecoder {
	result := &BinkDecoder{
		streamReader: d2datautils.CreateStreamReader(source),
	}

	result.loadHeaderInformation()

	return result
}

// GetNextFrame gets the next frame
func (v *BinkDecoder) GetNextFrame() {
	//nolint:gocritic // v.streamReader.SetPosition(uint64(v.FrameIndexTable[i] & 0xFFFFFFFE))
	lengthOfAudioPackets := v.streamReader.GetUInt32() - 4 //nolint:gomnd // decode magic
	samplesInPacket := v.streamReader.GetUInt32()

	v.streamReader.SkipBytes(int(lengthOfAudioPackets))

	log.Printf("Frame %d:\tSamp: %d", v.frameIndex, samplesInPacket)

	v.frameIndex++
}

//nolint:gomnd // Decoder magic
func (v *BinkDecoder) loadHeaderInformation() {
	v.streamReader.SetPosition(0)
	headerBytes := v.streamReader.ReadBytes(3)

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
		v.AudioTracks[i].AudioTrackID = v.streamReader.GetUInt32()
	}

	v.FrameIndexTable = make([]uint32, v.numberOfFrames+1)

	for i := 0; i < int(v.numberOfFrames+1); i++ {
		v.FrameIndexTable[i] = v.streamReader.GetUInt32()
	}
}
