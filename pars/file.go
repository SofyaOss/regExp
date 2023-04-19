package pars

import "regexp"

func Split(data []byte) [][]string {
	reg := regexp.MustCompile(`(\d+)([+-])(\d+)=`)
	return reg.FindAllStringSubmatch(string(data), -1)
}
