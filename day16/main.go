package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInt(s []string) int64 {
	n, _ := strconv.ParseInt(strings.Join(s, ""), 2, 0)

	return n
}

func parsePackets(packets []string, length int64) int64 {
	if len(packets) < 10 {
		return 0
	}

	v := packets[0:3]
	length += 3
	packets = packets[3:]
	t := packets[0:3]
	length += 3
	packets = packets[3:]

	pv := toInt(v)
	pt := toInt(t)

	//fmt.Println(pv)

	if pt == 4 {
		moreBits := true

		for moreBits {
			foo := packets[0:5]
			packets = packets[5:]
			length += 5

			if toInt(foo[0:1]) == 0 {
				moreBits = false
			}
		}
	} else {
		id := packets[0:1]
		packets = packets[1:]

		if toInt(id) == 0 {
			packets = packets[15:]

		} else {
			packets = packets[11:]
		}
	}

	pv += parsePackets(packets, length+pv)

	return pv
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var input string

	for scanner.Scan() {
		input = scanner.Text()
	}

	var bits string

	for _, x := range strings.Split(input, "") {
		foo, _ := strconv.ParseInt(x, 16, 0)
		bits += fmt.Sprintf("%04b", foo)
	}

	bitSlice := strings.Split(bits, "")

	fmt.Println("part 1:", parsePackets(bitSlice, 0))

}
