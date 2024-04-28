package controllers

import (
	"auto_excel/helpers"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"

	models "auto_excel/models"
	"html/template"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	temp := helpers.Include("dashboard")

	view, err := template.ParseFiles(temp...)

	helpers.ErrH("Error in Dashboard Index: ", err)
	data := make(map[string]interface{})
	data["Students"] = models.Person{}.ReadAll()

	view.ExecuteTemplate(w, "Index", data)
}

func (dashboard Dashboard) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	temp := helpers.Include("add")

	view, err := template.ParseFiles(temp...)

	helpers.ErrH("Error in Dashboard Index: ", err)
	data := make(map[string]interface{})
	data["Students"] = models.Person{}.ReadAll()

	view.ExecuteTemplate(w, "Index", data)
}

func (dashboard Dashboard) Excel(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	fmt.Println("+++EXCEL")
	//ExcelModelsfromArray([[Value,Column,Row,Width,Height,IsImage]])
	var excelData [][]interface{}
	arr := []interface{}{"T/b", "A", 1, 5, nil, nil}
	excelData = append(excelData, arr)
	arr = []interface{}{"Familiýasy", "B", 1, 20, nil, nil}
	excelData = append(excelData, arr)
	arr = []interface{}{"Ady", "C", 1, 20, nil, nil}
	excelData = append(excelData, arr)
	arr = []interface{}{"Atasynyň ady", "D", 1, 20, nil, nil}
	excelData = append(excelData, arr)
	arr = []interface{}{"Doglan senesi", "E", 1, 15, nil, nil}
	excelData = append(excelData, arr)
	arr = []interface{}{"Salgysy", "F", 1, 25, nil, nil}
	excelData = append(excelData, arr)
	arr = []interface{}{"Işleýän/okaýan ýeri, wezipesi", "G", 1, 25, nil, nil}
	excelData = append(excelData, arr)
	arr = []interface{}{"Telefon belgisi", "H", 1, 15, nil, nil}
	excelData = append(excelData, arr)
	arr = []interface{}{"ID Number", "I", 1, 15, nil, nil}
	excelData = append(excelData, arr)
	arr = []interface{}{"EPC Number", "J", 1, 15, nil, nil}
	excelData = append(excelData, arr)
	arr = []interface{}{"Fotosuraty", "K", 1, 20, nil, nil}
	excelData = append(excelData, arr)

	jsonData := models.Person{}.ReadAll()
	for i, model := range jsonData {
		column := "A"
		arr = []interface{}{i + 1, column, i + 2, nil, nil, false}
		excelData = append(excelData, arr)
		for _, v := range model.ToArr() {
			column = string(rune(rune(column[0]) + 1))
			arr = []interface{}{v, column, i + 2, nil, nil, false}
			excelData = append(excelData, arr)
		}
		if *model.Photo != ".base/uploads/empty.jpeg" {
			column = string(rune(rune(column[0]) + 1))
			arr = []interface{}{*model.Photo, column, i + 2, nil, 125, true}
			excelData = append(excelData, arr)
		}
	}

	excelModels, err := helpers.ExcelModelsfromArray(excelData)
	helpers.ErrH("Error ExcelModelsfromArra: ", err)
	fileName := "netije.xlsx"
	helpers.ExcelCreate(fileName, excelModels)
	//Download========
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(fileName))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, fileName)
	//================
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (dashboard Dashboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// slug := slug.Make(title)
	lastName := r.FormValue("last_name")
	firstName := r.FormValue("first_name")
	midleName := r.FormValue("midle_name")
	address := r.FormValue("address")
	workAddress := r.FormValue("work_address")
	phone := r.FormValue("phone")
	idNumber := r.FormValue("id_number")
	epcNumber := r.FormValue("epc_number")

	//Birth day start
	birthDay := r.FormValue("birth_day")
	date := birthDay

	// var date time.Time
	// if birthDay != "" {
	// 	birthDay = "00/00/0000"
	// 	split := strings.Split(birthDay, "/")
	// 	fmt.Println("split:", split)
	// 	dateFormat := split[len(split)-1] + "-" + split[0] + "-" + split[1]
	// 	_date, err := time.Parse("2006-01-02", dateFormat)
	// 	date = _date
	// 	helpers.ErrH("Error str to time:", err)

	// }
	//Birth day end

	//Upload start
	filePath := ".base/uploads/empty.jpeg"
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("photo")
	if err != nil {
		fmt.Println("Error in file upload:", err)
	} else {
		rand1 := strconv.Itoa(rand.Intn(1000) + 1000)
		rand2 := strconv.Itoa(rand.Intn(1000) + 1000)
		rand3 := strconv.Itoa(rand.Intn(1000) + 1000)
		rand4 := strconv.Itoa(rand.Intn(1000) + 1000)
		random := rand1 + rand2 + rand3 + rand4
		filePath = ".base/uploads/" + random + header.Filename

		f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		helpers.ErrH(err)
		_, err = io.Copy(f, file)
		helpers.ErrH(err)

	}
	//Upload end

	models.Person{
		LastName:    lastName,
		FirstName:   firstName,
		MidleName:   midleName,
		Phone:       phone,
		Address:     &address,
		WorkAddress: &workAddress,
		IDNumber:    &idNumber,
		EPCNumber:   &epcNumber,
		Photo:       &filePath,
		Birthday:    &date,
	}.Add()
	// helpers.SetAlert(w, r, "Post Goşuldy!", "fa fa-check-circle", "notification-success")
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (dashboard Dashboard) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	helpers.ErrH("Error in Delete:", err)
	person := models.Person{}.Read(id)
	fmt.Println("Delete person:", person)
	if *person.Photo != ".base/uploads/empty.jpeg" {
		helpers.DeleteFile(*person.Photo)
	}
	person.Delete()
	// if person.ID == uint(id) {
	// 	person.Delete()
	// 	// helpers.SetAlert(w, r, "This item deleted successsfully!", "fa fa-check-circle", "notification-success")
	// } else {
	// 	// helpers.SetAlert(w, r, "This item was not found!", "fa fa-times-circle-o", "notification-danger")
	// }
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
