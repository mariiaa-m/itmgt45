package main

import "fmt"

// Shift letter
func shiftLetter(letter string, shift int) string {
	if letter == " " {
		return " "
	}
	char := int(letter[0])
	newChar := ((char-'A')+shift)%26 + 'A'
	return string(newChar)
}

// Caesar cipher
func caesarCipher(message string, shift int) string {
	result := ""
	for _, ch := range message {
		result += shiftLetter(string(ch), shift)
	}
	return result
}

// Shift by letter
func shiftByLetter(letter string, letterShift string) string {
	if letter == " " {
		return " "
	}
	shiftValue := int(letterShift[0] - 'A')
	return shiftLetter(letter, shiftValue)
}

// Vigenere cipher
func vigenereCipher(message string, key string) string {
	result := ""
	keyLength := len(key)
	keyIndex := 0
	for _, ch := range message {
		if ch == ' ' {
			result += " "
			continue
		}
		result += shiftByLetter(string(ch), string(key[keyIndex%keyLength]))
		keyIndex++
	}
	return result
}

// Scytale cipher
func scytaleCipher(message string, shift int) string {
	length := len(message)
	for length%shift != 0 {
		message += "_"
		length++
	}

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		newIndex := (i / shift) + (length/shift)*(i%shift)
		result[newIndex] = message[i]
	}
	return string(result)
}

// Scytale decipher
func scytaleDecipher(message string, shift int) string {
	length := len(message)
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		originalIndex := (i % shift) * (length / shift) + (i / shift)
		result[i] = message[originalIndex]
	}

	return string(result)
}

func main() {
	fmt.Println(shiftLetter("Z", 1))          
	fmt.Println(shiftLetter("M", 13))         
	fmt.Println(shiftLetter("X", 3))        
	fmt.Println(caesarCipher("XYZ", 3))       
	fmt.Println(caesarCipher("HELLO WORLD", 26)) 
	fmt.Println(shiftByLetter("C", "C"))       
	fmt.Println(shiftByLetter("X", "Y"))      
	fmt.Println(vigenereCipher("HELLO WORLD", "KEY")) 
	fmt.Println(vigenereCipher("XYZ", "ABC"))  
	fmt.Println(scytaleDecipher("IMNNAF_TAOIGROE", 3)) 
}
