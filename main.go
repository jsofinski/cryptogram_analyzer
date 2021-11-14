package main
import (
    "fmt"
    "io/ioutil"
    "os"
    "log"  
    "strconv"
)

var inputArray [][]byte
var inputArrayString [][]string
var allKeys [][]string
var realKey []string
var rangeArray []int

var size = 4
var message []string

func main() {
    test()
    // testUTF8()
    // getFromUser()
}

func getFromUser() {
    parameters := os.Args[1:]

    // fmt.Println(len(parameters))
    readFileMessage(parameters[0])
    for i := 0; i < len(parameters); i++ {
        readFile(parameters[i], i)
    }
    getKey(len(parameters))

    result := xorBytes(message, realKey)
    for i := 0; i < len(result); i++ {
        if result[i] == "" {
            fmt.Print("?")
        } else {
            fmt.Print(stringByteToByte(result[i]))
        }
    } 
    fmt.Println()

}

func testUTF8() {
    readFileMessage("data/utf8_1.txt")
    readFile("data/utf8_1.txt", 0)
    readFile("data/utf8_2.txt", 1)
    readFile("data/utf8_3.txt", 2)
    readFile("data/utf8_4.txt", 3)
    readFile("data/utf8_5.txt", 4)

    getKey(5)
    
    result := xorBytes(message, realKey)
    for i := 0; i < len(result); i++ {
        if result[i] == "" {
            fmt.Print("?")
        } else {
            fmt.Print(stringByteToByte(result[i]))
        }
    }
    fmt.Println()

}
func test() {
    readFileMessage("data/message.txt")
    readFile("data/cryptogram1.txt", 0)
    readFile("data/cryptogram2.txt", 1)
    readFile("data/cryptogram3.txt", 2)
    readFile("data/cryptogram4.txt", 3)
    readFile("data/cryptogram5.txt", 4)
    readFile("data/cryptogram6.txt", 5)
    readFile("data/cryptogram7.txt", 6)
    readFile("data/cryptogram8.txt", 7)
    readFile("data/cryptogram9.txt", 8)
    readFile("data/cryptogram10.txt", 9)
    readFile("data/cryptogram11.txt", 10)
    readFile("data/cryptogram12.txt", 11)
    readFile("data/cryptogram13.txt", 12)
    readFile("data/cryptogram14.txt", 13)
    readFile("data/cryptogram15.txt", 14)
    readFile("data/cryptogram16.txt", 15)
    readFile("data/cryptogram17.txt", 16)
    readFile("data/cryptogram18.txt", 17)
    readFile("data/cryptogram19.txt", 18)
    readFile("data/cryptogram20.txt", 19)

    for i := 2; i <= 20; i++ {
        fmt.Print(strconv.Itoa(i) + ": ")
        getKey(i)
        result := xorBytes(message, realKey)
        for i := 0; i < len(result); i++ {
            if result[i] == "" {
                fmt.Print("?")
            } else {
                fmt.Print(stringByteToByte(result[i]))
            }
        }
        fmt.Println()
    }
    result := xorBytes(inputArrayString[11], realKey)
    for i := 0; i < len(result); i++ {
        if result[i] == "" {
            fmt.Print("?")
        } else {
            fmt.Print(stringByteToByte(result[i]))
        }
    }
    fmt.Println()
}

