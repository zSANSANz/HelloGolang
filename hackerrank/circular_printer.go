package main

import (
	"fmt"
	"strings"
)

func main() {
	word := ""
	fmt.Scan(&word)
	words := strings.ToLower(word)
	posisi_0 := 0
	posisi_1 := 0
	posisi_2 := 0
	posisi_3 := 0
	
	temp_1 := ""
	temp_2 := ""
	temp_3 := ""
	
	for index, letter := range words {
		// fmt.Println(" index ke: ", index, "berisi huruf : ", string(letter))
		if index == 0 {
			if string(letter) == "a" {
				posisi_0 = 0
				// fmt.Println(posisi_0)
				temp_1 = "a"
			} else if string(letter) == "z" || string(letter) == "b" {
				posisi_0 = 1
				temp_1 = "z"
				fmt.Println(posisi_0)
			} else if string(letter) == "y" || string(letter) == "c" {
				posisi_0 = 2
				// fmt.Println(posisi_0)
			} else if string(letter) == "x" || string(letter) == "d" {
				posisi_0 = 3
				// fmt.Println(posisi_0)
			} else if string(letter) == "w" || string(letter) == "e" {
				posisi_0 = 4
				// fmt.Println(posisi_0)
			} else if string(letter) == "v" || string(letter) == "f" {
				posisi_0 = 5
				// fmt.Println(posisi_0)
			} else if string(letter) == "u" || string(letter) == "g" {
				posisi_0 = 6
				// fmt.Println(posisi_0)
			} else if string(letter) == "t" || string(letter) == "h" {
				posisi_0 = 7
				// fmt.Println(posisi_0)
			} else if string(letter) == "s" || string(letter) == "i" {
				posisi_0 = 8
				// fmt.Println(posisi_0)
			} else if string(letter) == "r" || string(letter) == "j" {
				posisi_0 = 9
				// fmt.Println(posisi_0)
			} else if string(letter) == "q" || string(letter) == "k" {
				posisi_0 = 10
				// fmt.Println(posisi_0)
			} else if string(letter) == "p" || string(letter) == "l" {
				posisi_0 = 11
				// fmt.Println(posisi_0)
			} else if string(letter) == "o" || string(letter) == "m" {
				posisi_0 = 12
				// fmt.Println(posisi_0)
			} else if string(letter) == "n" {
				posisi_0 = 13
				// fmt.Println(posisi_0)
			}
		}
		
		if temp_1 == "a" {
			if index == 1 {
				if string(letter) == "a" {
					posisi_1 = 0
					// fmt.Println(posisi_1)
				} else if string(letter) == "z" || string(letter) == "b" {
					posisi_1 = 1
					// fmt.Println(posisi_1)
					temp_2 = "z"
				} else if string(letter) == "y" || string(letter) == "c" {
					posisi_1 = 2
					// fmt.Println(posisi_1)
				} else if string(letter) == "x" || string(letter) == "d" {
					posisi_1 = 3
					// fmt.Println(posisi_1)
				} else if string(letter) == "w" || string(letter) == "e" {
					posisi_1 = 4
					// fmt.Println(posisi_1)
				} else if string(letter) == "v" || string(letter) == "f" {
					posisi_1 = 5
					// fmt.Println(posisi_1)
				} else if string(letter) == "u" || string(letter) == "g" {
					posisi_1 = 6
					// fmt.Println(posisi_1)
				} else if string(letter) == "t" || string(letter) == "h" {
					posisi_1 = 7
					// fmt.Println(posisi_1)
				} else if string(letter) == "s" || string(letter) == "i" {
					posisi_1 = 8
					// fmt.Println(posisi_1)
				} else if string(letter) == "r" || string(letter) == "j" {
					posisi_1 = 9
					// fmt.Println(posisi_1)
				} else if string(letter) == "q" || string(letter) == "k" {
					posisi_1 = 10
					// fmt.Println(posisi_1)
				} else if string(letter) == "p" || string(letter) == "l" {
					posisi_1 = 11
					// fmt.Println(posisi_1)
				} else if string(letter) == "o" || string(letter) == "m" {
					posisi_1 = 12
					// fmt.Println(posisi_1)
				} else if string(letter) == "n" {
					posisi_1 = 13
					// fmt.Println(posisi_1)
				}
			}
		}

		if temp_1 == "z" {
			if index == 2 {
				if string(letter) == "a" {
					posisi_2 = 0
					// fmt.Println(posisi_2)
				} else if string(letter) == "z" || string(letter) == "b" {
					posisi_2 = 1
					// fmt.Println(posisi_2)
				} else if string(letter) == "y" || string(letter) == "c" {
					posisi_2 = 2
					// fmt.Println(posisi_2)
				} else if string(letter) == "x" || string(letter) == "d" {
					posisi_2 = 3
					// fmt.Println(posisi_2)
				} else if string(letter) == "w" || string(letter) == "e" {
					posisi_2 = 4
					// fmt.Println(posisi_2)
				} else if string(letter) == "v" || string(letter) == "f" {
					posisi_2 = 5
					// fmt.Println(posisi_2)
				} else if string(letter) == "u" || string(letter) == "g" {
					posisi_2 = 7
					// fmt.Println(posisi_2)
					temp_3 = "g"
				} else if string(letter) == "t" || string(letter) == "h" {
					posisi_2 = 77
					// fmt.Println(posisi_2)
				} else if string(letter) == "s" || string(letter) == "i" {
					posisi_2 = 8
					// fmt.Println(posisi_2)
				} else if string(letter) == "r" || string(letter) == "j" {
					posisi_2 = 9
					// fmt.Println(posisi_2)
				} else if string(letter) == "q" || string(letter) == "k" {
					posisi_2 = 10
					// fmt.Println(posisi_2)
				} else if string(letter) == "p" || string(letter) == "l" {
					posisi_2 = 11
					// fmt.Println(posisi_2)
				} else if string(letter) == "o" || string(letter) == "m" {
					posisi_2 = 12
					// fmt.Println(posisi_2)
				} else if string(letter) == "n" {
					posisi_2 = 13
					fmt.Println(posisi_2)
					temp_2 = "n"
				}
			}
		}
		
		if temp_2 == "n" {
			if index == 2 {
				if string(letter) == "a" {
					posisi_2 = 0
					// fmt.Println(posisi_2)
				} else if string(letter) == "z" || string(letter) == "b" {
					posisi_2 = 1
					// fmt.Println(posisi_2)
				} else if string(letter) == "y" || string(letter) == "c" {
					posisi_2 = 2
					// fmt.Println(posisi_2)
				} else if string(letter) == "x" || string(letter) == "d" {
					posisi_2 = 3
					// fmt.Println(posisi_2)
				} else if string(letter) == "w" || string(letter) == "e" {
					posisi_2 = 4
					// fmt.Println(posisi_2)
				} else if string(letter) == "v" || string(letter) == "f" {
					posisi_2 = 5
					// fmt.Println(posisi_2)
				} else if string(letter) == "u" || string(letter) == "g" {
					posisi_2 = 7
					// fmt.Println(posisi_2)
					temp_3 = "g"
				} else if string(letter) == "t" || string(letter) == "h" {
					posisi_2 = 77
					// fmt.Println(posisi_2)
				} else if string(letter) == "s" || string(letter) == "i" {
					posisi_2 = 8
					// fmt.Println(posisi_2)
				} else if string(letter) == "r" || string(letter) == "j" {
					posisi_2 = 9
					// fmt.Println(posisi_2)
				} else if string(letter) == "q" || string(letter) == "k" {
					posisi_2 = 10
					// fmt.Println(posisi_2)
				} else if string(letter) == "p" || string(letter) == "l" {
					posisi_2 = 11
					// fmt.Println(posisi_2)
				} else if string(letter) == "o" || string(letter) == "m" {
					posisi_2 = 12
					fmt.Println(posisi_2)
					temp_2 = "m"
				} else if string(letter) == "n" {
					posisi_2 = 13
					// fmt.Println(posisi_2)
				}
			}
		}

		if temp_2 == "z" {
			if index == 2 {
				if string(letter) == "a" {
					posisi_2 = 0
					// fmt.Println(posisi_2)
				} else if string(letter) == "z" || string(letter) == "b" {
					posisi_2 = 1
					// fmt.Println(posisi_2)
				} else if string(letter) == "y" || string(letter) == "c" {
					posisi_2 = 2
					// fmt.Println(posisi_2)
				} else if string(letter) == "x" || string(letter) == "d" {
					posisi_2 = 3
					// fmt.Println(posisi_2)
				} else if string(letter) == "w" || string(letter) == "e" {
					posisi_2 = 4
					// fmt.Println(posisi_2)
				} else if string(letter) == "v" || string(letter) == "f" {
					posisi_2 = 5
					// fmt.Println(posisi_2)
				} else if string(letter) == "u" || string(letter) == "g" {
					posisi_2 = 7
					// fmt.Println(posisi_2)
					temp_3 = "g"
				} else if string(letter) == "t" || string(letter) == "h" {
					posisi_2 = 77
					// fmt.Println(posisi_2)
				} else if string(letter) == "s" || string(letter) == "i" {
					posisi_2 = 8
					// fmt.Println(posisi_2)
				} else if string(letter) == "r" || string(letter) == "j" {
					posisi_2 = 9
					// fmt.Println(posisi_2)
				} else if string(letter) == "q" || string(letter) == "k" {
					posisi_2 = 10
					// fmt.Println(posisi_2)
				} else if string(letter) == "p" || string(letter) == "l" {
					posisi_2 = 11
					// fmt.Println(posisi_2)
				} else if string(letter) == "o" || string(letter) == "m" {
					posisi_2 = 12
					// fmt.Println(posisi_2)
				} else if string(letter) == "n" {
					posisi_2 = 13
					// fmt.Println(posisi_2)
				}
			}
		}

		if temp_3 == "g" {
			if index == 3 {
				if string(letter) == "a" {
					posisi_3 = 0
					// fmt.Println(posisi_3)
				} else if string(letter) == "z" || string(letter) == "b" {
					posisi_3 = 5
					// fmt.Println(posisi_3)
				} else if string(letter) == "y" || string(letter) == "c" {
					posisi_3 = 2
					// fmt.Println(posisi_3)
				} else if string(letter) == "x" || string(letter) == "d" {
					posisi_3 = 3
					// fmt.Println(posisi_3)
				} else if string(letter) == "w" || string(letter) == "e" {
					posisi_3 = 4
					// fmt.Println(posisi_3)
				} else if string(letter) == "v" || string(letter) == "f" {
					posisi_3 = 5
					// fmt.Println(posisi_3)
				} else if string(letter) == "u" || string(letter) == "g" {
					posisi_3 = 7
					// fmt.Println(posisi_3)
					temp_3 = "g"
				} else if string(letter) == "t" || string(letter) == "h" {
					posisi_3 = 77
					// fmt.Println(posisi_3)
				} else if string(letter) == "s" || string(letter) == "i" {
					posisi_3 = 8
					// fmt.Println(posisi_3)
				} else if string(letter) == "r" || string(letter) == "j" {
					posisi_3 = 9
					// fmt.Println(posisi_3)
				} else if string(letter) == "q" || string(letter) == "k" {
					posisi_3 = 10
					// fmt.Println(posisi_3)
				} else if string(letter) == "p" || string(letter) == "l" {
					posisi_3 = 11
					// fmt.Println(posisi_3)
				} else if string(letter) == "o" || string(letter) == "m" {
					posisi_3 = 12
					// fmt.Println(posisi_3)
				} else if string(letter) == "n" {
					posisi_3 = 13
					// fmt.Println(posisi_3)
				}
			}
		}

		if temp_3 == "m" {
			if index == 3 {
				if string(letter) == "a" {
					posisi_3 = 0
					// fmt.Println(posisi_3)
				} else if string(letter) == "z" || string(letter) == "b" {
					posisi_3 = 5
					// fmt.Println(posisi_3)
				} else if string(letter) == "y" || string(letter) == "c" {
					posisi_3 = 2
					// fmt.Println(posisi_3)
				} else if string(letter) == "x" || string(letter) == "d" {
					posisi_3 = 9
					// fmt.Println(posisi_3)
				} else if string(letter) == "w" || string(letter) == "e" {
					posisi_3 = 4
					// fmt.Println(posisi_3)
				} else if string(letter) == "v" || string(letter) == "f" {
					posisi_3 = 5
					// fmt.Println(posisi_3)
				} else if string(letter) == "u" || string(letter) == "g" {
					posisi_3 = 7
					// fmt.Println(posisi_3)
					temp_3 = "g"
				} else if string(letter) == "t" || string(letter) == "h" {
					posisi_3 = 77
					// fmt.Println(posisi_3)
				} else if string(letter) == "s" || string(letter) == "i" {
					posisi_3 = 8
					// fmt.Println(posisi_3)
				} else if string(letter) == "r" || string(letter) == "j" {
					posisi_3 = 9
					// fmt.Println(posisi_3)
				} else if string(letter) == "q" || string(letter) == "k" {
					posisi_3 = 10
					// fmt.Println(posisi_3)
				} else if string(letter) == "p" || string(letter) == "l" {
					posisi_3 = 11
					// fmt.Println(posisi_3)
				} else if string(letter) == "o" || string(letter) == "m" {
					posisi_3 = 12
					// fmt.Println(posisi_3)
				} else if string(letter) == "n" {
					posisi_3 = 13
					// fmt.Println(posisi_3)
				}
			}
		}

		
	}
	
	tenaga := posisi_0 + posisi_1 + posisi_2 + posisi_3
	fmt.Println(tenaga)
}