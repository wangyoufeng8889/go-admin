package aliyunjobs

var DtuToBms map[string]string
var BmsToDtu map[string]string

func init()  {
	DtuToBms = make(map[string]string)
	BmsToDtu = make(map[string]string)
}