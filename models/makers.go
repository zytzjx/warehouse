package models

var needloadmaker bool = true
var makers []Maker

func LoadMakers() []Maker {
	if needloadmaker {
		err := DB.Find(&makers).Order("maker").Error
		if err != nil {
			needloadmaker = true
		} else {
			needloadmaker = false
		}
	}
	return makers
}

func SetLoadMaker() {
	needloadmaker = true
}
