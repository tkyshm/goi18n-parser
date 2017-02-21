package sample

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/i18n"
)

type A struct {
	T i18n.TranslateFunc
}

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

	T2, _ := i18n.Tfunc("en-US")
	a := A{T: T2}

	fmt.Println(a.T("sample_uniq_key_4"))

	fmt.Printf("%s,%s,%s", a.T("sample_uniq_key_4"), T("sample_uniq_key_5"), T("sample_uniq_key_2"))
	return
}
