package sampledata

var SocialGraph map[string]map[string]string = map[string]map[string]string{
	"@bongolpoc": {
		"first_name": "Joselito",
		"last_name":  "Olpoc",
		"following":  "",
	},
	"@joaquin": {
		"first_name": "Joaquin",
		"last_name":  "Gonzales",
		"following":  "@chums,@jobenilagan",
	},
	"@chums": {
		"first_name": "Matthew",
		"last_name":  "Uy",
		"following":  "@bongolpoc,@miketan,@rudyang,@joeilagan",
	},
	"@jobenilagan": {
		"first_name": "Joben",
		"last_name":  "Ilagan",
		"following":  "@eeebeee,@joeilagan,@chums,@joaquin",
	},
	"@joeilagan": {
		"first_name": "Joe",
		"last_name":  "Ilagan",
		"following":  "@eeebeee,@jobenilagan,@chums",
	},
	"@eeebeee": {
		"first_name": "EB",
		"last_name":  "Ilagan",
		"following":  "@jobenilagan,@joeilagan",
	},
}

var Board [][]string = [][]string{
	{"X", "X", "O"},
	{"O", "X", "O"},
	{"", "O", "X"},
}

var RouteMap map[string]map[string]int = map[string]map[string]int{
	"upd,admu": {
		"travel_time_mins": 10,
	},
	"admu,dlsu": {
		"travel_time_mins": 35,
	},
	"dlsu,upd": {
		"travel_time_mins": 55,
	},
}
