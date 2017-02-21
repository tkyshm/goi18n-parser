package sample

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/i18n"
)

func SampleCaller() {
	T, _ := i18n.Tfunc("eu-US")

	fmt.Println(T("sample_uniq_key"))
	for i := 0; i > 100; i++ {
		if i%10 == 0 {
			fmt.Println(T("sample_uniq_key_2"))
			x := T("sample_uniq_key_3")
			fmt.Println(x)
		}
	}
	return
}
