package main

import "testing"

func TestHelloWorldyp1( t *testing.T){

	result := HelloWorld("yp1")

	if result != ("Hi yp1"){
		// Fail jika ada kesalahan masih melanjutkan test sampai selesai
		t.Fail()
	}

}

func TestHelloWorldyp2( t *testing.T){

	result := HelloWorld("yp2")

	if result != ("Hi yp2"){
		// FailNow jika ada kesalahan masih akan memberhentikan test
		t.FailNow()
	}

}

func TestHelloWorldyp3( t *testing.T){

	result := HelloWorld("yp3")

	if result != ("Hi yp3"){
		// Error jika ada kesalahan masih melanjutkan test sampai selesai dan juga mengirim error message
		t.Error("result not match")
	}

}

func TestHelloWorldyp4( t *testing.T){

	result := HelloWorld("yp4")

	if result != ("Hi yp4"){
		// Error jika ada kesalahan masih melanjutkan test sampai selesai dan juga mengirim error message
		t.Error("result not match")
	}

}