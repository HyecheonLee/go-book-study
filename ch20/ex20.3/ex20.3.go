package main

import (
	"github.com/HyecheonLee/go-book-study/ch20/fedex"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	//sender := &koreaPost.PostSender{}
	//SendBook("어린 왕자", sender)
	//SendBook("그리스인 조르바", sender)
}
