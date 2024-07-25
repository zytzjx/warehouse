package controller

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zytzjx/warehouse/models"
	"gorm.io/gorm"

	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// /database id array for batch task
type RequestIDs struct {
	Ids []int `json:"ids"`
}

// GET /devices
// Find all devices
func FindAllHandset(c *gin.Context) {
	var devices []models.Handsets
	models.DB.Find(&devices)

	c.JSON(http.StatusOK, devices)
}

func GetLoctions() string {
	locations := models.QueryLocation()

	tmpl := template.Must(template.ParseFiles("template/selectlocation.tpl"))
	var doc bytes.Buffer
	tmpl.Execute(&doc, locations)
	return doc.String()
}

func Home(c *gin.Context) {
	filename := "static/index.html"
	if strings.Contains(c.Request.URL.Path, "edit") {
		filename = "static/edit.html"
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", data)
}

func AdminHome(c *gin.Context) {
	locations := models.QueryLocation()
	c.HTML(http.StatusOK, "index.tpl", locations)
}

func LoginOrUI(c *gin.Context) {
	if val, ok := c.Get("user"); ok {
		user := val.(models.Users)
		if user.Role != 10 {
			c.Redirect(http.StatusFound, "/index")
			return
		} else {
			c.Redirect(http.StatusFound, "/admin")
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "unkown the user"})
		return
	}
}

func Registration(c *gin.Context) {
	data, err := os.ReadFile("static/registration.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", data)
}

func Login(c *gin.Context) {
	//UserName := c.Query("username")
	//Password := c.Query("password") // shortcut for c.Request.URL.Query().Get("lastname")
	var info models.LOGINUSER
	//if UserName == "" || Password == "" {
	if err := c.ShouldBindQuery(&info); err != nil {
		fmt.Println(err)
		data, err := os.ReadFile("static/login.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	} else {
		// info := models.LOGINUSER{
		// 	Email:    UserName,
		// 	Password: Password,
		// }
		fmt.Println(info)
		u, err := models.Login(info)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(err.Error())
			return
		}
		fmt.Println(u)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userid": u.ID,
			"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
		})
		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Token string could not be created",
			})
			return
		}
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Auth", tokenString, 3600*24*30, "", "", false, true)
		fmt.Printf("%#v\n", u)
		if u.Role == 10 {
			c.Redirect(http.StatusFound, "/admin")
		} else {
			c.Redirect(http.StatusFound, "/index")
		}
		//c.JSON(http.StatusOK, u.ToLoginResponse())
	}
}

func Logout(c *gin.Context) {
	c.SetCookie("Auth", "", -1, "", "", false, true)
	c.Redirect(http.StatusFound, "/login")
}

func FDLogin(c *gin.Context) {
	var info models.LOGINUSER
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	u, err := models.Login(info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}
	//c.JSON(http.StatusOK, u.ToLoginResponse())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": u.ID,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Token string could not be created",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 3600*24*30, "", "", false, true)
	fmt.Printf("%#v\n", u)
	if u.Role == 10 {
		c.Redirect(http.StatusFound, "/admin")
	} else {
		c.Redirect(http.StatusFound, "/index")
	}
}

func Signup(c *gin.Context) {
	//{"firstname":"A","lastname":"b","email":"z@b.com","password":"!errors.Is","title":"AT"}
	var userinfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Fistname string `json:"firstname"`
		Lastname string `json:"lastname"`
		Title    string `json:"title"`
	}
	if err := c.BindJSON(&userinfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "requst format err:" + err.Error(),
		})
		return
	}

	fmt.Printf("%#v\n", userinfo)
	var user models.Users
	bUpdate := false
	err := models.DB.First(&user, "email = ?", userinfo.Email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("not find email")
		err = models.DB.Where("first_name = ? AND last_name = ?", userinfo.Fistname, userinfo.Lastname).First(&user).Error
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			// c.JSON(http.StatusInternalServerError, gin.H{
			// 	"error": "Query db failed:" + err.Error(),
			// })
			// return

		} else {
			fmt.Println("find user: " + userinfo.Fistname)
			bUpdate = true
		}
	} else {
		fmt.Println("found email")
		bUpdate = true
	}

	newuser := models.Users{
		FirstName: userinfo.Fistname,
		LastName:  userinfo.Lastname,
		Email:     userinfo.Email,
		Title:     userinfo.Title,
		Password:  models.HashPassword(userinfo.Password),
		FullName:  fmt.Sprintf("%s %s", userinfo.Fistname, userinfo.Lastname),
	}
	if !bUpdate {
		err = models.DB.Create(&newuser).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "User could not be created:" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "User Created",
		})
	} else {
		//newuser.ID = user.ID
		err = models.DB.Model(&user).Updates(newuser).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "User could not be updated:" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "User updated",
		})
	}
}

