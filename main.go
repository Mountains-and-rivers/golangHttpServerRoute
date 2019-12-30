package main

import (
    "log"
    "net/http"
    "mypkg/httpHandle"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

func postBody(w http.ResponseWriter, r *http.Request, _ httpHandle.Params) {

    var user map[string]interface{}
    body, _ := ioutil.ReadAll(r.Body)

    json.Unmarshal(body, &user)

    fmt.Println("获取json中的username:", user["username"])
    fmt.Println("获取json中的password:", user["password"].(string))

    fmt.Fprint(w, "post请求!\n")
}
func getBody(w http.ResponseWriter, r *http.Request, _ httpHandle.Params) {

    keys, ok := r.URL.Query()["key"]

    if !ok || len(keys) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }
    
    // Query()["key"] will return an array of items,
    // we only want the single item.
    key := keys[0]
    log.Println("获取json中的key: " + string(key))
    fmt.Fprint(w, "get 请求!\n")
}
func Hello(w http.ResponseWriter, r *http.Request, ps httpHandle.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
func HelloPost(w http.ResponseWriter, r *http.Request, ps httpHandle.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
    router := httpHandle.New()
    router.POST("/post", postBody)
    router.GET("/get", getBody)
    router.GET("/hello/:name", Hello)

    log.Fatal(http.ListenAndServe(":8080", router))
}