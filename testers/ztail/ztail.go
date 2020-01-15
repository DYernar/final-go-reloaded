package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var data []byte
	var length, nbr int64

	argsCount := 0
	for range os.Args {
		argsCount++
	}

	var filenames []string
	filesNum := 0
	nbrStr := ""
	znak := 0
	optIndex := -999
	numIndex := -999
	optNumIndex := -999
	nBytes := ""
	for index, arg := range os.Args {
		if arg == "--help" {
			fmt.Printf("Usage: ./ztail -c [+]NUM [FILE]...\n\n")
			fmt.Printf("Outputs the last NUM bytes; or use -c +NUM to output starting with byte NUM of each file.\n\n")
			fmt.Printf("NUM may have a multiplier suffix: b 512, kB 1000, K 1024, MB 1000*1000, M 1024*1024, GB 1000*1000*1000, G 1024*1024*1024, and so on for T, P, E, Z, Y.\n")
			fmt.Printf("Binary prefixes can be used, too: KiB=K, MiB=M, and so on.\n")
			os.Exit(0)
		}

		if arg == "-c" {
			optIndex = index
			optNumIndex = -999
			if optIndex == argsCount-1 {
				fmt.Printf("ztail: option requires an argument -- 'c'\nTry './ztail --help' for more information.\n")
				os.Exit(1)
			}
			nBytes = os.Args[index+1]
			if os.Args[index+1] >= "0" && os.Args[index+1] < ":" {
				continue
			} else if os.Args[index+1] >= "+0" && os.Args[index+1] < "+:" {
				continue
			} else if os.Args[index+1] >= "-0" && os.Args[index+1] < "-:" {
				continue
			} else if os.Args[index+1] == "KiB" || os.Args[index+1] == "MiB" || os.Args[index+1] == "GiB" ||
				os.Args[index+1] == "kiB" || os.Args[index+1] == "miB" ||
				os.Args[index+1] == "TiB" || os.Args[index+1] == "PiB" || os.Args[index+1] == "EiB" ||
				os.Args[index+1] == "ZiB" || os.Args[index+1] == "YiB" ||
				os.Args[index+1] == "K" || os.Args[index+1] == "M" || os.Args[index+1] == "G" ||
				os.Args[index+1] == "k" || os.Args[index+1] == "m" ||
				os.Args[index+1] == "T" || os.Args[index+1] == "P" || os.Args[index+1] == "E" ||
				os.Args[index+1] == "Z" || os.Args[index+1] == "Y" ||
				os.Args[index+1] == "kB" || os.Args[index+1] == "MB" || os.Args[index+1] == "GB" ||
				os.Args[index+1] == "KB" || os.Args[index+1] == "mB" ||
				os.Args[index+1] == "TB" || os.Args[index+1] == "PB" || os.Args[index+1] == "EB" ||
				os.Args[index+1] == "ZB" || os.Args[index+1] == "YB" {
				continue
			} else {
				if os.Args[index+1][0] == '-' {
					fmt.Printf("ztail: invalid number of bytes: ‘%v’\n", os.Args[index+1][1:])
				} else {
					fmt.Printf("ztail: invalid number of bytes: ‘%v’\n", os.Args[index+1])
				}
				os.Exit(1)
			}
		}
		if arg >= "-c0" && arg < "-c:" {
			optNumIndex = index
			optIndex = -999

			nbrStr = arg[2:]
			nBytes = arg[2:]
			continue
		}
		if arg >= "-c+0" && arg < "-c+:" {
			optNumIndex = index
			optIndex = -999
			znak = 1
			nbrStr = arg[3:]
			nBytes = arg[2:]
			continue
		}
		if arg >= "-c-0" && arg < "-c-:" {
			optNumIndex = index
			optIndex = -999

			nbrStr = arg[3:]
			nBytes = arg[2:]
			continue
		}

		if arg >= "0" && arg < ":" {
			numIndex = index
			
			nbrStr = arg
			continue
		}
		if arg >= "+0" && arg < "+:" {
			numIndex = index
			znak = 1
			nbrStr = arg[1:]
			continue
		}
		if arg >= "-0" && arg < "-:" {
			numIndex = index
			
			nbrStr = arg[1:]
			continue
		}
		if arg == "-cKiB" || arg == "-ckiB" || arg == "-cMiB" || arg == "-cmiB" || arg == "-cGiB" ||
			arg == "-cTiB" || arg == "-cPiB" || arg == "-cEiB" ||
			arg == "-cZiB" || arg == "-cYiB" ||
			arg == "-cK" || arg == "-ck" || arg == "-cM" || arg == "-cm" || arg == "-cG" ||
			arg == "-cT" || arg == "-cP" || arg == "-cE" ||
			arg == "-cZ" || arg == "-cY" {
			optNumIndex = index
			nbrStr = "1" + arg[2:]
			nBytes = arg[2:]
			continue
		}
		if arg == "-c-KiB" || arg == "-c-kiB" || arg == "-c-MiB" || arg == "-c-miB" || arg == "-c-GiB" ||
			arg == "-c-TiB" || arg == "-c-PiB" || arg == "-c-EiB" ||
			arg == "-c-ZiB" || arg == "-c-YiB" ||
			arg == "-c-K" || arg == "-c-k" || arg == "-c-M" || arg == "-c-m" || arg == "-c-G" ||
			arg == "-c-T" || arg == "-c-P" || arg == "-c-E" ||
			arg == "-c-Z" || arg == "-c-Y" ||
			arg == "-c-kB" || arg == "-c-KB" || arg == "-c-MB" || arg == "-c-mB" || arg == "-c-GB" ||
			arg == "-c-TB" || arg == "-c-PB" || arg == "-c-EB" ||
			arg == "-c-ZB" || arg == "-c-YB" {
			optNumIndex = index
			nbrStr = "1" + arg[3:]
			nBytes = arg[2:]
			continue
		}
		if arg == "-c+KiB" || arg == "-c+kiB" || arg == "-c+MiB" || arg == "-c+miB" || arg == "-c+GiB" ||
			arg == "-c+TiB" || arg == "-c+PiB" || arg == "-c+EiB" ||
			arg == "-c+ZiB" || arg == "-c+YiB" ||
			arg == "-c+K" || arg == "-c+k" || arg == "-c+M" || arg == "-c+m" || arg == "-c+G" ||
			arg == "-c+T" || arg == "-c+P" || arg == "-c+E" ||
			arg == "-c+Z" || arg == "-c+Y" ||
			arg == "-c+kB" || arg == "-c+KB" || arg == "-c+MB" || arg == "-c+mB" || arg == "-c+GB" ||
			arg == "-c+TB" || arg == "-c+PB" || arg == "-c+EB" ||
			arg == "-c+ZB" || arg == "-c+YB" {
			fmt.Printf("ztail: invalid number of bytes: ‘%v’\n", arg[2:])
			os.Exit(1)
		}
		if arg == "KiB" || arg == "kiB" || arg == "MiB" || arg == "miB" || arg == "GiB" ||
			arg == "TiB" || arg == "PiB" || arg == "EiB" ||
			arg == "ZiB" || arg == "YiB" ||
			arg == "K" || arg == "k" || arg == "M" || arg == "m" || arg == "G" ||
			arg == "T" || arg == "P" || arg == "E" ||
			arg == "Z" || arg == "Y" ||
			arg == "kB" || arg == "KB" || arg == "mB" || arg == "MB" || arg == "GB" ||
			arg == "TB" || arg == "PB" || arg == "EB" ||
			arg == "ZB" || arg == "YB" {
			numIndex = index
			nbrStr = "1" + arg
			nBytes = arg
			continue
		}

		if index != 0 {
			filenames = append(filenames, arg)
			filesNum++
		}
	}
	
	if optIndex == -999 && numIndex == -999 {
		if optNumIndex == -999 {
			fmt.Printf("ztail: program requires an option\n")
			os.Exit(1)
		}
	}
	if optNumIndex == -999 {
		if numIndex-optIndex != 1 {
			if argsCount-1 > optIndex {
				fmt.Printf("ztail: invalid number of bytes: ‘%v’\n", os.Args[optIndex+1])
				os.Exit(1)
			} else {
				fmt.Printf("ztail: option requires an argument -- 'c'\n")
				os.Exit(1)
			}
		}
	}

	nbr = findNbr(nbrStr, nBytes)
	if nbr < 0 {
		fmt.Printf("ztail: invalid number of bytes: ‘%v’: Value too large for defined data type\n", nBytes)
		os.Exit(1)
	}

	for index, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("ztail: cannot open '%v' for reading: No such file or directory\n", filename)
			continue
		}
		if index != 0 {
			z01.PrintRune('\n')
		}
		stats, _ := file.Stat()
		length = stats.Size()
		var i int64
		for i = 0; i < length; i++ {
			data = append(data, 0)
		}
		file.Read(data)
		file.Close()
		if filesNum > 1 {
			fmt.Printf("==> %v <==\n", filename)
		}
		if nbr > 999999999999999 && znak == 1 {
			fmt.Printf("ztail: %v: cannot seek to relative offset %v: Invalid argument\n", filename, nbr-1)
			os.Exit(1)
		}
		printC(nbr, length, znak, data)
	}
	os.Exit(0)
}

