package errordetection

// LRC return one byte longitudinal redundancy check sum
func LRC(data []byte) byte {
	res := data[0]
	for i := 1; i < len(data); i++ {
		res ^= data[i]
	}
	return res
}

// Compare comparing original lrc with new recalculate data
func Compare(data []byte, lrc byte) bool {
	return LRC(data) == lrc
}