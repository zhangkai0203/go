package main
//文件上传
import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter,r *http.Request)  {
	temp,_ := template.ParseFiles("D:/go/src/fileDemo/index.html")
    temp.Execute(w,nil)
}
func upload(w http.ResponseWriter,r *http.Request)  {
    //文件名称
	name := r.FormValue("name")

    //文件对象   文件信息   错误信息
    file,hader,_ := r.FormFile("file")
    //截取扩展名
    ex := hader.Filename[strings.LastIndex(hader.Filename,"."):]
    //读取信息
    b,_ := ioutil.ReadAll(file)
    //保存信息
    ioutil.WriteFile("d:/" + name + ex,b,0777)
}
func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/upload",upload)
	http.ListenAndServe(":9000",nil)
}
