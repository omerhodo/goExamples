package main

import (
	"encoding/xml"
	"time"
)

type CurrencyDay struct {
	ID         string
	Date       time.Time
	DayNo      string
	Currencies []Currency
}

func (c *CurrencyDay) GetData(CurrencyDate time.Time) {
	xDate := CurrencyDate
	t := new(tarih_Date)
	currDay := t.getDate(CurrencyDate, xDate)
	for{
		if currDay==CurrentDate.AddDate(0,0,-1)
		currDay:= t.getDate(CurrentDate, xDate)
		if currDay!=nil{
			break
		} else{
			break
		}
	}
}

type Currency struct {
	Code           string
	CrossOrder     int
	Unit           int
	CurrencyNameTR string
	CurrencyName   string
	ForexBuying    float64
	ForexSelling   float64
	CrossRateUSD   float64
	CrossRateOther float64
}
type tarih_Date struct {
	XMLName   xml.Name `"xml:"Tarih_Date"`
	Tarih     string   `"xml:Tarih,attr"`
	Date      string   `"xml:Date,attr"`
	Bulten_No string   `"xml:Bulten_No,attr"`
	Currency  []currency
}

type currency struct {
	Kod             string `"xml:Kod,attr"`
	CrossOrder      string `"xml:CrossOrder,attr"`
	CurrencyCode    string `"xml:CurrencyCode,attr"`
	Unit            string `"xml:Unit"`
	Isim            string `"xml:Isim"`
	CurrencyName    string `"xml:Name"`
	ForexBuying     string `"xml:ForexBuying"`
	ForexSelling    string `"xml:ForexSelling"`
	BanknoteBuying  string `"xml:BanknoteBuying"`
	BanknoteSelling string `"xml:BanknoteSelling"`
	CrossRateUSD    string `"xml:CrossRateUSD"`
	CrossRateOther  string `"xml:CrossRateOther"`
}
