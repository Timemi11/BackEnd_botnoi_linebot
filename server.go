package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

func main() {
	bot, err := linebot.New(
		//1.Channel secret
		//2.Channel access token
		"76a10f93fe6dacee4bf963aabebea8b2",
		"frLr21hjDvTbx039Msl6QlSarzDzG302R3SK3fkPf5Q+0ICrrxgWCMEuLmx8zPrO7ki7TzY/5Ny23hnvmOgEjUP6ct16wme3sLW9wbOKpJLFToRuuSnVVFmHemtiwadT01oQd2MrDS0o6Ojz8ooJOgdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/chat", handleChat(bot))
	http.ListenAndServe(":5000", nil)
}

func handleChat(bot *linebot.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			handleError(err, w)
			return
		}
		for _, event := range events {
			switch event.Type {
			case linebot.EventTypeMessage:
				handleMessage(event, bot)
			}
		}
	}
}

func handleMessage(event *linebot.Event, bot *linebot.Client) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		handleTextMessage(message, event, bot)
	}
}

func handleTextMessage(message *linebot.TextMessage, event *linebot.Event, bot *linebot.Client) {
	text := strings.ToLower(strings.TrimSpace(message.Text))

	switch text {
	case "quick":
		sendQuickReply(event, bot)
	case "text":
		sendTextMessage(event, bot)
	case "ช่วยเหลือ":
		sendHelpMessage(event, bot)
	case "ประวัติ":
		sendHistoryMessage(event, bot)
	case "อังกฤษ":
		sendFlexMessageUK(event, bot)
	case "อิตาลี":
		sendFlexMessageItaly(event, bot)
	case "จีน":
		sendFlexMessageChina(event, bot)
	case "ฝรั่งเศส":
		sendFlexMessageFrance(event, bot)
	case "image":
		sendImageMessage(event, bot)
	case "button":
		sendButtonMessage(event, bot)
	case "carousel":
		sendCarouselMessage(event, bot)
	default:
		sendUnknownMessage(event, bot)
	}
}

func sendQuickReply(event *linebot.Event, bot *linebot.Client) {
	// Quick reply implementation
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เลือกรายการประเทศที่สนใจ").WithQuickReplies(
		linebot.NewQuickReplyItems(
			linebot.NewQuickReplyButton(
				"https://upload.wikimedia.org/wikipedia/commons/thumb/0/03/Flag_of_Italy.svg/125px-Flag_of_Italy.svg.png",
				linebot.NewMessageAction("อิตาลี", "อิตาลี")),
			linebot.NewQuickReplyButton(
				"https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/Flag_of_France_%281794%E2%80%931815%2C_1830%E2%80%931974%29.svg/125px-Flag_of_France_%281794%E2%80%931815%2C_1830%E2%80%931974%29.svg.png",
				linebot.NewMessageAction("ฝรั่งเศส", "ฝรั่งเศส")),
			linebot.NewQuickReplyButton(
				"https://th.bing.com/th?id=ODL.b82e3bf33d5e242bb99ce1748c8de669&w=143&h=95&c=10&rs=1&qlt=99&o=6&pid=13.1",
				linebot.NewMessageAction("จีน", "จีน")),
			linebot.NewQuickReplyButton(
				"https://upload.wikimedia.org/wikipedia/commons/thumb/a/a5/Flag_of_the_United_Kingdom_%281-2%29.svg/125px-Flag_of_the_United_Kingdom_%281-2%29.svg.png",
				linebot.NewMessageAction("อังกฤษ", "อังกฤษ")),
			linebot.NewQuickReplyButton(
				"",
				linebot.NewLocationAction("ส่งที่อยู่")),
		),
	)).Do(); err != nil {
		log.Println(err)
	}
}

func sendTextMessage(event *linebot.Event, bot *linebot.Client) {

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("สวัสดีนี้คือข้อความเริ่มต้น ลองพิมพ์ quick button carousel ดูสิ!! ")).Do(); err != nil {
		log.Println(err)
	}
}

