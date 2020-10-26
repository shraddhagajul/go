package secret

//Handshake Given a decimal number, convert it to the appropriate sequence of events for a secret handshake.
  /*1 = wink
	10 = double blink
	100 = close your eyes
	1000 = jump
	10000 = Reverse the order of the operations in the secret handshake.*/
func Handshake(decimalNo uint) []string{
	secret := make([]string,0) 

	handshake := [5]string{"wink", "double blink", "close your eyes" , "jump"}
	binaryNo := Convert(decimalNo)
	for counter, high := range binaryNo{
			if counter != 4 && high == 1{
				secret = append(secret, handshake[counter])
			}

			if counter == 4 && high == 1{
			revSecret := make([]string,0)
			for i,j:=len(secret),0; i > 0 ;i--{
				revSecret = append(revSecret,secret[i-1])
				j++
			}
			return revSecret
			}
	}

	return secret
}
func Convert(decimalNo uint) []uint{
		binary := make([]uint, 0)
		dividend := decimalNo
	for {
		remainder := dividend%2
		dividend = dividend/2
		binary = append(binary, remainder) 
		if dividend == 0{
			break
		}
	}
	 return binary
}
