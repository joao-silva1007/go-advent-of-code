package day1

import "testing"

func Test_part1(t *testing.T) {
	result := part1([]string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	})

	if result != 142 {
		t.Fatalf("result should be equal to 142 and not %d", result)
	}
}

func Test_part2(t *testing.T) {
	result := part2([]string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	})

	if result != 281 {
		t.Fatalf("result should be equal to 281 and not %d", result)
	}
}

func Test_run(t *testing.T) {
	run()
}
