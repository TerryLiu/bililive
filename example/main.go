package main

import (
	"flag"
	"log"

	"github.com/zboyco/bililive"
)

func main() {
	roomID := flag.Int("id", 0, "id")
	flag.Parse()
	if *roomID <= 0 {
		log.Fatalln("房间号错误!")
		return
	}
	liveRoom := &bililive.LiveRoom{
		RoomID: *roomID,
		ReceivePopularValue: func(v uint32) {
			log.Printf("【人气值】%v", v)
		},
		UserEnter: func(m *bililive.UserEnterModel) {
			log.Printf("【用户进入】欢迎 %v(%v) 进入直播间", m.UserName, m.UserID)
		},
		GuardEnter: func(m *bililive.GuardEnterModel) {
			log.Printf("【舰长进入】欢迎 舰长 - %v(%v) 进入直播间", m.UserName, m.UserID)
		},
		ReceiveMsg: func(msg *bililive.MsgModel) {
			log.Printf("【弹幕】%v:  %v", msg.UserName, msg.Content)
		},
		ReceiveGift: func(gift *bililive.GiftModel) {
			log.Printf("【礼物】%v:  %v(%v) * %v  连击 %v", gift.UserName, gift.GiftName, gift.GiftID, gift.Num, gift.Combo)
		},
		GiftComboEnd: func(m *bililive.ComboEndModel) {
			log.Printf("【连击结束】%v 赠送 %v(价值%v) 总共连击 %v 次", m.UserName, m.GiftName, m.Price, m.ComboNum)
		},
		GuardBuy: func(m *bililive.GuardBuyModel) {
			log.Printf("【上船】欢迎 %v - %v(%v) 上船", m.GiftName, m.UserName, m.UserID)
		},
	}
	liveRoom.Start()
}
