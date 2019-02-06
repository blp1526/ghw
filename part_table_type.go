package ghw

import "strings"

type PartTableType int

const (
	PART_TABLE_TYPE_UNKNOWN PartTableType = iota
	PART_TABLE_TYPE_NONE
	PART_TABLE_TYPE_MBR
	PART_TABLE_TYPE_GPT
)

var (
	partTableTypeString = map[PartTableType]string{
		PART_TABLE_TYPE_UNKNOWN: "Unknown",
		PART_TABLE_TYPE_NONE:    "None",
		PART_TABLE_TYPE_MBR:     "MBR",
		PART_TABLE_TYPE_GPT:     "GPT",
	}
)

func (ptt PartTableType) String() string {
	return partTableTypeString[ptt]
}

func (ptt PartTableType) MarshalJSON() ([]byte, error) {
	return []byte("\"" + strings.ToLower(ptt.String()) + "\""), nil
}
