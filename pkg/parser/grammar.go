package parser

type SMTPGrammar struct {
	Hello *HELOCommand `parser:"@@"`
}

type HELOCommand struct {
	FQDN *string `parser:"HELO ws @FQDN"`
}
