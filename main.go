package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var nums []int
	for _, i := range strings.Split(scanner.Text(), " ") {
		a, _ := strconv.Atoi(i)
		nums = append(nums, a)
	}
	sort.Ints(nums)
	fmt.Println(nums[1])
}