func sendImageMessage(event *linebot.Event, bot *linebot.Client) {

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://d3h1lg3ksw6i6b.cloudfront.net/media/image/2022/01/27/57fb54945d744a78b9aff62df8ffdc41_How+To+Celebrate+Like+A+True+Italian5.jpg", "https://d3h1lg3ksw6i6b.cloudfront.net/media/image/2022/01/27/57fb54945d744a78b9aff62df8ffdc41_How+To+Celebrate+Like+A+True+Italian5.jpg")).Do(); err != nil {
		log.Println(err)
	}

}

func sendButtonMessage(event *linebot.Event, bot *linebot.Client) {
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(`
						{
							"type": "bubble",
							"header": {
								"type": "box",
								"layout": "horizontal",
								"contents": [
									{
										"type": "text",
										"text": "ประเทศจีน",
										"size": "lg",
										"align": "center"
									}
								]
							},
							"hero": {
								"type": "image",
								"url": "https://cdn.britannica.com/89/179589-138-3EE27C94/Overview-Great-Wall-of-China.jpg?w=800&h=450&c=crop",
								"size": "full",
								"aspectMode": "cover",
								"aspectRatio": "320:213"
							},
							"body": {
								"type": "box",
								"layout": "vertical",
								"contents": [
									{
										"type": "text",
										"text": "กำแพงเมืองจีน",
										"size": "md",
										"wrap": true
									}
								]
							},
							"footer": {
								"type": "box",
								"layout": "horizontal",
								"contents": [
									{
										"type": "button",
										"action": {
											"type": "message",
											"label": "ประวัติ",
											"text": "ประวัติ"
										}
									},
									{
										"type": "button",
										"action": {
											"type": "uri",
											"label": "สถานที่ตั้ง",
											"uri": "https://maps.app.goo.gl/mMWDCoJzV7nwyLxG8"
										}
									}
								]
							}
						}`))
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return
	}
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("FlexWithJSON", flexContainer)).Do(); err != nil {
		log.Println("Error sending Flex Message:", err)
	}
}

