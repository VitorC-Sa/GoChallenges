package rot13decipher

func MySolution(msg string) string {
	var resp []rune
	for _, v := range msg {
		if v >= 'A' && v <= 'Z' {
			v = 'A' + (v-'A'+13)%26
		} else if v >= 'a' && v <= 'z' {
			v = 'a' + (v-'a'+13)%26
		}
		resp = append(resp, v)
	}
	return string(resp)
}
