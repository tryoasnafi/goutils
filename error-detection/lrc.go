package errordetection

// LRC return one byte sum which implement Block Character Check (BCC).
func LRC(data []byte) byte {
	res := data[0]
	for i := 1; i < len(data); i++ {
		res ^= data[i]
	}
	return res
}

// LRC1155 return one byte sum which implement LRC ISO 1155.
func LRC1155(data []byte) byte {
	res := data[0]
	for i := 1; i < len(data); i++ {
		res = (res + data[i]) & 0xFF
	}
	return ((res ^ 0xFF) + 1) & 0xFF
}

// Compare comparing original lrc with new recalculate data (BCC).
func CompareLRC(data []byte, lrc byte) bool {
	return LRC(data) == lrc
}

// Compare comparing original lrc with new recalculate data (LRC ISO 1155).
func CompareLRC1155(data []byte, lrc byte) bool {
	return LRC1155(data) == lrc
}