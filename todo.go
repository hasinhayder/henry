package main
import (
	"fmt"
	"strings"
	"gopkg.in/alecthomas/kingpin.v2"
)
var (
	/*debug   = kingpin.Flag("debug", "Enable debug mode.").Bool()
	timeout = kingpin.Flag("timeout", "Timeout waiting for ping.").Default("5s").OverrideDefaultFromEnvar("PING_TIMEOUT").Short('t').Duration()
	ip      = kingpin.Arg("ip", "IP address to ping.").Required().IP()
	count   = kingpin.Arg("count", "Number of packets to send").Int()*/

	add  = kingpin.Command("add", "Task Command")
	atask     = add.Arg("task", "Add Task").Strings()
	adue      = add.Flag("due", "Completion Time").Default("24h").Short('d').Duration()
	apriority = add.Flag("priority", "Task Priority").Short('p').Default("0").Int()

	list  = kingpin.Command("list", "Task Command").Default()
	lfrom      = add.Flag("from", "List From").Default("0").Short('f').Int()
	lto      = add.Flag("to", "List To").Short('t').Int()


)
func main(){
	kingpin.Version("0.0.1")
	switch kingpin.Parse() {
	case add.FullCommand():
		fmt.Print("Your Task Is : ", strings.Join(*atask, " "),", Priority: ",*apriority,", Duration: ",adue.Seconds())
	case list.FullCommand():
		fmt.Print("Listing ")
	}
}
