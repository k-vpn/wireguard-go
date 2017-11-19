package main

import (
	"testing"
)

func TestCookieMAC1(t *testing.T) {

	// setup generator / checker

	var (
		generator CookieGenerator
		checker   CookieChecker
	)

	sk, err := newPrivateKey()
	if err != nil {
		t.Fatal(err)
	}
	pk := sk.publicKey()

	generator.Init(pk)
	checker.Init(pk)

	// check mac1

	src := []byte{192, 168, 13, 37, 10, 10, 10}

	checkMAC1 := func(msg []byte) {
		generator.AddMacs(msg)
		if !checker.CheckMAC1(msg) {
			t.Fatal("MAC1 generation/verification failed")
		}
		if checker.CheckMAC2(msg, src) {
			t.Fatal("MAC2 generation/verification failed")
		}
	}

	checkMAC1([]byte{
		0x99, 0xbb, 0xa5, 0xfc, 0x99, 0xaa, 0x83, 0xbd,
		0x7b, 0x00, 0xc5, 0x9a, 0x4c, 0xb9, 0xcf, 0x62,
		0x40, 0x23, 0xf3, 0x8e, 0xd8, 0xd0, 0x62, 0x64,
		0x5d, 0xb2, 0x80, 0x13, 0xda, 0xce, 0xc6, 0x91,
		0x61, 0xd6, 0x30, 0xf1, 0x32, 0xb3, 0xa2, 0xf4,
		0x7b, 0x43, 0xb5, 0xa7, 0xe2, 0xb1, 0xf5, 0x6c,
		0x74, 0x6b, 0xb0, 0xcd, 0x1f, 0x94, 0x86, 0x7b,
		0xc8, 0xfb, 0x92, 0xed, 0x54, 0x9b, 0x44, 0xf5,
		0xc8, 0x7d, 0xb7, 0x8e, 0xff, 0x49, 0xc4, 0xe8,
		0x39, 0x7c, 0x19, 0xe0, 0x60, 0x19, 0x51, 0xf8,
		0xe4, 0x8e, 0x02, 0xf1, 0x7f, 0x1d, 0xcc, 0x8e,
		0xb0, 0x07, 0xff, 0xf8, 0xaf, 0x7f, 0x66, 0x82,
		0x83, 0xcc, 0x7c, 0xfa, 0x80, 0xdb, 0x81, 0x53,
		0xad, 0xf7, 0xd8, 0x0c, 0x10, 0xe0, 0x20, 0xfd,
		0xe8, 0x0b, 0x3f, 0x90, 0x15, 0xcd, 0x93, 0xad,
		0x0b, 0xd5, 0x0c, 0xcc, 0x88, 0x56, 0xe4, 0x3f,
	})

	checkMAC1([]byte{
		0x33, 0xe7, 0x2a, 0x84, 0x9f, 0xff, 0x57, 0x6c,
		0x2d, 0xc3, 0x2d, 0xe1, 0xf5, 0x5c, 0x97, 0x56,
		0xb8, 0x93, 0xc2, 0x7d, 0xd4, 0x41, 0xdd, 0x7a,
		0x4a, 0x59, 0x3b, 0x50, 0xdd, 0x7a, 0x7a, 0x8c,
		0x9b, 0x96, 0xaf, 0x55, 0x3c, 0xeb, 0x6d, 0x0b,
		0x13, 0x0b, 0x97, 0x98, 0xb3, 0x40, 0xc3, 0xcc,
		0xb8, 0x57, 0x33, 0x45, 0x6e, 0x8b, 0x09, 0x2b,
		0x81, 0x2e, 0xd2, 0xb9, 0x66, 0x0b, 0x93, 0x05,
	})

	checkMAC1([]byte{
		0x9b, 0x96, 0xaf, 0x55, 0x3c, 0xeb, 0x6d, 0x0b,
		0x13, 0x0b, 0x97, 0x98, 0xb3, 0x40, 0xc3, 0xcc,
		0xb8, 0x57, 0x33, 0x45, 0x6e, 0x8b, 0x09, 0x2b,
		0x81, 0x2e, 0xd2, 0xb9, 0x66, 0x0b, 0x93, 0x05,
	})

	// exchange cookie reply

	func() {
		msg := []byte{
			0x6d, 0xd7, 0xc3, 0x2e, 0xb0, 0x76, 0xd8, 0xdf,
			0x30, 0x65, 0x7d, 0x62, 0x3e, 0xf8, 0x9a, 0xe8,
			0xe7, 0x3c, 0x64, 0xa3, 0x78, 0x48, 0xda, 0xf5,
			0x25, 0x61, 0x28, 0x53, 0x79, 0x32, 0x86, 0x9f,
			0xa0, 0x27, 0x95, 0x69, 0xb6, 0xba, 0xd0, 0xa2,
			0xf8, 0x68, 0xea, 0xa8, 0x62, 0xf2, 0xfd, 0x1b,
			0xe0, 0xb4, 0x80, 0xe5, 0x6b, 0x3a, 0x16, 0x9e,
			0x35, 0xf6, 0xa8, 0xf2, 0x4f, 0x9a, 0x7b, 0xe9,
			0x77, 0x0b, 0xc2, 0xb4, 0xed, 0xba, 0xf9, 0x22,
			0xc3, 0x03, 0x97, 0x42, 0x9f, 0x79, 0x74, 0x27,
			0xfe, 0xf9, 0x06, 0x6e, 0x97, 0x3a, 0xa6, 0x8f,
			0xc9, 0x57, 0x0a, 0x54, 0x4c, 0x64, 0x4a, 0xe2,
			0x4f, 0xa1, 0xce, 0x95, 0x9b, 0x23, 0xa9, 0x2b,
			0x85, 0x93, 0x42, 0xb0, 0xa5, 0x53, 0xed, 0xeb,
			0x63, 0x2a, 0xf1, 0x6d, 0x46, 0xcb, 0x2f, 0x61,
			0x8c, 0xe1, 0xe8, 0xfa, 0x67, 0x20, 0x80, 0x6d,
		}
		generator.AddMacs(msg)
		reply, err := checker.CreateReply(msg, 1377, src)
		if err != nil {
			t.Fatal("Failed to create cookie reply:", err)
		}
		if !generator.ConsumeReply(reply) {
			t.Fatal("Failed to consume cookie reply")
		}
	}()

	// check mac2

	checkMAC2 := func(msg []byte) {
		generator.AddMacs(msg)

		if !checker.CheckMAC1(msg) {
			t.Fatal("MAC1 generation/verification failed")
		}
		if !checker.CheckMAC2(msg, src) {
			t.Fatal("MAC2 generation/verification failed")
		}

		msg[5] ^= 0x20

		if checker.CheckMAC1(msg) {
			t.Fatal("MAC1 generation/verification failed")
		}
		if checker.CheckMAC2(msg, src) {
			t.Fatal("MAC2 generation/verification failed")
		}

		msg[5] ^= 0x20

		srcBad1 := []byte{192, 168, 13, 37, 40, 01}
		if checker.CheckMAC2(msg, srcBad1) {
			t.Fatal("MAC2 generation/verification failed")
		}

		srcBad2 := []byte{192, 168, 13, 38, 40, 01}
		if checker.CheckMAC2(msg, srcBad2) {
			t.Fatal("MAC2 generation/verification failed")
		}
	}

	checkMAC2([]byte{
		0x03, 0x31, 0xb9, 0x9e, 0xb0, 0x2a, 0x54, 0xa3,
		0xc1, 0x3f, 0xb4, 0x96, 0x16, 0xb9, 0x25, 0x15,
		0x3d, 0x3a, 0x82, 0xf9, 0x58, 0x36, 0x86, 0x3f,
		0x13, 0x2f, 0xfe, 0xb2, 0x53, 0x20, 0x8c, 0x3f,
		0xba, 0xeb, 0xfb, 0x4b, 0x1b, 0x22, 0x02, 0x69,
		0x2c, 0x90, 0xbc, 0xdc, 0xcf, 0xcf, 0x85, 0xeb,
		0x62, 0x66, 0x6f, 0xe8, 0xe1, 0xa6, 0xa8, 0x4c,
		0xa0, 0x04, 0x23, 0x15, 0x42, 0xac, 0xfa, 0x38,
	})

	checkMAC2([]byte{
		0x0e, 0x2f, 0x0e, 0xa9, 0x29, 0x03, 0xe1, 0xf3,
		0x24, 0x01, 0x75, 0xad, 0x16, 0xa5, 0x66, 0x85,
		0xca, 0x66, 0xe0, 0xbd, 0xc6, 0x34, 0xd8, 0x84,
		0x09, 0x9a, 0x58, 0x14, 0xfb, 0x05, 0xda, 0xf5,
		0x90, 0xf5, 0x0c, 0x4e, 0x22, 0x10, 0xc9, 0x85,
		0x0f, 0xe3, 0x77, 0x35, 0xe9, 0x6b, 0xc2, 0x55,
		0x32, 0x46, 0xae, 0x25, 0xe0, 0xe3, 0x37, 0x7a,
		0x4b, 0x71, 0xcc, 0xfc, 0x91, 0xdf, 0xd6, 0xca,
		0xfe, 0xee, 0xce, 0x3f, 0x77, 0xa2, 0xfd, 0x59,
		0x8e, 0x73, 0x0a, 0x8d, 0x5c, 0x24, 0x14, 0xca,
		0x38, 0x91, 0xb8, 0x2c, 0x8c, 0xa2, 0x65, 0x7b,
		0xbc, 0x49, 0xbc, 0xb5, 0x58, 0xfc, 0xe3, 0xd7,
		0x02, 0xcf, 0xf7, 0x4c, 0x60, 0x91, 0xed, 0x55,
		0xe9, 0xf9, 0xfe, 0xd1, 0x44, 0x2c, 0x75, 0xf2,
		0xb3, 0x5d, 0x7b, 0x27, 0x56, 0xc0, 0x48, 0x4f,
		0xb0, 0xba, 0xe4, 0x7d, 0xd0, 0xaa, 0xcd, 0x3d,
		0xe3, 0x50, 0xd2, 0xcf, 0xb9, 0xfa, 0x4b, 0x2d,
		0xc6, 0xdf, 0x3b, 0x32, 0x98, 0x45, 0xe6, 0x8f,
		0x1c, 0x5c, 0xa2, 0x20, 0x7d, 0x1c, 0x28, 0xc2,
		0xd4, 0xa1, 0xe0, 0x21, 0x52, 0x8f, 0x1c, 0xd0,
		0x62, 0x97, 0x48, 0xbb, 0xf4, 0xa9, 0xcb, 0x35,
		0xf2, 0x07, 0xd3, 0x50, 0xd8, 0xa9, 0xc5, 0x9a,
		0x0f, 0xbd, 0x37, 0xaf, 0xe1, 0x45, 0x19, 0xee,
		0x41, 0xf3, 0xf7, 0xe5, 0xe0, 0x30, 0x3f, 0xbe,
		0x3d, 0x39, 0x64, 0x00, 0x7a, 0x1a, 0x51, 0x5e,
		0xe1, 0x70, 0x0b, 0xb9, 0x77, 0x5a, 0xf0, 0xc4,
		0x8a, 0xa1, 0x3a, 0x77, 0x1a, 0xe0, 0xc2, 0x06,
		0x91, 0xd5, 0xe9, 0x1c, 0xd3, 0xfe, 0xab, 0x93,
		0x1a, 0x0a, 0x4c, 0xbb, 0xf0, 0xff, 0xdc, 0xaa,
		0x61, 0x73, 0xcb, 0x03, 0x4b, 0x71, 0x68, 0x64,
		0x3d, 0x82, 0x31, 0x41, 0xd7, 0x8b, 0x22, 0x7b,
		0x7d, 0xa1, 0xd5, 0x85, 0x6d, 0xf0, 0x1b, 0xaa,
	})
}
