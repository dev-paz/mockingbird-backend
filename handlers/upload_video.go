package handler

import(
  "net/http"
  "net/http/httputil"
  "net/url"
)

func handleUploadVideoRequest(w http.ResponseWriter, req *http.Request) {
  checkOpenshotIP()
   url, err := url.Parse("http://" + "3.236.10.84" + "/files/")
   if err != nil {
     panic(err)
   }
   req.URL, _ = url.Parse("/")

   proxy := httputil.NewSingleHostReverseProxy(url)

   req.URL.Host = url.Host
   req.URL.Scheme = url.Scheme
   req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
   req.SetBasicAuth(username, passwd)
   req.Host = url.Host

   proxy.ServeHTTP(w, req)
}
