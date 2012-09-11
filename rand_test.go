package tools

import "testing"

func Testirand(t *testing.T) {

	if IRand(3, 5) > 5 || IRand(3, 5) < 3 {
		t.Fail()
	}
	println(IRand(3, 5))
}

func TestSRand(t *testing.T) {
	println(SRand(5, 5, true))
	println(SRand(5, 5, true))
	println(SRand(5, 5, true))
	println(SRand(5, 5, false))
	println(SRand(5, 7, true))
	println(SRand(5, 10, false))
	println(SRand(5, 50, true))
	println(SRand(5, 10, false))
	println(SRand(5, 50, true))
	println(SRand(5, 10, false))
	println(SRand(5, 50, true))
	println(SRand(5, 10, false))
	println(SRand(5, 50, true))
	println(SRand(5, 4, true))
	println(SRand(5, 400, true))
	println(SRand(6, 5, true))
	println(SRand(6, 5, true))
	println(SRand(6, 5, true))
	println(SRand(6, 5, true))
	println(SRand(6, 5, true))
	println(SRand(6, 5, true))
	println(SRand(6, 5, true))
	println(SRand(6, 5, true))
	println(SRand(6, 5, true))
	println(SRand(6, 5, true))
	println(SRand(6, 5, true))
	println(SRand(6, 5, true))

}