func sendCarouselMessage(event *linebot.Event, bot *linebot.Client) {

	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(`
	{
		"type": "carousel",
		"contents": [
		  {
			"type": "bubble",
			"direction": "ltr",
			"hero": {
			  "type": "image",
			  "url": "https://p.mgcdn.me/Z1KCKH5x-zpVgxn6w1ys9Neasx0=//065/031/000/20231130_101648_1701314208_6567fea04ce54.jpeg",
			  "size": "full",
			  "aspectRatio": "20:13",
			  "aspectMode": "cover"
			},
			"body": {
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "นางพยัคฆ์",
				  "weight": "bold",
				  "size": "xl",
				  "wrap": true,
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "Master Tigress is one of the main supporting characters of the Kung Fu Panda franchise. She is a member of the Furious Five as well as one of Master Shifu's students at the Jade Palace. She is a master of the Tiger Style of kung fu.",
				  "contents": []
				}
			  ]
			},
			"footer": {
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "button",
				  "action": {
					"type": "uri",
					"label": "ข้อมูล",
					"uri": "https://hmong.in.th/wiki/Master_Shifu"
				  },
				  "color": "#2DCF1EFF",
				  "style": "primary"
				}
			  ]
			}
		  },
		  {
			"type": "bubble",
			"hero": {
			  "type": "image",
			  "url": "https://down-th.img.susercontent.com/file/132305b22e0c2256c6336eb4a33465e7",
			  "size": "full",
			  "aspectRatio": "20:13",
			  "aspectMode": "cover"
			},
			"body": {
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "อาจารย์ชิฟู",
				  "weight": "bold",
				  "size": "xl",
				  "wrap": true,
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "Master Shifu is one of the main supporting characters of the Kung Fu Panda franchise. He is the current senior master of the Jade Palace and trainer of many kung fu warriors, including Tai Lung, the Furious Five, and the Dragon Warrior Po.",
				  "contents": []
				}
			  ]
			},
			"footer": {
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "button",
				  "action": {
					"type": "uri",
					"label": "ข้อมูล",
					"uri": "https://hmong.in.th/wiki/Master_Shifu"
				  },
				  "flex": 2,
				  "color": "#2DCF1EFF",
				  "style": "primary"
				}
			  ]
			}
		  },
		  {
			"type": "bubble",
			"hero": {
			  "type": "image",
			  "url": "https://i.pinimg.com/236x/f1/57/77/f15777adb8e82aa538d010a24aa12f25.jpg",
			  "size": "full",
			  "aspectRatio": "20:13",
			  "aspectMode": "cover",
			  "action": {
				"type": "uri",
				"label": "Line",
				"uri": "https://linecorp.com/"
			  }
			},
			"body": {
			  "type": "box",
			  "layout": "vertical",
			  "contents": [
				{
				  "type": "text",
				  "text": "โป",
				  "weight": "bold",
				  "size": "xl",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "Master Po Ping (known simply as Po, and originally named Lotus[1]) is the main protagonist of the Kung Fu Panda franchise. He is the adopted son of Mr. Ping and the biological son of Li Shan and Li's wife, as well as one of Master Shifu's students at the Jade Palace. He is also the foretold Dragon W",
				  "contents": []
				}
			  ]
			},
			"footer": {
			  "type": "box",
			  "layout": "vertical",
			  "flex": 0,
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "button",
				  "action": {
					"type": "uri",
					"label": "ข้อมูล",
					"uri": "https://hmong.in.th/wiki/Master_Shifu"
				  },
				  "color": "#2DCF1EFF",
				  "style": "primary"
				}
			  ]
			}
		  },
		  {
			"type": "bubble",
			"hero": {
			  "type": "image",
			  "url": "https://static.wikia.nocookie.net/kungfupanda/images/f/f9/KFP3-promo-monkey1.jpg/revision/latest?cb=20150726165704",
			  "size": "full",
			  "aspectRatio": "20:13",
			  "aspectMode": "cover",
			  "action": {
				"type": "uri",
				"label": "Line",
				"uri": "https://linecorp.com/"
			  }
			},
			"body": {
			  "type": "box",
			  "layout": "vertical",
			  "contents": [
				{
				  "type": "text",
				  "text": "ลิง",
				  "weight": "bold",
				  "size": "xl",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "Master Monkey is one of the supporting characters of the Kung Fu Panda franchise. He is a member of the Furious Five as well as one of Master Shifu's students at the Jade Palace. He is a master of the Monkey Style of kung fu.",
				  "contents": []
				}
			  ]
			},
			"footer": {
			  "type": "box",
			  "layout": "vertical",
			  "flex": 0,
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "button",
				  "action": {
					"type": "uri",
					"label": "ข้อมูล",
					"uri": "https://hmong.in.th/wiki/Master_Shifu"
				  },
				  "color": "#2DCF1EFF",
				  "style": "primary"
				}
			  ]
			}
		  },
		  {
			"type": "bubble",
			"hero": {
			  "type": "image",
			  "url": "https://static.wikia.nocookie.net/kungfupanda/images/e/ef/KFP3-promo-mantis1.jpg/revision/latest?cb=20150726231824",
			  "size": "full",
			  "aspectRatio": "20:13",
			  "aspectMode": "cover",
			  "action": {
				"type": "uri",
				"label": "Line",
				"uri": "https://linecorp.com/"
			  }
			},
			"body": {
			  "type": "box",
			  "layout": "vertical",
			  "contents": [
				{
				  "type": "text",
				  "text": "ตั๊กแตน",
				  "weight": "bold",
				  "size": "xl",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "Master Mantis is one of the supporting characters of the Kung Fu Panda franchise. He is a member of the Furious Five as well as one of Master Shifu's students at the Jade Palace. He is a master of the Mantis Style of kung fu.",
				  "contents": []
				}
			  ]
			},
			"footer": {
			  "type": "box",
			  "layout": "vertical",
			  "flex": 0,
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "button",
				  "action": {
					"type": "uri",
					"label": "ข้อมูล",
					"uri": "https://hmong.in.th/wiki/Master_Shifu"
				  },
				  "color": "#2DCF1EFF",
				  "style": "primary"
				}
			  ]
			}
		  },
		  {
			"type": "bubble",
			"hero": {
			  "type": "image",
			  "url": "https://static.wikia.nocookie.net/kungfupanda/images/d/de/KFP3-promo-crane1.jpg/revision/latest?cb=20150726165530",
			  "align": "start",
			  "size": "full",
			  "aspectRatio": "20:13",
			  "aspectMode": "fit",
			  "action": {
				"type": "uri",
				"label": "Line",
				"uri": "https://linecorp.com/"
			  }
			},
			"body": {
			  "type": "box",
			  "layout": "vertical",
			  "contents": [
				{
				  "type": "text",
				  "text": "กระเรียน",
				  "weight": "bold",
				  "size": "xl",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "Master Crane is one of the supporting characters of the Kung Fu Panda franchise. He is a member of the Furious Five as well as one of Master Shifu's students at the Jade Palace. He is a master of the Crane Style of kung fu.",
				  "contents": []
				}
			  ]
			},
			"footer": {
			  "type": "box",
			  "layout": "vertical",
			  "flex": 0,
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "button",
				  "action": {
					"type": "uri",
					"label": "ข้อมูล",
					"uri": "https://hmong.in.th/wiki/Master_Shifu"
				  },
				  "color": "#2DCF1EFF",
				  "style": "primary"
				}
			  ]
			}
		  },
		  {
			"type": "bubble",
			"hero": {
			  "type": "image",
			  "url": "https://static.wikia.nocookie.net/kungfupanda/images/d/db/KFP3-promo-viper.png/revision/latest?cb=20160223200500",
			  "align": "start",
			  "size": "full",
			  "aspectRatio": "20:13",
			  "aspectMode": "fit",
			  "action": {
				"type": "uri",
				"label": "Line",
				"uri": "https://linecorp.com/"
			  }
			},
			"body": {
			  "type": "box",
			  "layout": "vertical",
			  "contents": [
				{
				  "type": "text",
				  "text": "อสรพิษ",
				  "weight": "bold",
				  "size": "xl",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "Master Viper is one of the supporting characters of the Kung Fu Panda franchise. She is a member of the Furious Five and the daughter of Great Master Viper, as well as one of Master Shifu's students at the Jade Palace. She is a master of the Viper Style of kung fu.",
				  "contents": []
				}
			  ]
			},
			"footer": {
			  "type": "box",
			  "layout": "vertical",
			  "flex": 0,
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "button",
				  "action": {
					"type": "uri",
					"label": "ข้อมูล",
					"uri": "https://hmong.in.th/wiki/Master_Shifu"
				  },
				  "color": "#2DCF1EFF",
				  "style": "primary"
				}
			  ]
			}
		  }
		]
	  }`))
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return
	}
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("FlexWithJSON", flexContainer)).Do(); err != nil {
		log.Println("Error sending Flex Message:", err)
	}

}

