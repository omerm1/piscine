package piscine

func Itoa(n int) string {
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	length := 1
	temp := n
	for temp >= 10 {
		temp /= 10
		length++
	}

	result := make([]byte, length+len(sign))
	
	idx := len(result) - 1
	for n >= 10 {
		result[idx] = byte(n%10) + '0'
		n /= 10
		idx--
	}
	result[idx] = byte(n) + '0'

	if sign != "" {
		result[idx-1] = sign[0]
	}

	return string(result)
}
