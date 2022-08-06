package interest


import (
	"net/http"
	"html/template"
	"github.com/dingtra/rundb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"strconv"
	"fmt"
)

type InterestedStruct struct{
	Details template.HTML
	Success bool
}


func (l *InterestedStruct) VerifyInterested (r *http.Request, id string){
	if strings.ToLower(r.Method) == "post" {
		body := strings.TrimSpace(r.FormValue("body"))

		fmt.Println(body)
	
		if len(body) > 1 {
			getcoll := rundb.Connect("app").Collection("register")
			
			theid, _:= primitive.ObjectIDFromHex(id)
			
			getcredentials := rundb.FindOne(getcoll, bson.M{"_id":theid})

			if len(getcredentials) > 0 {
				
				if getcredentials["interested"] == nil {
					if strings.TrimSpace(r.FormValue("oprate")) == "ax"{
						rundb.UpdateOne(getcoll, bson.M{"interested":"#"+body}, id)
					}
				}else {
				
					wrk := strings.Split(getcredentials["interested"].(string), "#")
					
					if len(wrk) > 1 {
						if strings.TrimSpace(r.FormValue("oprate")) == "ax"{
							var vrfy bool
							for _, k := range wrk{
								if k != ""{
									if body != k {
										vrfy = true
									}
								}
							}
	
							if vrfy == true {
								rundb.UpdateOne(getcoll, bson.M{"interested":getcredentials["interested"].(string)+"#"+body}, id)
	
								det := make(map[string]string)
								det["interested"] = getcredentials["interested"].(string)+"#"+body
								det["usersid"] = id
	
								l.Details = template.HTML( InterestedGraph(det, id))
								l.Success = true
							}
						}else {
							notdel := []string{}
							for _, k := range wrk{
								if k != ""{
									if body != k {
										notdel = append(notdel, "#"+k)
									}
								}
							}
	
							rundb.UpdateOne(getcoll, bson.M{"interested":strings.Join(notdel, "")}, id)

							det := make(map[string]string)
							det["interested"] = strings.Join(notdel, "")
							det["usersid"] = id
	
							l.Details = template.HTML( InterestedGraph(det, id))
							l.Success = true
						}
					}else{
						
						if strings.TrimSpace(r.FormValue("oprate")) == "ax"{
							rundb.UpdateOne(getcoll, bson.M{"interested":"#"+body}, id)

							det := make(map[string]string)
							det["interested"] = "#"+body
							det["usersid"] = id
	
							l.Details = template.HTML( InterestedGraph(det, id))
							l.Success = true
						}
						fmt.Println("in here")
					}

					
					
				}
			}
			
		}
	}
}


func GetInterests() []string {
    interests := []string{"&#129302 technology", "&#128200 trading", "&#128202 stocks", "&#128211 books", "&#128177 finance", "&#127971 education", "&#128176 cryptocurrency", " NFT"}

    return interests
}


func InterestedGraph(details map[string]string, id string) string {
	generate := []string{}
	getint := make(map[string]bool)

	if details["usersid"] == id {
		if details["interested"] != ""{
			//interested fieldset
			generate = append(generate, "<div id='bx009x'  style='padding-bottom:15px;'><fieldset><legend style='width:fit-content;margin:auto;'>interested areas</legend>")
			
			s := strings.Split(details["interested"], "#")

			for _, item := range s {
				if item != ""{
					getint[item] = true
				}
			}

			grabtruthvals := []string{}

			for i, item := range GetInterests() {
				if getint[strings.Split(item, " ")[1]] == true {
					grabtruthvals = append( grabtruthvals, "<div class='interestedjk78'>") 
					grabtruthvals = append( grabtruthvals, "<div style='border:1px solid #ccc;border-radius:10px;' id='xybprc' data-name='"+strings.Split(item, " ")[1]+"' data-ind='"+strconv.Itoa(i)+"'>")
					grabtruthvals = append( grabtruthvals, "<span style='pointer-events:none;'><div id='bpaste"+strconv.Itoa(i)+"'>"+item+"</div></span>")
					//open
					grabtruthvals = append( grabtruthvals, "<div id='framedelrem"+strconv.Itoa(i)+"' style='margin-top:5px;display:none;'>")
					grabtruthvals = append( grabtruthvals, "<span style='padding-right:10px;border:white;' id='bprc' data-name='"+strings.Split(item, " ")[1]+"' data-ind='"+strconv.Itoa(i)+"' data-operate='dx'><button style='pointer-events:none;background:white;border:1px solid #ccc;'>remove</button></span>")
					grabtruthvals = append( grabtruthvals, "<span><button style='background:white;border:1px solid #ccc;'>explore</button></span>")
					grabtruthvals = append( grabtruthvals,"</div>")
					//close
					grabtruthvals = append( grabtruthvals, "</div>")
					grabtruthvals = append( grabtruthvals, "</div>")
				}
			}
			
			generate = append(generate, "<div style='width:fit-content;margin:auto;' class='show-interested-wrap'>"+strings.Join(grabtruthvals, "")+"</div>")

			generate = append(generate, "</fieldset></div>")
			//interested fieldset
		}

		//suggested fieldset
		capture := []string{}
		generate = append(generate, "<fieldset><legend style='width:fit-content;margin:auto;'>suggested</legend>")
		for i, item := range GetInterests() {
			if getint[strings.Split(item, " ")[1]] == true {
				capture= append(capture, "<div class='interestedjk78'>") 
				capture= append(capture, "<div style='border:1px solid green;' id='bprc' data-name='"+strings.Split(item, " ")[1]+"' data-ind='"+strconv.Itoa(i)+"'  data-operate='ax'>")
				capture= append(capture, "<span style='pointer-events:none;'><div id='bpaste"+strconv.Itoa(i)+"'>"+item+"</div></span>")
				capture= append(capture, "</div>")
				capture= append(capture, "</div>")
			}else{
				capture= append(capture, "<div class='interestedjk78'>")
				capture= append(capture, "<div id='bprc' data-name='"+strings.Split(item, " ")[1]+"' data-ind='"+strconv.Itoa(i)+"'  data-operate='ax'>")
				capture= append(capture, "<span style='pointer-events:none;'><div id='bpaste"+strconv.Itoa(i)+"'>"+item+"</div></span>")
				capture= append(capture, "</div>")
				capture= append(capture, "</div>")
			}
		}
		generate = append(generate, "<div style='width:fit-content;margin:auto;' class='show-interested-wrap'>"+strings.Join(capture, "")+"</div>")
		generate = append(generate, "</fieldset>")
		//suggested fieldset

	}

	return strings.Join(generate, "")
	
}