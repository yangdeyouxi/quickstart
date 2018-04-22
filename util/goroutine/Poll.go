package goroutine

import (
	"fmt"
)

type Pool struct {
	Queue chan func() error;
	RuntineNumber int;
	Total int;

	Result chan error;
	FinishCallback func();
}

/**
	初始化
	runtineNumber	并行任务数
	total			总任务数，超出并行任务数的会进入队列
 */
func (self *Pool) Init(runtineNumber int,total int)  {
	self.RuntineNumber = runtineNumber;
	self.Total = total;
	self.Queue = make(chan func() error, total);
	self.Result = make(chan error, total);
}

//开始
func (self *Pool) Start()  {
	//开启 number 个goruntine
	for i:=0;i<self.RuntineNumber;i++ {
		go func() {
			for {
				task,ok := <-self.Queue
				if !ok {
					break;
				}
				err := task();
				self.Result <- err;
			}
		}();
	}

	//获取每个任务的处理结果
	for j:=0;j<self.RuntineNumber;j++ {
		res,ok := <-self.Result;
		if !ok {
			break;
		}
		if res != nil {
			fmt.Println(res);
		}
	}

	//结束回调函数
	if self.FinishCallback != nil {
		self.FinishCallback();
	}
}

//关闭
func (self *Pool) Stop()  {
	close(self.Queue);
	close(self.Result);
}

//添加任务
func (self *Pool) AddTask(task func() error)  {
	self.Queue <- task;
}

//设置结束回调
func (self *Pool) SetFinishCallback(fun func())  {
	self.FinishCallback = fun;
}


func init(){
	var p Pool;
	p.Init(9, 20);
}

//下面是例子
//func main()  {
//	var p Pool;
//	url := []string{"11111","22222","33333","444444","55555","66666","77777","88888","999999"};
//	p.Init(9, len(url));
//
//	for i := range url {
//		u := url[i];
//		p.AddTask(func() error {
//			return Download(u);
//		});
//	}
//
//	p.SetFinishCallback(DownloadFinish);
//	p.Start();
//	p.Stop();
//}
//
//func Download(url string) error {
//	time.Sleep(1*time.Second);
//	fmt.Println("Download " + url);
//	return nil;
//}
//
//func DownloadFinish()  {
//	fmt.Println("Download finsh");
//}
