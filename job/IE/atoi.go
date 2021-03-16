package main

import (
	"errors"
	"fmt"
)

func main() {
	str := "-+aaaaaaaaa啊aaaaaaaaaaaaaaaa" //please input the string here
	base := 16                            //please input the base number here
	res, err := atoi(str, base)
	if err != nil {
		fmt.Println(err)
		fmt.Println(res)
		return
	}
	fmt.Println(res)
}

func atoi(str string, base int) (int, error) {
	if len(str) == 0 {
		return 0, errors.New("empty string")
	}
	if base < 2 || base > 16 {
		return 0, errors.New("invalid base")
	}
	validNumMap := map[byte]int{
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'a': 10,
		'A': 10,
		'b': 11,
		'B': 11,
		'c': 12,
		'C': 12,
		'd': 13,
		'D': 13,
		'e': 14,
		'E': 14,
		'f': 15,
		'F': 15,
	}
	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX
	strByt := []byte(str)
	isNegative := false
	bit := -1
	sum := 0
	for k, v := range strByt {
		if v != ' ' {
			bit = k
			break
		}
	}
	if bit == -1 {
		return 0, errors.New("empty string")
	}
	if strByt[bit] == '-' {
		isNegative = true
	} else if strByt[bit] == '+' {
		isNegative = false
	} else if (validNumMap[strByt[bit]] != 0 || strByt[bit] == '0') && validNumMap[strByt[bit]] < base {
		sum = validNumMap[strByt[bit]]
	} else {
		invalidChar := string(str[bit])
		if str[bit] == ' ' {
			invalidChar = "space"
		}
		return 0, errors.New("invalid char: " + invalidChar)
	}

	bit++

	for ; bit < len(str); bit++ {
		num := validNumMap[strByt[bit]]
		if (num == 0 && strByt[bit] != '0') || num >= base {
			invalidChar := string(str[bit])
			if str[bit] == ' ' {
				invalidChar = "a space"
			}
			return 0, errors.New("invalid char " + invalidChar)
		}
		//溢出检查
		if !isNegative && sum > (INT_MAX-num)/base {
			return INT_MAX, errors.New("input overflow int64")
		}
		if isNegative && (-1)*sum < (INT_MIN+num)/base {
			return INT_MIN, errors.New("input overflow int64")
		}
		sum = sum*base + num
	}

	if isNegative {
		return -1 * sum, nil
	}
	return sum, nil
}
