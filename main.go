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
var key [][]string
var realKey []string
var size = 4
var task []string

func main() {
    fmt.Println("hello world")
    readFileTask("zad1.txt")
    readFile("cryptogram1.txt", 0)
    readFile("cryptogram2.txt", 1)
    readFile("cryptogram3.txt", 2)
    readFile("cryptogram4.txt", 3)
    readFile("cryptogram5.txt", 4)
    readFile("cryptogram6.txt", 5)
    readFile("cryptogram7.txt", 6)
    readFile("cryptogram8.txt", 7)
    readFile("cryptogram9.txt", 8)
    readFile("cryptogram10.txt", 9)
    readFile("cryptogram11.txt", 10)
    readFile("cryptogram12.txt", 11)
    readFile("cryptogram13.txt", 12)
    readFile("cryptogram14.txt", 13)
    readFile("cryptogram15.txt", 14)
    readFile("cryptogram16.txt", 15)
    readFile("cryptogram17.txt", 16)
    readFile("cryptogram18.txt", 17)
    readFile("cryptogram19.txt", 18)
    readFile("cryptogram20.txt", 19)
    // fmt.Println(xorBytes(inputArrayString[0], inputArrayString[1]))
    // fmt.Println((inputArrayString[0]))
    // fmt.Println(xorBytes(inputArrayString[3], inputArrayString[0])[9])
    // fmt.Println(xorBytes(inputArrayString[3], inputArrayString[1])[9])
    // fmt.Println(xorBytes(inputArrayString[3], inputArrayString[2])[9])
    // fmt.Println(xorBytes(inputArrayString[0], inputArrayString[1]))
    // fmt.Println(xorBytes(inputArrayString[1], inputArrayString[1])[3])
    // fmt.Println(xorByteString(inputArrayString[0][3], "01010011"))
    // fmt.Println(xorByteString(inputArrayString[1][3], "01010011"))
    getKey(20)
    fmt.Println(realKey)
    // cryptogram1 := xorBytes(inputArrayString[0], realKey)
    result := xorBytes(task, realKey)
    for i := 0; i < len(inputArrayString[0]); i++ {
        if result[i] == "" {
            fmt.Print("?")
        } else {
            fmt.Print(stringByteToByte(result[i]))
        }
    } 
    fmt.Println()
}

func getKey(size int) {
    keyLength := 0;
    for i := 0; i < size; i++ {
        keyLength = max(keyLength, len(inputArrayString[i]))
    }
    key := make([][]string, keyLength)
    for i := 0; i < keyLength; i++ {
        for j := 0; j < size; j++ {
            key[i] = append(key[i], "")
        }
    }

    for i := 0; i < size; i++ {
        tempKey := []string{}
        for k := 0; k < len(inputArrayString[i]); k++ {
            specialCount := 0
            for j := 0; j < size; j++ {
                maxK := min(len(inputArrayString[i]), len(inputArrayString[j]))
                if k >= maxK {
                    break
                }
                if i != j {
                    if checkIfSpecial(xorByteString(inputArrayString[i][k], inputArrayString[j][k])) {
                        specialCount++
                    }
                }
            }
            if specialCount >= size/2 {
                tempKey = append(tempKey, xorByteString(inputArrayString[i][k], "00100000"))
            } else {
                tempKey = append(tempKey, "")
            }
        }
        for k := 0; k < len(inputArrayString[i]); k++ {
            key[k][i] = tempKey[k] 
        }
    }

    // fmt.Println(key)
    realKey = []string{}
    for i:=0; i < len(key); i++ {
        realKey = append(realKey, mostFrequent(key[i]))
    }
    // fmt.Println(mostFrequent(key[3]))
    // fmt.Println()
    // fmt.Println()

    // fmt.Println(realKey)

}

func stringByteToByte(bitString string) string {
    result, err := strconv.ParseUint(bitString, 2, 8)
    if err != nil {
        panic(err)
    }
    // fmt.Println("cyk")
    // fmt.Println(string(result))
    return string(result)
}

func stringToByteArray(text []byte) []byte {
    tempInputArray := []byte{}
    
    bitstring := ""
    for i := 0; i < len(string(text)); i++ {
        if string(text[i]) != " " {
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
    for i := 0; i < len(string(text)); i++ {
        if string(text[i]) != " " {
            bitstring += string(text[i])
        } else {
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
    // fmt.Println(inputArray[textPosition])
    // fmt.Println(inputArrayString[textPosition])
}

func readFileTask(fileName string) {
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
    task = stringToStringArray(text)
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
