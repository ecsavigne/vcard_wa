package main

import (
	"bytes"
	"fmt"

	// vc "github.com/ecsavigne/vcard_wa"
	vc "github.com/ecsavigne/vcard_wa"
)

func main() {
	text := "BEGIN:VCARD\nVERSION:3.0\nN:família);Agente comunitario de saúde (clinica;;;\nFN:Agente comunitario de saúde (clinica família)\nX-WA-BIZ-NAME:Agente De Saúde Leonardo\nX-WA-BIZ-DESCRIPTION:Agente de Saúde Leonardo da Equipe Herminia - Clínica da Família Erivaldo Fernandes Nóbrega \n\nHorário de funcionamento da clínica: Segunda a sexta das 07h às 18h\n\nMeu horário de atendimento é das 07 às 16h\nORG:Agente De Saúde Leonardo;\nitem2.TEL;waid=5521977178948:+55 21 97717-8948\nitem2.X-ABLabel:Celular\nEND:VCARD"

	m := vc.ParseVCard(bytes.NewReader([]byte(text)))
	fmt.Println(m.GetPartRecord())
	fmt.Println("Tags:")
	fmt.Println(m.GetTags())
	fmt.Println("GetTagValue:")
	for _, k := range m.GetTags() {
		fmt.Printf("key: '%s' = value: '%s\n", k, m.GetTagValue(k))
	}

	fmt.Println("GetTagContentValues:")
	fmt.Printf("key: '%s' = value: '%s\n", "X-WA-BIZ-NAME", m.GetTagContentValues("X-WA-BIZ-NAME"))
	fmt.Printf("key: '%s' = value: '%s\n", "TEL", m.GetTagContentValues("TEL"))
}
