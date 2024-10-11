package bot

//var cmd_list_type = []string {
//	"cmd_list_only",
//	"cmd_list_no_white_space",
//	"cmd_list_white_space",
//}

// 명령어"만" 쳐야함
var cmd_list_only = map[string]string{}

// 공백 문자 허용 x, 다른 거 있으면 상관없
var cmd_list_no_white_space = map[string]string{}

// 공백문자를 허용 및 다른 거 있어도 상관없
var cmd_list_white_space = map[string]string{}
