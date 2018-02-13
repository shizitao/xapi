package main
import (
    "io/ioutil"
    "net/http"
    "net/url"
    "fmt"
    "encoding/json"
)
 
//----------------------------------
// 笑话大全调用示例代码 － 聚合数据
// 在线接口文档：http://www.juhe.cn/docs/95
//----------------------------------
 
const APPKEY = "52c5c99d07dd27fc6b085dd30c19fa3a" //您申请的APPKEY
 
//1.按更新时间查询笑话
func Request1(){
    //请求地址
    juheURL :="http://japi.juhe.cn/joke/content/list.from"
 
    //初始化参数
    param:=url.Values{}
 
    //配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
    param.Set("sort","") //类型，desc:指定时间之前发布的，asc:指定时间之后发布的
    param.Set("page","") //当前页数,默认1
    param.Set("pagesize","") //每次返回条数,默认1,最大20
    param.Set("time","") //时间戳（10位），如：1418816972
    param.Set("key",APPKEY) //您申请的key
 
 
    //发送请求
    data,err:=Get(juheURL,param)
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err)
    }else{
        var netReturn map[string]interface{}
        json.Unmarshal(data,&netReturn)
        if netReturn["error_code"].(float64)==0{
            fmt.Printf("接口返回result字段是:\r\n%v",netReturn["result"])
        }
    }
}
 
//2.最新笑话
func Request2(){
    //请求地址
    juheURL :="http://japi.juhe.cn/joke/content/text.from"
 
    //初始化参数
    param:=url.Values{}
 
    //配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
    param.Set("page","") //当前页数,默认1
    param.Set("pagesize","") //每次返回条数,默认1,最大20
    param.Set("key",APPKEY) //您申请的key
 
 
    //发送请求
    data,err:=Get(juheURL,param)
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err)
    }else{
        var netReturn map[string]interface{}
        json.Unmarshal(data,&netReturn)
        if netReturn["error_code"].(float64)==0{
            fmt.Printf("接口返回result字段是:\r\n%v",netReturn["result"])
        }
    }
}
 
//3.按更新时间查询趣图
func Request3(){
    //请求地址
    juheURL :="http://japi.juhe.cn/joke/img/list.from"
 
    //初始化参数
    param:=url.Values{}
 
    //配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
    param.Set("sort","") //类型，desc:指定时间之前发布的，asc:指定时间之后发布的
    param.Set("page","") //当前页数,默认1
    param.Set("pagesize","") //每次返回条数,默认1,最大20
    param.Set("time","") //时间戳（10位），如：1418816972
    param.Set("key",APPKEY) //您申请的key
 
 
    //发送请求
    data,err:=Get(juheURL,param)
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err)
    }else{
        var netReturn map[string]interface{}
        json.Unmarshal(data,&netReturn)
        if netReturn["error_code"].(float64)==0{
            fmt.Printf("接口返回result字段是:\r\n%v",netReturn["result"])
        }
    }
}
 
//4.最新趣图
func Request4(){
    //请求地址
    juheURL :="http://japi.juhe.cn/joke/img/text.from"
 
    //初始化参数
    param:=url.Values{}
 
    //配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
    param.Set("page","") //当前页数,默认1
    param.Set("pagesize","") //每次返回条数,默认1,最大20
    param.Set("key",APPKEY) //您申请的key
 
 
    //发送请求
    data,err:=Get(juheURL,param)
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err)
    }else{
        var netReturn map[string]interface{}
        json.Unmarshal(data,&netReturn)
        if netReturn["error_code"].(float64)==0{
            fmt.Printf("接口返回result字段是:\r\n%v",netReturn["result"])
        }
    }
}
 
 
 
// get 网络请求
func Get(apiURL string,params url.Values)(rs[]byte ,err error){
    var Url *url.URL
    Url,err=url.Parse(apiURL)
    if err!=nil{
        fmt.Printf("解析url错误:\r\n%v",err)
        return nil,err
    }
    //如果参数中有中文参数,这个方法会进行URLEncode
    Url.RawQuery=params.Encode()
    resp,err:=http.Get(Url.String())
    if err!=nil{
        fmt.Println("err:",err)
        return nil,err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
 
// post 网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values)(rs[]byte,err error){
    resp,err:=http.PostForm(apiURL, params)
    if err!=nil{
        return nil ,err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
