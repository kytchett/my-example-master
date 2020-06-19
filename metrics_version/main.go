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
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/abc", index)
	http.Handle("/metrics", promhttp.Handler())
	metrics.Register()
	err := http.ListenAndServe(":5565", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	timer:=metrics.NewAdmissionLatency()
	metrics.RequestIncrease()
	num:=os.Getenv("Num")
	if num==""{
		ans10:=Factorial(10)
		str:="there is no env Num. Computed factorial of 10. The answer is "+strconv.Itoa(ans10)+" \n"
		_,err:=w.Write([]byte(str))
		if err!=nil{
			log.Println("err:"+err.Error()+" No\n")
		}
	}else{
		numInt,_:=strconv.Atoi(num)
		randNum:=rand.Intn(numInt) + 1
		ans:=Factorial(randNum)
		str:="there is env Num. Computed factorial of "+strconv.Itoa(randNum)+". The answer is "+strconv.Itoa(ans)+" \n"
		_,err:=w.Write([]byte(str))
		if err!=nil{
			log.Println("err:"+err.Error()+" Yes\n")
		}
	}
	timer.Observe()
}

func Factorial(n int)int{
	if n<=2{
		return n
	}else{
		return n * Factorial(n-1)
	}
}