func sendHelpMessage(event *linebot.Event, bot *linebot.Client) {

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("โปรดพิมพ์หนึ่งในรายการต่อไปนี้\ntext : บอทตอบทักทายเล็กน้อย \nquick : บอทแสดงปุ่ม quick reply ขึ้นมา \nbutton : บอทแสดงกล่องข้อความขึ้นมาพร้อมรูปภาพ \ncarousel : บอทแสดงข้อความกล่องแบบเลื่อนได้")).Do(); err != nil {
		log.Println(err)
	}
}

func sendHistoryMessage(event *linebot.Event, bot *linebot.Client) {

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("กำแพงถูกสร้างขึ้นเพื่อป้องกันการรุกรานขจากเผ่าภายนอก")).Do(); err != nil {
		log.Println(err)
	}
}

func sendFlexMessageUK(event *linebot.Event, bot *linebot.Client) {

	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(`{
		"type": "bubble",
		"hero": {
		  "type": "image",
		  "url": "https://traveldigg.com/wp-content/uploads/2016/05/Tower-Of-London-Images.jpg",
		  "size": "full",
		  "aspectRatio": "20:13",
		  "aspectMode": "cover",
		  "action": {
			"type": "uri",
			"label": "Line",
			"uri": "https://linecorp.com/"
		  }
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "ประเทศอังกฤษ The Tower Of London ",
			  "weight": "bold",
			  "size": "xl",
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "margin": "lg",
			  "contents": [
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "contents": [
					{
					  "type": "text",
					  "text": "ที่ตั้ง",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "flex": 1,
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "The Tower Of London, St Katharine's & Wapping, London EC3N 4AB",
					  "size": "sm",
					  "color": "#666666",
					  "flex": 5,
					  "wrap": true,
					  "contents": []
					}
				  ]
				}
			  ]
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "vertical",
		  "flex": 0,
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "button",
			  "action": {
				"type": "uri",
				"label": "ตั๋ว",
				"uri": "https://www.hrp.org.uk/tower-of-london/visit/tickets-and-prices/#gs.6k5yd2"
			  },
			  "height": "sm",
			  "style": "link"
			},
			{
			  "type": "button",
			  "action": {
				"type": "uri",
				"label": "ข้อมูล",
				"uri": "https://www.hrp.org.uk/tower-of-london/#gs.6k5wpp"
			  },
			  "height": "sm",
			  "style": "link"
			},
			{
			  "type": "spacer",
			  "size": "sm"
			}
		  ]
		}
	  }
	`))
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return
	}
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("FlexWithJSON", flexContainer)).Do(); err != nil {
		log.Println("Error sending Flex Message:", err)
	}
}

