package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"strconv"
	"example/metrics"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().Unix()) 			// 设置随机数种子
	http.HandleFunc("/abc", index)  		//普通服务请求处理
	http.Handle("/metrics", promhttp.Handler()) 	
	metrics.Register() 				//通过url方式访问prometheus数据
	err := http.ListenAndServe(":5565", nil) 	// 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	timer:=metrics.NewAdmissionLatency()
	metrics.RequestIncrease() 		// 请求计数器加一，标记当前时间
	num:=os.Getenv("Num")    		// 获取环境变量
	if num==""{ 				// 环境变量无效
		ans10:=Factorial(10) 		// 认计算10的阶乘
		str:="there is no env Num. Computed factorial of 10. The answer is "+strconv.Itoa(ans10)+" \n"
		_,err:=w.Write([]byte(str))
		if err!=nil{
			log.Println("err:"+err.Error()+" No\n")
		}
	}else{ 					//环境变量有效
		numInt,_:=strconv.Atoi(num)
		randNum:=rand.Intn(numInt) + 1	// 使计算的阶乘数大于0，获取需计算阶乘的随机数
		ans:=Factorial(randNum)		// 计算阶乘
		str:="there is env Num. Computed factorial of "+strconv.Itoa(randNum)+". The answer is "+strconv.Itoa(ans)+" \n"
		_,err:=w.Write([]byte(str))
		if err!=nil{
			log.Println("err:"+err.Error()+" Yes\n")
		}
	}
	timer.Observe() 			//记录所需时间
}

func Factorial(n int)int{  // 计算阶乘
	if n<=2{
		return n
	}else{
		return n * Factorial(n-1)
	}
}
