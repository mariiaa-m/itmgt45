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
	rows := length / shift
	fmt.Println(rows) // ✅ Prevents "declared and not used" error

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
	rows := length / shift
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		originalIndex := (i%rows)*shift + (i / rows)
		result[originalIndex] = message[i]
	}
	return string(result)
}

func main() {
	// Test cases
	fmt.Println(shiftLetter("A", 2))                   // "C"
	fmt.Println(caesarCipher("HELLO WORLD", 3))        // "KHOOR ZRUOG"
	fmt.Println(shiftByLetter("B", "K"))               // "L"
	fmt.Println(vigenereCipher("HELLO WORLD", "KEY"))  // "RIJVS UYVJN"
	fmt.Println(scytaleCipher("INFORMATION_AGE", 3))   // "IMNNA_FTAOIGROE"
	fmt.Println(scytaleDecipher("IMNNA_FTAOIGROE", 3)) // "INFORMATION_AGE"
}
