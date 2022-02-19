package main

import "testing"

func TestCase1Part(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("王强", 0.32)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 1 {
			t.Fatalf("预期王强第一,但是得到的是:%d", randOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂是:0.32,但是得到的是:%f", fatRateOfWQ)
		}
	}
}

func TestCase1(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("王强", 0.32)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 1 {
			t.Fatalf("预期王强第一,但是得到的是:%d", randOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂是:0.32,但是得到的是:%f", fatRateOfWQ)
		}
	}
	inputRecord("李静", 0.28)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期王强第二,但是得到的是:%d", randOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂是:0.32,但是得到的是:%f", fatRateOfWQ)
		}
	}
	{
		randOfLG, fatRateOflj := getRand("李静")
		if randOfLG != 1 {
			t.Fatalf("预期王强第一,但是得到的是:%d", randOfLG)
		}
		if fatRateOflj != 0.28 {
			t.Fatalf("预期王强的体脂是:0.32,但是得到的是:%f", fatRateOflj)
		}
	}

}

func TestCase2(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("张伟", 0.38)
	inputRecord("李静", 0.28)
	{
		randOfLG, fatRateOflj := getRand("李静")
		if randOfLG != 1 {
			t.Fatalf("预期李静第一,但是得到的是:%d", randOfLG)
		}
		if fatRateOflj != 0.28 {
			t.Fatalf("预期李静的体脂是:0.28,但是得到的是:%f", fatRateOflj)
		}
	}
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期王强第二,但是得到的是:%d", randOfWQ)
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("预期王强的体脂是:0.38,但是得到的是:%f", fatRateOfWQ)
		}
	}
	{
		randOfZW, fatRateOfZW := getRand("张伟")
		if randOfZW != 2 {
			t.Fatalf("预期张伟第二,但是得到的是:%d", randOfZW)
		}
		if fatRateOfZW != 0.38 {
			t.Fatalf("预期张伟的体脂是:0.38,但是得到的是:%f", fatRateOfZW)
		}
	}

}

func TestCase3(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("李静", 0.28)
	inputRecord("张伟")

	{
		randOfLG, fatRateOflj := getRand("李静")
		if randOfLG != 1 {
			t.Fatalf("预期李静第一,但是得到的是:%d", randOfLG)
		}
		if fatRateOflj != 0.28 {
			t.Fatalf("预期李静的体脂是:0.28,但是得到的是:%f", fatRateOflj)
		}
	}
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期王强第二,但是得到的是:%d", randOfWQ)
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("预期王强的体脂是:0.38,但是得到的是:%f", fatRateOfWQ)
		}
	}
	{
		randOfZW, _ := getRand("张伟")
		if randOfZW != 3 {
			t.Fatalf("预期张伟第二,但是得到的是:%d", randOfZW)
		}

	}

}