func printC(nbr, length int64, znak int, data []byte) {
	var startFrom int64
	if znak == 1 {
		startFrom = nbr - 1
		if startFrom < 0 {
			startFrom = 0
		}
	} else {
		startFrom = length - nbr
		if startFrom < 0 {
			startFrom = 0
		}
	}
	for i := startFrom; i < length; i++ {
		fmt.Printf("%c", data[i])
	}
}

func findNbr(s, nbrStr string) int64 {
	var nbr int64
	lnNbr := 0
	for range s {
		lnNbr++
	}
	
	if s[lnNbr-1] == 'b' || s[lnNbr-1] == 'K' || s[lnNbr-1] == 'k' || s[lnNbr-1] == 'M' || s[lnNbr-1] == 'm' || s[lnNbr-1] == 'G' ||
		s[lnNbr-1] == 'T' || s[lnNbr-1] == 'P' || s[lnNbr-1] == 'E' || s[lnNbr-1] == 'Z' || s[lnNbr-1] == 'Y' {
		nbr = atoi(s[:lnNbr-1])
		if nbr < 0 {
			fmt.Printf("ztail: invalid number of bytes: ‘%v’\n", nbrStr)
			os.Exit(1)
		}
		nbr = nbr * multiSuffix(s[lnNbr-1:])
		
	} else if s[lnNbr-1] == 'B' && (s[lnNbr-2] == 'k' || s[lnNbr-2] == 'K' || s[lnNbr-2] == 'M' || s[lnNbr-2] == 'm' || s[lnNbr-2] == 'G' ||
		s[lnNbr-2] == 'T' || s[lnNbr-2] == 'P' || s[lnNbr-2] == 'E' || s[lnNbr-2] == 'Z' || s[lnNbr-2] == 'Y') {
		nbr = atoi(s[:lnNbr-2])
		if nbr < 0 {
			fmt.Printf("ztail: invalid number of bytes: ‘%v’\n", nbrStr)
			os.Exit(1)
		}
		nbr = nbr * multiSuffix(s[lnNbr-2:])
	} else if lnNbr >= 3 && (s[lnNbr-3:] == "KiB" || s[lnNbr-3:] == "kiB" || s[lnNbr-3:] == "MiB" || s[lnNbr-3:] == "miB" || s[lnNbr-3:] == "GiB" ||
		s[lnNbr-3:] == "TiB" || s[lnNbr-3:] == "PiB" || s[lnNbr-3:] == "EiB" || s[lnNbr-3:] == "ZiB" || s[lnNbr-3:] == "YiB") {
		nbr = atoi(s[:lnNbr-3])
		if nbr < 0 {
			fmt.Printf("ztail: invalid number of bytes: ‘%v’\n", nbrStr)
			os.Exit(1)
		}
		nbr = nbr * multiSuffix(s[lnNbr-3:])
	} else {
		nbr = atoi(s)
		if nbr < 0 {
			fmt.Printf("ztail: invalid number of bytes: ‘%v’\n", nbrStr)
			os.Exit(1)
		}
	}
	return nbr
}

