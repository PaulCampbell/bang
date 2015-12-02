package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestAddsBangToBangList(t *testing.T) {
	b := NewBangs("testFile.json")
	b.SetBang("breed", "Cocker Spaniel")
	if len(b.bangList) == 0 {
		panic("Item not added")
	}
	os.Remove("testFile.json")
}

func TestRetrievesBang(t *testing.T) {
	b := NewBangs("testFile.json")
	b.SetBang("breed", "Cocker Spaniel")
	if b.GetBang("breed") != "Cocker Spaniel" {
		panic("Item not retrieved")
	}
	os.Remove("testFile.json")
}

func TestListsBangs(t *testing.T) {
	b := NewBangs("testFile.json")
	b.SetBang("breed", "Cocker Spaniel")
	b.SetBang("colour", "Black")
	if len(b.ListBangs()) != 2 {
		println(b.ListBangs())
	}
	os.Remove("testFile.json")
}

func TestCreatesBangFileIfItDoesntExist(t *testing.T) {
	NewBangs("testFile.json")
	if _, err := os.Stat("testFile.json"); os.IsNotExist(err) {
		panic("failed to create bang file")
	}
	os.Remove("testFile.json")
}

func TestReadsBangsFromFile(t *testing.T) {
	if _, err := os.Stat("exisingBangFile.json"); os.IsNotExist(err) {
		w, err := os.Create("exisingBangFile.json")
		if err != nil {
			panic(err)
		}
		defer w.Close()
		w.Write([]byte(`{"some": "item"}`))
		b := NewBangs("exisingBangFile.json")
		if b.GetBang("some") != "item" {
			panic("existing file read failed")
		}
		os.Remove("exisingBangFile.json")
	}
}

func TestWritesBangToFile(t *testing.T) {
	b := NewBangs("testFile.json")
	b.SetBang("breed", "Cocker Spaniel")
	if len(b.bangList) == 0 {
		panic("Item not added")
	}
	encoded, _ := ioutil.ReadFile(b.bangFilePath)
	if string(encoded) != `{"breed":"Cocker Spaniel"}` {
		panic("Bang not written to file")
	}
	os.Remove("testFile.json")
}
