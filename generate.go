package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	// "strings"

	"github.com/fatih/color"
)

type Publication struct {
	Description   string `json:"description"`
	ImgUrl        string `json:"img_url"`
	Title         string `json:"title"`
	WriteupUrl    string `json:"writeup_url"`
	OrderPosition int16  `json:"order_position"`
}

type Project struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Uri           string `json:"uri"`
	ImgUrl        string `json:"img_url"`
	Technologies  string `json:"technologies"`
	GithubUrl     string `json:"github_url"`
	WriteupUrl    string `json:"writeup_url"`
	OrderPosition int16  `json:"order_position"`
}

type PersonalData struct {
	FirstName               string        `json:"first_name"`
	LastName                string        `json:"last_name"`
	Slug                    string        `json:"slug"`
	GithubHandle            string        `json:"github_handle"`
	ProfessionalTitle       string        `json:"professional_title"`
	WorkDescription         string        `json:"work_description"`
	AboutShortDescription   string        `json:"about_short_description"`
	AboutDescription        string        `json:"about_description"`
	ProfileImgUrl           string        `json:"profile_img_url"`
	ContactDescription      string        `json:"contact_description"`
	Email                   string        `json:"email"`
	DribbleHandle           string        `json:"dribbble_handle"`
	LinkedinHandle          string        `json:"linkedin_handle"`
	ResumeUrl               string        `json:"resume_document_url"`
	PublicationsDescription string        `json:"publication_description"`
	Publications            []Publication `json:"publications"`
	Projects                []Project     `json:"projects"`
}

type Data struct {
	Content    PersonalData
	StylesPath string
}

var boldGreen *color.Color = color.New(color.FgGreen, color.Bold)
var boldMagenta *color.Color = color.New(color.FgMagenta, color.Bold)
var boldRed *color.Color = color.New(color.FgRed, color.Bold)
var boldWhite *color.Color = color.New(color.FgWhite, color.Bold)

func main() {
	filename := flag.String("file", "filename", "name of file to parse")
	directory := flag.String("dir", "directory", "name of directory of files to parse")
	flag.Parse()

	fileNum, size := create(*filename, *directory)

	boldGreen.Print("Success!")
	fmt.Print(" Generated ")
	boldWhite.Print(fileNum)
	fmt.Printf(" pages (%.1fkB total)\n", size)

}

func create(filename, directory string) (int16, float64) {
	var fileNum int16 = 0
	var size float64 = 0
	if directory != "directory" {
		fileNum, size = createManyHTML(directory)
	}

	if filename != "filename" {
		data := loadFileContent(filename)
		fileNum, size = createHTML("", strings.SplitN(filename, ".", 2)[0], "portfolio.tmpl", data)
	}
	return fileNum, size
}

func loadFileContent(filename string) Data {
	var person PersonalData
	if strings.SplitN(filename, ".", 2)[1] == "json" {
		data, err := ioutil.ReadFile(filename)
		handleError(err)
		err = json.Unmarshal(data, &person)
	}
	sort.SliceStable(person.Projects, func(i, j int) bool {
		return person.Projects[i].OrderPosition < person.Projects[j].OrderPosition
	})

	sort.SliceStable(person.Publications, func(i, j int) bool {
		return person.Publications[i].OrderPosition < person.Publications[j].OrderPosition
	})

	var dircount string = ""
	for i := 0; i < len(strings.SplitN(filename, "/", -1))-1; i++ {
		dircount = dircount + "../"
	}

	return Data{person, dircount + "styles.css"}
}

func createHTML(dir, filename, templ string, data Data) (int16, float64) {
	var path string = ""
	if dir != "" {
		path = dir + "/"
	}
	htmlFile, osErr := os.Create(path + filename + ".html")
	handleError(osErr)
	t := template.Must(template.ParseFiles(templ))
	execErr := t.Execute(htmlFile, data)
	handleError(execErr)
	return 1, calculateSize(path + filename + ".html")
}

func createManyHTML(directory string) (int16, float64) {
	files, err := os.ReadDir(directory)
	handleError(err)
	var counter int16 = 0
	var size float64 = 0
	for _, file := range files {
		path := directory + "/" + file.Name()
		stat, errStat := os.Stat(path)
		handleError(errStat)
		if !stat.IsDir() && (strings.SplitN(stat.Name(), ".", 2)[1] == "txt" || strings.SplitN(stat.Name(), ".", 2)[1] == "md") {
			counter++
			data := loadFileContent(path)
			c, s := createHTML(directory, strings.SplitN(file.Name(), ".", 2)[0], "template.tmpl", data)
			counter += c
			size += s
		} else if stat.IsDir() {
			createManyHTML(path)
		} else {
			boldMagenta.Print("Warning:")
			fmt.Print(" File ")
			boldWhite.Print(stat.Name())
			fmt.Println(" does not match type .txt or .md")
		}
	}
	return counter, size
}

func handleError(err error) {
	if err != nil {
		boldRed.Print("Error:")
		fmt.Print(err)
		panic(err)
	}
}

func calculateSize(filename string) float64 {
	file, errOpen := os.Open(filename)
	handleError(errOpen)
	size, errSeek := file.Seek(0, 2)
	handleError(errSeek)
	return float64(size) / 1024.0
}