func multiSuffix(s string) int64 {
	if s == "b" {
		return 512
	}
	if s == "kB" || s == "KB" {
		return 1000
	}
	if s == "K" || s == "k" || s == "KiB" || s == "kiB" {
		return 1024
	}
	if s == "MB" || s == "mB" {
		return 1000 * 1000
	}
	if s == "M" || s == "m" || s == "MiB" || s == "miB" {
		return 1024 * 1024
	}
	if s == "GB" {
		return 1000 * 1000 * 1000
	}
	if s == "G" || s == "GiB" {
		return 1024 * 1024 * 1024
	}
	if s == "TB" {
		return 1000 * 1000 * 1000 * 1000
	}
	if s == "T" || s == "TiB" {
		return 1024 * 1024 * 1024 * 1024
	}
	if s == "PB" {
		return 1000 * 1000 * 1000 * 1000 * 1000
	}
	if s == "P" || s == "PiB" {
		return 1024 * 1024 * 1024 * 1024 * 1024
	}
	if s == "EB" {
		return 1000 * 1000 * 1000 * 1000 * 1000 * 1000
	}
	if s == "E" || s == "EiB" {
		return 1024 * 1024 * 1024 * 1024 * 1024 * 1024
	}
	return -1
}

func atoi(s string) int64 {
	slice := []rune(s)
	var nbr int64
	length := 0
	for _, char := range slice {
		if char < '0' || char > '9' {
			return -999
		} else {
			var c int64
			for a := '0'; a < char; a++ {
				c++
			}
			if length == 18 {
				if nbr > 922337203685477580 {
					return -999
				}
				if nbr == 922337203685477580 && c > 7 {
					return -999
				}
			}
			if length == 19 {
				return -999
			}
			nbr = nbr*10 + c
			length++
		}
	}
	return nbr
}