func DeleteDevice(c *gin.Context) {
	if val, ok := c.Get("user"); ok {
		user := val.(models.Users)
		if user.Role != 10 {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin can delete device, please use admin login"})
			return
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "unkown the user"})
		return
	}
	var request RequestIDs
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Model(models.Device{}).Where("id IN ?", request.Ids).Update("isDeleted", 1).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
		"ids":     request.Ids,
	})
}

func ChangeBorrower(c *gin.Context) {
	if val, ok := c.Get("user"); ok {
		user := val.(models.Users)
		if user.Role != 10 {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin can delete device, please use admin login"})
			return
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "unkown the user"})
		return
	}
	var request struct {
		RequestIDs
		Borrower string
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var loc models.Location
	var iduser int
	var err error
	if iduser, err = strconv.Atoi(request.Borrower); err == nil {
		fmt.Printf("%q looks like a number.\n", request.Borrower)
	}
	if iduser > 0 {
		loc.ID = iduser
	}
	if loc.ID == 0 {
		if err = models.DB.Where("location = ?", request.Borrower).First(&loc).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if err = models.DB.Model(models.Device{}).Where("id IN ?", request.Ids).Update("current_location_id", loc.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "change location",
		"ids":     request.Ids,
	})
}

// update to phone room
func ReturnWarehouse(c *gin.Context) {
	if val, ok := c.Get("user"); ok {
		user := val.(models.Users)
		if user.Role != 10 {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin can delete device, please use admin login"})
			return
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "unkown the user"})
		return
	}
	var request struct {
		User string `json:"user"`
		IDs  []int  `json:"ids"`
	}
	var err error
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if request.User == "" {
		//return device
		//1 is phone room
		if err := models.UpdateHandsetBorrower(1, request.IDs); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		//query user browerid from
		ssds := models.QueryLocation()
		curid := 0
		for _, element := range ssds {
			if element.Location == request.User {
				curid = element.ID
				break
			}
		}

		if curid == 0 {
			curid, err = models.InsertBorrower(request.User)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		if curid == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "not find borrower id"})
			return
		}

		//borrower
		if err := models.UpdateHandsetBorrower(curid, request.IDs); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "update borrower"})
}

// new device or update device
func Device(c *gin.Context) {
	if val, ok := c.Get("user"); ok {
		user := val.(models.Users)
		if user.Role != 10 {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin can delete device, please use admin login"})
			return
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "unkown the user"})
		return
	}

	var request models.Device
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// var locmaker models.Maker
	// if err := models.DB.Where("maker = ?", request.Maker).First(&locmaker).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// var loc models.Location
	// if err := models.DB.Where("location = ?", request.Borrower).First(&loc).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// newdev := models.Device{
	// 	MakerID:       locmaker.ID,
	// 	Model:         request.Model,
	// 	MarketingName: request.MarketingName,
	// 	Carrier:       request.Carrier,
	// 	ESN:           request.ESN,
	// 	PhoneNumber:   request.PhoneNumber,
	// 	FD_Model:      request.FD_Model,
	// 	Location:      request.Location,
	// 	BorrowerID:    loc.ID,
	// 	BarCode:       request.BarCode,
	// 	Note:          request.Note,
	// }
	if request.ID == 0 {
		//create new record
		if err := models.DB.Create(&request).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Add new device"})
	} else {
		//update new record
		//newdev.ID = request.ID
		if err := models.DB.Save(&request).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "update a device"})
	}
}

func EditNew(c *gin.Context) {
	locations := models.QueryLocation()
	makers := models.LoadMakers()

	c.HTML(http.StatusOK, "edit_new.tpl", gin.H{
		"Borrowers": locations,
		"Makers":    makers,
	})
}

func UpdateNote(c *gin.Context) {
	if val, ok := c.Get("user"); ok {
		user := val.(models.Users)
		if user.Role != 10 {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin can delete device, please use admin login"})
			return
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "unkown the user"})
		return
	}
	var request struct {
		Id    int    `json:"id"`
		Notes string `json:"note"`
	}
	var err error
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.DB.Model(&models.Device{ID: request.Id}).Updates(models.Device{Note: request.Notes}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "update note"})
}