func sendFlexMessageItaly(event *linebot.Event, bot *linebot.Client) {

	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(`
	{
		"type": "bubble",
		"hero": {
		  "type": "image",
		  "url": "https://th.bing.com/th?id=OLC.tNrK9vQzZJJuhQ480x360&w=280&h=140&c=8&rs=1&qlt=90&o=6&pid=3.1&rm=2",
		  "size": "full",
		  "aspectRatio": "20:13",
		  "aspectMode": "cover",
		  "action": {
			"type": "uri",
			"label": "Line",
			"uri": "https://linecorp.com/"
		  }
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "ประเทศอิตาลี Colosseum",
			  "weight": "bold",
			  "size": "xl",
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "margin": "lg",
			  "contents": [
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "contents": [
					{
					  "type": "text",
					  "text": "ที่ตั้ง",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "flex": 1,
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "Colosseo Square - 00184 Rome (Rome)",
					  "weight": "regular",
					  "size": "xxs",
					  "color": "#666666",
					  "flex": 5,
					  "gravity": "bottom",
					  "wrap": true,
					  "contents": []
					}
				  ]
				}
			  ]
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "vertical",
		  "flex": 0,
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "button",
			  "action": {
				"type": "uri",
				"label": "ตั๋ว",
				"uri": "https://tourkrub.co/destinations/destination-colosseum"
			  },
			  "height": "sm",
			  "style": "link"
			},
			{
			  "type": "button",
			  "action": {
				"type": "uri",
				"label": "ข้อมูล",
				"uri": "https://happylongway.com/colosseum/"
			  },
			  "height": "sm",
			  "style": "link"
			},
			{
			  "type": "spacer",
			  "size": "sm"
			}
		  ]
		}
	  }
	`))
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return
	}
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("FlexWithJSON", flexContainer)).Do(); err != nil {
		log.Println("Error sending Flex Message:", err)
	}

}

