package example

import (
	"fmt"
)

func GetInfoUsage() {
	info, err := PosAPI.GetInformation()
	if err != nil {
		panic(err)
	}
	fmt.Println(info.RegisterNo)
	fmt.Println(info.PosID)
}
