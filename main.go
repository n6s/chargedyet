package main

import (
	// "fmt"
	"github.com/0xAX/notificator"
	"time"
)

var notify *notificator.Notificator

func main() {
	title := "Charged?"
	message := "Has it finished charging yet?"
	iconpath := "/home/rbrown/Pictures/charging.png"
	for {
		notifyMe(title, message, iconpath)

		timer1 := time.NewTimer(time.Minute * 30)
		<-timer1.C
	}
}

func notifyMe(title, message, icon string) {

	notify = notificator.New(notificator.Options{
		DefaultIcon: "/usr/share/icons/elementary-xfce/notifications/24/notification-message-im.png",
		AppName:     "My test App",
	})

	notify.Push(title, message, icon, notificator.UR_CRITICAL)
}
