package exporter

import(
	"go_classes/go_classes/Interfaceclass"
	"fmt"
	"encoding/json"
	"encoding/xml"
	"os"
)

// Export feed
func Export(u interfaceclass.SocialMedia, filename string, format string) error {
	filename = filename+"."+format // fbdata, txt ==> fbdata.txt
	f, err:= os.OpenFile(filename, os.O_CREATE | os.O_WRONLY, 0755)
	if err != nil{
		return err
	}

	if format == "txt"{
		for _, fd := range u.Feed() {
			n, err := f.Write([]byte(fd + "\n"))
			if err != nil{
				return err
			}
			fmt.Printf("Wrote %d bytes\n", n)
		}
	}else if format == "json"{
		for n, fd := range u.Feed() {
			js := map[int]string{n:fd}
			json, _ := json.Marshal(js)
			n, err := f.Write([]byte(string(json)))
			if err != nil{
				return err
			}
			fmt.Printf("Wrote %d bytes\n", n)
		}
	}else if format == "xml"{
		for _, fd := range u.Feed(){
			out, _ := xml.MarshalIndent(fd,"","")
			n, err := f.Write([]byte(string(out)))
			if err != nil{
				return err
			}
			fmt.Printf("Wrote %d bytes\n", n)
		}
	}else if format == "yaml"{
	}
	return nil
