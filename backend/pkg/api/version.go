package api

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// our struct which contains the complete
// array of all Users in the file
type Users struct {
	XMLName    xml.Name   `xml:"testsuites"`
	TestSuites TestSuites `xml:"testsuite"`
}

// a simple struct which contains all our
// social links
type TestSuites struct {
	XMLName   xml.Name    `xml:"testsuite"`
	TestSuite []TestCases `xml:"testcase"`
}

type TestCases struct {
	XMLName xml.Name `xml:"testcase"`
	Failure Failure  `xml:"failure,omitempty"`
	Name    string   `xml:"name,attr"`
	Status  string   `xml:"status,attr"`
	Time    string   `xml:"time,attr"`
}

type Failure struct {
	XMLName xml.Name `xml:"failure,omitempty"`
	Message string   `xml:"message,attr"`
}

// Version godoc
// @Summary Version
// @Description returns quality backend version
// @Tags Version API
// @Produce json
// @Router /api/version [get]
// @Success 200 {object} api.MapResponse
func (s *Server) versionHandler(w http.ResponseWriter, r *http.Request) {
	/*result := map[string]string{
		"version": version.VERSION,
	}*/

	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	fmt.Println(r.PostForm)
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	// copy example
	f, _ := os.OpenFile("./tmp", os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(f, file)

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	xmlFile, err := os.Open(f.Name())
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// Remove the file after reading
	os.Remove("./tmp")

	// we initialize our Users array
	var users Users
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &users)

	for _, tc := range users.TestSuites.TestSuite {
		fmt.Println(tc.Name)
		fmt.Println(tc.Status)
		fmt.Println(tc.Time)
	}

	s.JSONResponse(w, r, users)
}
