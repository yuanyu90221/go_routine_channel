# go_routine_channel

## 關鍵概念

兩個goruntine 

一個負責寫入channel ,
一個負責讀出channel

首先需要 用sync.WaitGroup讓 main routine等待 兩個routien執行完在結束

go func(ch <- chan int){

  for i := range ch { // channel 有讀到值就會印出來

    fmt.Println(i)

}

wg.Done()

}(ch)

go func(ch chan<- int) {

  ch <- 100

  ch <- 101

  ch<- 102

  // 注意這邊需要把channel close起來 因為結束之後如果 channel不close

  // read channel就會卡在那邊造成Deadlock

close(ch)

wg.Done()

}(ch)


如果是buffered  channel 因為不是blocking

所以如果塞入buffer channel的值 多於讀出的值 部會造成deadlock



可以利用 v, ok := <-ch 判斷 channel 是否被close