package randomly

// RandPrintStrLen ASCII码表32 ~ 126为可打印字符
func RandPrintStrLen(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandIntGap(32, 126))
	}
	return string(bytes)
}

// RandNormStrLen 去除可能对SQL语句造成影响的字符
func RandNormStrLen(l int) string {
	bytes := make([]byte, l)
	candiChars := "!#$%&()*+,-.:;<=>@0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for i := 0; i < l; i++ {
		bytes[i] = candiChars[RandIntGap(0, len(candiChars)-1)]
	}
	return string(bytes)
}

// RandAlphabetStrLen 只考虑字母
func RandAlphabetStrLen(l int) string {
	bytes := make([]byte, l)
	candiChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for i := 0; i < l; i++ {
		bytes[i] = candiChars[RandIntGap(0, len(candiChars)-1)]
	}
	return string(bytes)
}

// RandHexStrLen 十六进制字符串
func RandHexStrLen(l int) string {
	bytes := make([]byte, l)
	candiChars := "0123456789ABCDE"
	for i := 0; i < l; i++ {
		bytes[i] = candiChars[RandIntGap(0, len(candiChars)-1)]
	}
	return string(bytes)
}