func sendFlexMessageChina(event *linebot.Event, bot *linebot.Client) {

	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(`
	{
		"type": "bubble",
		"hero": {
		  "type": "image",
		  "url": "https://ik.imagekit.io/tvlk/blog/2019/03/Blue-Moon-Valley-1-1-800x531.jpg?tr=dpr-2,w-675",
		  "size": "full",
		  "aspectRatio": "20:13",
		  "aspectMode": "cover",
		  "action": {
			"type": "uri",
			"label": "Line",
			"uri": "https://linecorp.com/"
		  }
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "ประเทศจีน Blue Moon Valley",
			  "weight": "bold",
			  "size": "xl",
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "margin": "lg",
			  "contents": [
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "contents": [
					{
					  "type": "text",
					  "text": "ที่ตั้ง",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "flex": 1,
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "Blue Moon Valley, or Baishui River, is located nearby Yunshanping (cable car) and Ganhaizi, at the eastern foot of Jade Dragon Snow Mountain, approximately 30km (18.6 miles) north of Lijiang Ancient Town. The river was formed by the streams melted from the glaciers and snow of Jade Dragon Snow Mount",
					  "weight": "regular",
					  "size": "xxs",
					  "color": "#666666",
					  "flex": 5,
					  "gravity": "bottom",
					  "wrap": true,
					  "contents": []
					}
				  ]
				}
			  ]
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "vertical",
		  "flex": 0,
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "button",
			  "action": {
				"type": "uri",
				"label": "ตั๋ว",
				"uri": "https://www.chinadiscovery.com/yunnan/lijiang/blue-moon-valley.html"
			  },
			  "height": "sm",
			  "style": "link"
			},
			{
			  "type": "button",
			  "action": {
				"type": "uri",
				"label": "ข้อมูล",
				"uri": "https://www.travelchinaguide.com/attraction/yunnan/lijiang/white-water-river.htm"
			  },
			  "height": "sm",
			  "style": "link"
			},
			{
			  "type": "spacer",
			  "size": "sm"
			}
		  ]
		}
	  }

	`))
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return
	}
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("FlexWithJSON", flexContainer)).Do(); err != nil {
		log.Println("Error sending Flex Message:", err)
	}
}

func sendFlexMessageFrance(event *linebot.Event, bot *linebot.Client) {

	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(`
	{
		"type": "bubble",
		"hero": {
		  "type": "image",
		  "url": "https://th.bing.com/th/id/OIP.Ese6RWA_64j7cZIavbQA2QHaLH?w=616&h=924&rs=1&pid=ImgDetMain",
		  "size": "full",
		  "aspectRatio": "20:13",
		  "aspectMode": "cover",
		  "action": {
			"type": "uri",
			"label": "Line",
			"uri": "https://linecorp.com/"
		  }
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "ประเทศฝรั่งเศส โกตดาซูร์",
			  "weight": "bold",
			  "size": "xl",
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "margin": "lg",
			  "contents": [
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "contents": [
					{
					  "type": "text",
					  "text": "ที่ตั้ง",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "flex": 1,
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "Provence-Alpes-Côte d'Azur, also known as Région Sud, is one of the eighteen administrative regions of France,",
					  "weight": "regular",
					  "size": "xxs",
					  "color": "#666666",
					  "flex": 5,
					  "gravity": "bottom",
					  "wrap": true,
					  "contents": []
					}
				  ]
				}
			  ]
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "vertical",
		  "flex": 0,
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "button",
			  "action": {
				"type": "uri",
				"label": "ตั๋ว",
				"uri": "https://www.skyscanner.co.th/flights/flights-to-region/44293386/cheap-flights-to-provence-alpes-cote-d-azur.html"
			  },
			  "height": "sm",
			  "style": "link"
			},
			{
			  "type": "button",
			  "action": {
				"type": "uri",
				"label": "ข้อมูล",
				"uri": "https://leelawadeeholiday.com/default.asp?content=contentdetail&id=21101"
			  },
			  "height": "sm",
			  "style": "link"
			},
			{
			  "type": "spacer",
			  "size": "sm"
			}
		  ]
		}
	  }
	`))
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return
	}
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("FlexWithJSON", flexContainer)).Do(); err != nil {
		log.Println("Error sending Flex Message:", err)
	}
}

func sendUnknownMessage(event *linebot.Event, bot *linebot.Client) {

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ไม่พบคำที่ต้องการ กรุณาพิมพ์ \"ช่วยเหลือ\" ")).Do(); err != nil {
		log.Println(err)
	}

}

func handleError(err error, w http.ResponseWriter) {
	if err == linebot.ErrInvalidSignature {
		w.WriteHeader(400)
	} else {
		w.WriteHeader(500)
	}
}
