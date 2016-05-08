package main



import (
"net/http"
"os"
A "github.com/Abdul2/thecroydonprojectdata/dataserver"
)



func main(){



go func() {

if r := http.ListenAndServe(":8080", A.Fileserverrouter()); r != nil {

os.Exit(1)
}

}()

http.HandleFunc("/results", A.Display)
http.ListenAndServe(":9000", nil)



}
