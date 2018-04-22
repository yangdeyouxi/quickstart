package biquge

import (
	"net/http"
	"quickstart/util/log"
	//"fmt"
	"io"
	//"github.com/axgle/mahonia"
	"quickstart/util/convert"
)

//这个类用来扒笔趣阁的小说：http://www.xbqge.com/

//扒首页
func getHome(){
	log.Console("start reptile")
	resp,err := http.Get("http://www.xbqge.com/")
	if err != nil{
		log.Error("笔趣阁首页出错:" + err.Error())
		return;
	}

	defer resp.Body.Close()

	statusCode := resp.StatusCode;
	if(statusCode == 200){
		buf := make([]byte,1024*4)
		var data string;
		for{
			n,err := resp.Body.Read(buf)
			if err != nil{
				if err != io.EOF {
					log.Warn(err.Error())
				}
				break
			}
			data += string(buf[:n])
		}
		//获取到网页源码：data
		//log.Info(data)
		result := convert.ConvertToString(data,"gbk","utf-8")
		log.Info(result)
		//fmt.Println("biquge:" + result)
	}
}

func init(){
	go getHome();
}
