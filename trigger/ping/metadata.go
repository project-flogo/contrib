package ping


type Settings struct {
	Port int `md:"port,required"`
}

type PingResponse struct {
	Version        string
	Appversion     string
	Appdescription string
}