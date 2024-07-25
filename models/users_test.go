package models

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	info := LOGINUSER{
		Email:    "jefferyz@futuredial.com",
		Password: "20040413",
	}
	ConnectDatabase()
	u, _ := Login(info)
	fmt.Print(u)
}

func TestTemplete(t *testing.T) {
	ConnectDatabase()
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)
	locations := QueryLocation()
	if err != nil {
		t.Error("error")
		return
	}
	abc, err := template.ParseFiles("../template/selectlocation.tpl")
	if err != nil {
		t.Error(err)
	}
	tmpl := template.Must(abc, err)
	var doc bytes.Buffer
	tmpl.Execute(&doc, locations)
	fmt.Print(doc.String())
}

func TestGetUser(t *testing.T) {
	ConnectDatabase()
	var user Users
	err := DB.Where("first_name = ? AND last_name = ?", "Jet", "Li").First(&user).Error
	if err != nil {
		fmt.Println(err)
	}
}
