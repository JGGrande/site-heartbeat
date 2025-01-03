package internal

type Site struct {
	Uuid string
	Nome string
	Url  string
}

type Log struct {
	SiteUuid string
	Texto    string
	Ativo    bool
	Data     string
}
