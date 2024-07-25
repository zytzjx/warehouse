package models

var needloadborrower bool = true
var loctions []Location

func QueryLocation() []Location {
	if needloadborrower {
		err := DB.Find(&loctions).Where("isdelete=?", 0).Order("location").Error
		if err != nil {
			needloadborrower = true
		} else {
			needloadborrower = false
		}
	}
	return loctions
}

func SetLoadCurLocation() {
	needloadborrower = true
}

func InsertBorrower(name string) (int, error) {
	loc := Location{Location: name}
	ret := DB.Create(loc)
	needloadborrower = true
	return loc.ID, ret.Error
}
