// +build !containers,!disable_notifcation

package app

import "github.com/bedminer1/distributing/notify"

func send_notification(msg string) {
	n := notify.New("Pomodoro", msg, notify.SeverityNormal)
	n.Send()
}
