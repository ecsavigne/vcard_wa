package vcard

var tags = []string{
	"N", "FN", "F", "X-WA-BIZ-NAME",
	"END", "BEGIN", "VERSION", "TEL",
	"X-WA-BIZ-DESCRIPTION", "TEL", "X-ABLabel", "ORG",
}

var replacer = []string{
	"\nitem1.", "\n",
	"\nitem2.", "\n",
	"\nitem3.", "\n",
}