func getKey(size int) {
    // set key
    keyLength := 0;
    for i := 0; i < size; i++ {
        keyLength = max(keyLength, len(inputArrayString[i]))
    }
    allKeys = make([][]string, keyLength)
    rangeArray = make([]int, keyLength)

    // set array of number of cryptograms that have at least i length
    for i := 0; i < size; i++ {
        for j := 0; j < len(inputArrayString[i]); j++ {
            if j == 0 {
                rangeArray = append(rangeArray, 0)
            }
            rangeArray[j]++
        }
    }

    // init allKeys
    for i := 0; i < keyLength; i++ {
        for j := 0; j < size; j++ {
            allKeys[i] = append(allKeys[i], "")
        }
    }

    // iterate over all cryprograms
    for i := 0; i < size; i++ {
        tempKey := []string{}
        for k := 0; k < len(inputArrayString[i]); k++ {
            // special sumbols counter (both crypt1 and crypt2 will show that they have special at i, but only one of them has it)
            specialCount := 0
            // compare with other
            for j := 0; j < size; j++ {
                maxK := min(len(inputArrayString[i]), len(inputArrayString[j]))
                if k >= maxK {
                    continue
                }
                if i != j {
                    // if is it space
                    if checkIfSpecial(xorByteString(inputArrayString[i][k], inputArrayString[j][k])) {
                        specialCount++
                    }
                }
            }
            // if special occures at least with half of the cryptogram
            if specialCount >= (rangeArray[k])/2 {
                tempKey = append(tempKey, xorByteString(inputArrayString[i][k], "00100000"))
            } else {
                tempKey = append(tempKey, "")
            }
        }
        for k := 0; k < len(inputArrayString[i]); k++ {
            allKeys[k][i] = tempKey[k] 
        }
    }

    realKey = []string{}
    for i:=0; i < len(allKeys); i++ {
        realKey = append(realKey, mostFrequent(allKeys[i]))
    }

}

func stringByteToByte(bitString string) string {
    result, err := strconv.ParseUint(bitString, 2, 8)
    if err != nil {
        panic(err)
    }
    return string(result)
}

func stringToByteArray(text []byte) []byte {
    tempInputArray := []byte{}
    
    bitstring := ""
    for i := 0; i < len(string(text)); i++ {
        if string(text[i]) != " "{
            bitstring += string(text[i])
        } else {
            newByte, err := strconv.ParseUint(bitstring, 2, 8)
            if err != nil {
                panic(err)
            }
            tempInputArray = append(tempInputArray, byte(newByte))
            bitstring = ""
        }
    }
    return tempInputArray
}
func stringToStringArray(text []byte) []string {
    tempInputArrayString := []string{}
    
    bitstring := ""
    counter := 0
    for i := 0; i < len(string(text)); i++ {
        if string(text[i]) != " " && counter != 8 {
            bitstring += string(text[i])
            counter++
        } else {
            counter = 0
            tempInputArrayString = append(tempInputArrayString, bitstring)
            bitstring = ""
        }
    }
    return tempInputArrayString
}

func checkIfSpecial(text string) bool {
    if text[1] == '1' {
        return true
    } else {
        return false
    }
}
func readFile(fileName string, textPosition int) {
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()


    text, err := ioutil.ReadAll(file)
    
    
    inputArray = append(inputArray, stringToByteArray(text))
    inputArrayString = append(inputArrayString, stringToStringArray(text))
}

func readFileMessage(fileName string) {
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()


    text, err := ioutil.ReadAll(file)
    message = stringToStringArray(text)
    // fmt.Println(inputArray[textPosition])
    // fmt.Println(inputArrayString[textPosition])
}

func xorByteString(first string, second string) string {
    result := ""
    for i := 0; i < len(first); i++ {
        if first[i] == second[i] {
            result += "0"
        } else {
            result += "1"
        }
    }
    return result
}

func xorBytes(firstArray []string, secondArray[]string) []string {
    result := []string{}

    if len(firstArray) > len(secondArray) {
        x := len(secondArray)
        for i := 0; i < x; i++ {
            tempString := ""
            for j:= 0; j < len(secondArray[0]); j++ {
                if firstArray[i][j] == secondArray[i][j] {
                    tempString += "0"
                } else {
                    tempString += "1"
                }
            }
            result = append(result, tempString)
        } 
    } else {
        x := len(firstArray)
        for i := 0; i < x; i++ {
            tempString := ""
            for j:= 0; j < len(secondArray[i]); j++ {
                if firstArray[i][j] == secondArray[i][j] {
                    tempString += "0"
                } else {
                    tempString += "1"
                }
            }
            result = append(result, tempString)
        }
    }
    return result
}
func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
func min(a int, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}

func mostFrequent(array []string) string {
    // fmt.Println("start")
    m := make(map[string]int)
    for _, word := range array {
        _, ok := m[word]
        if !ok {
            m[word] = 1
        } else {
            m[word]++
        }
    }

    max := 0
    result := ""
    for key, value := range m {
        if value > max && key != "" {
            result = key
            max = value
        }
    }
    return result
}
