package main

// import (
// 	"bytes"
// 	"fmt"

// 	// v1 "github.com/ecsavigne/vcard-wa/v1"
// )

// func main() {
// 	text := "BEGIN:VCARD\nVERSION:3.0\nN:;PeinadoBrasil;;;\nFN:PeinadoBrasil\nX-WA-BIZ-NAME:Tiarinha Mello\nX-WA-BIZ-DESCRIPTION:Afeto | Cuidado  |Beleza | Ancestralidade | Empoderamento.\n\nNosso salão é especializado em cuidados com o cabelo crespo. \n\nSe você procura um ambiente profissional, descontraído, acolhedor, com serviço de excelência e voltado exclusivamente para pessoas negras, você está no lugar certo. 📍💁🏾‍♀️\n\nTiarinha Mello e toda sua equipe estão aguardando por você! 🤎✊🏿\nVenha viver essa experiência!!✨ ✨✨\n\nAgendamentos somente aqui no Whatsapp.😉\n\n😍Conheça o meu trabalho \ninstagram.com/tiarinhamello\nTEL;type=CELL;type=VOICE;waid=5521996500815:+55 21 99650-0815\nEND:VCARD"

// 	m := v1.ParseVCard(bytes.NewReader([]byte(text)))
// 	fmt.Println(m.GetPartRecord())
// 	fmt.Println("Tags:")
// 	fmt.Println(m.GetTags())
// 	fmt.Println("GetTagValue:")
// 	for _, k := range m.GetTags() {
// 		fmt.Printf("key: '%s' = value: '%s\n", k, m.GetTagValue(k))
// 	}

// 	fmt.Println("GetTagContentValues:")
// 	fmt.Printf("key: '%s' = value: '%s\n", "X-WA-BIZ-NAME", m.GetTagContentValues("X-WA-BIZ-NAME"))
// 	fmt.Printf("key: '%s' = value: '%s\n", "TEL", m.GetTagContentValues("TEL"))
// }
