package models

// /view device_infos
type Handsets struct {
	ID            int    `gorm:"column:ID" json:"id"`
	Maker         string `json:"maker"`
	Model         string `json:"model"`
	MarketingName string `gorm:"column:MarketingName" json:"marketingname"`
	Carrier       string `json:"carrier"`
	ESN           string `gorm:"column:ESN"  json:"esn"`
	PhoneNumber   string `gorm:"column:PhoneNumber" json:"phonenumber"`
	FD_Model      string `gorm:"column:FD_Model" json:"fdmodel"`
	Location      string `json:"location"`
	Borrower      string `gorm:"column:CurLocation" json:"borrower"`
	BarCode       string `json:"barcode"`
	Note          string `json:"note"`
}

func (Handsets) TableName() string {
	return "device_infos" //this is tableview
}

// phone db table
type Device struct {
	ID            int    `gorm:"column:ID" json:"id"`
	MakerID       int    `json:"maker"`
	Model         string `json:"model"`
	MarketingName string `json:"marketingname"`
	Carrier       string `json:"carrier"`
	ESN           string `gorm:"column:esn" json:"esn"`
	PhoneNumber   string `json:"phonenumber"`
	FD_Model      string `gorm:"column:fd_model" json:"fdmodel"`
	Location      string `json:"location"`
	BorrowerID    int    `gorm:"column:current_location_id" json:"borrower"`
	BarCode       string `gorm:"column:barcode" json:"barcode"`
	Note          string `gorm:"column:notes" json:"note"`
	IsDeleted     int8   `gorm:"column:isDeleted" json:"-"`
}

func (Device) TableName() string {
	return "tbl_handset_infos"
}

func UpdateHandsetBorrower(bowrrorid int, ids []int) error {
	return DB.Model(Device{}).Where("id IN ?", ids).Updates(map[string]interface{}{"current_location_id": bowrrorid}).Error
}

// Role===10 it is admin, others are normal user
type Users struct {
	ID         int
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Title      string    `json:"title"`
	Password   string    `json:"password"`
	FullName   string    `json:"full_name" gorm:"->;type:GENERATED ALWAYS AS (concat(first_name,' ',last_name));default:(-);"`
	Role       int8      `json:"role" gorm:"default:1"`
	CreateDate LocalTime `json:"create_date" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdateDate LocalTime `json:"update_date"`
}

func (Users) TableName() string {
	return "tbl_users"
}

func (u *Users) ToLoginResponse() map[string]interface{} {
	dd := make(map[string]interface{})
	dd["first_name"] = u.FirstName
	dd["last_name"] = u.LastName
	dd["email"] = u.Email
	dd["title"] = u.Title
	dd["role"] = u.Role
	return dd
}

type Location struct {
	ID       int
	Location string
}

func (Location) TableName() string {
	return "tbl_cur_location"
}

type Maker struct {
	ID    int
	Maker string
}

func (Maker) TableName() string {
	return "tbl_makers"
}
