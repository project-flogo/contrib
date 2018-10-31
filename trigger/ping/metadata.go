package ping


type Settings struct {
	Port int `md:"port,required"`
	Version string `md:"version"`
}

type PingResponse struct {
	Version        string
	Appversion     string
	Appdescription string
}