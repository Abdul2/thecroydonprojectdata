//dataserver starts two servers. a file server to give access to files in datasource folder
//and http server to serve out the content
package dataserver

import (
	"net/http"

	"os"
	"fmt"

	"io/ioutil"
	"bufio"
)

var files []string
var content  []string

//Fileserverrouter returns a handle to source data folder (datasource)
func Fileserverrouter() http.Handler {

	return http.FileServer(http.FileSystem(http.Dir("datasource")))
}

//Display serves out the content
func Display(resp http.ResponseWriter, req *http.Request) {

	res, err := http.Get("http://localhost:8080/")

	if err != nil {
		fmt.Print(err)

		os.Exit(1)

	}

	scanner := bufio.NewScanner(res.Body)

	for scanner.Scan() {


		line := fmt.Sprint(scanner.Text())

		if len(line) > 10 {

			url := fmt.Sprintf("%s%s", "http://localhost:8080/", line[9:35])


			files = append(files, url)

		}

	}

	var counter int

	for v := range files {

		url := files[v]

		res, _ := http.Get(url)

		b, _ := ioutil.ReadAll(res.Body)

//		fmt.Println(string(b))

		content = append(content, string(b))

		counter++

	}

	//fmt.Println("numbr of files retrieved ", counter)

	fmt.Fprint(resp, content)

}
