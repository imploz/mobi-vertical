package types

import "math"

const (
	CBNCXSingle byte = 15  // 0x0F
	CBNCXParent byte = 111 // 0x6F
	CBNCXChild  byte = 31  // 0x1F
	CBSkeleton  byte = 10  // 0x0A
)

const INDXHeaderLength = 192 // 0xC0

type INDXHeader struct {
	INDX             [4]byte
	HeaderLength     uint32
	Unknown1         [4]byte
	HeaderType       uint32
	IndexType        uint32
	IDXTStart        uint32
	IndexRecordCount uint32
	IndexEncoding    uint32
	IndexLanguage    uint32
	IndexEntryCount  uint32
	ORDTStart        uint32
	LIGTStart        uint32
	LIGTCount        uint32
	CNCXCount        uint32
	Unknown2         [124]byte
	TAGXOffset       uint32
	Unknown3         [8]byte
}

func NewINDXHeader(RecordCount uint32, EntryCount uint32) INDXHeader {
	return INDXHeader{
		INDX:             [4]byte{'I', 'N', 'D', 'X'},
		HeaderLength:     INDXHeaderLength,
		HeaderType:       0,
		IndexType:        0, // 0: normal, 2: inflection
		IDXTStart:        0,
		IndexRecordCount: RecordCount,
		IndexEncoding:    65001,
		IndexLanguage:    math.MaxUint32,
		IndexEntryCount:  EntryCount,
		ORDTStart:        0,
		LIGTStart:        0,
		LIGTCount:        0,
		CNCXCount:        0,
		TAGXOffset:       INDXHeaderLength,
	}
}

const TAGXHeaderLength = 12 // 0x0C

type TAGXHeader struct {
	TAGX             [4]byte
	HeaderLength     uint32
	ControlByteCount uint32
}

func NewTAGXHeader() TAGXHeader {
	return TAGXHeader{
		TAGX:             [4]byte{'T', 'A', 'G', 'X'},
		HeaderLength:     TAGXHeaderLength,
		ControlByteCount: 1,
	}
}

const TAGXSingleHeaderLength = 32 // 0x20

type TAGXSingleHeader struct {
	TAGX             [4]byte
	HeaderLength     uint32
	ControlByteCount uint32
	TagTable         [5]TAGXTag
}

func NewTAGXSingleHeader() TAGXSingleHeader {
	return TAGXSingleHeader{
		TAGX:             [4]byte{'T', 'A', 'G', 'X'},
		HeaderLength:     TAGXSingleHeaderLength,
		ControlByteCount: 1,
		TagTable: [5]TAGXTag{
			TAGXTagEntryPosition,
			TAGXTagEntryLength,
			TAGXTagEntryNameOffset,
			TAGXTagEntryDepthLevel,
			TAGXTagEnd,
		},
	}
}

const TAGXTagLength = 4 // 0x04

type TAGXTag uint32

const (
	TAGXTagEntryPosition       TAGXTag = 0x01010100
	TAGXTagEntryLength         TAGXTag = 0x02010200
	TAGXTagEntryNameOffset     TAGXTag = 0x03010400
	TAGXTagEntryDepthLevel     TAGXTag = 0x04010800
	TAGXTagEntryParent         TAGXTag = 0x15011000
	TAGXTagEntryChild1         TAGXTag = 0x16012000
	TAGXTagEntryChildN         TAGXTag = 0x17014000
	TAGXTagEntryPosFid         TAGXTag = 0x06028000
	TAGXTagSkeletonChunkCount  TAGXTag = 0x01010300
	TAGXTagSkeletonGeometry    TAGXTag = 0x06020C00
	TAGXTagChunkCNCXOffset     TAGXTag = 0x02010100
	TAGXTagChunkFileNumber     TAGXTag = 0x03010200
	TAGXTagChunkSequenceNumber TAGXTag = 0x04010400
	TAGXTagChunkGeometry       TAGXTag = 0x06020800
	TAGXTagGuideTitle          TAGXTag = 0x01010100
	TAGXTagGuidePosFid         TAGXTag = 0x06020200
	TAGXTagEnd                 TAGXTag = 0x00000001
)

type TAGXTagTable []TAGXTag

var TAGXTableNCXSingle = TAGXTagTable{
	TAGXTagEntryPosition,
	TAGXTagEntryLength,
	TAGXTagEntryNameOffset,
	TAGXTagEntryDepthLevel,
	TAGXTagEnd,
}

var TAGXTableSkeleton = TAGXTagTable{
	TAGXTagSkeletonChunkCount,
	TAGXTagSkeletonGeometry,
	TAGXTagEnd,
}

var TAGXTableChunk = TAGXTagTable{
	TAGXTagChunkCNCXOffset,
	TAGXTagChunkFileNumber,
	TAGXTagChunkSequenceNumber,
	TAGXTagChunkGeometry,
	TAGXTagEnd,
}

var TAGXTableGuide = TAGXTagTable{
	TAGXTagGuideTitle,
	TAGXTagGuidePosFid,
	TAGXTagEnd,
}

const IDXTSingleHeaderLength = 6 // 0x06

type IDXTSingleHeader struct {
	IDXT   [4]byte
	Offset uint16
}

func NewIDXTSingleHeader(Offset uint16) IDXTSingleHeader {
	return IDXTSingleHeader{
		IDXT:   [4]byte{'I', 'D', 'X', 'T'},
		Offset: Offset,
	}
}

const IDXTHeaderLength = 4 // 0x04

type IDXTHeader struct {
	IDXT [4]byte
}

func NewIDXTHeader() IDXTHeader {
	return IDXTHeader{
		IDXT: [4]byte{'I', 'D', 'X', 'T'},
	}
}
