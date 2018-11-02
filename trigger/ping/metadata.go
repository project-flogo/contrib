package ping


type Settings struct {
	Port 		int 	`md:"port,required"`
	Version 	string 	`md:"version"`
	AppVersion 	string 	`md:"appversion"`
	AppDescription 	string 	`md:"appdescription"`
}

type MemoryStats struct{
	NumGoroutine 	int
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObjects	uint64
	NumGC		uint32
}

func PrintMemUsage() string{
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	var t MemoryStats

	t.NumGoroutine = runtime.NumGoroutine()

	// Misc memory stats
	t.Alloc = rtm.Alloc
	t.TotalAlloc = rtm.TotalAlloc
	t.Sys = rtm.Sys
	t.Mallocs = rtm.Mallocs
	t.Frees = rtm.Frees

	// Live objects = Mallocs - Frees
	t.LiveObjects = t.Mallocs - t.Frees

	//GC stats
	t.NumGC = rtm.NumGC

	result, _ := json.Marshal(t)
	return string(result)
}

func Valid(token string) bool{
	matched, _ := regexp.MatchString("^Bearer (.*)", token)
	return matched
}
