package main

import (
	"fmt"
	"log"
)

func (cm *Peg) AddParam(key, value string) {
	cm.Params[key] = value
}

func (cm *Peg) AddCropSubParam(key, subKey, value string) {
	cm.CropParams[subKey] = value
	cm.Params[key] = cm.CropParams
}

func (cm *Peg) SkipParam(text string) {
	log.Printf("SkipParam ======== %s", text)
}

func main() {

	testP := []string{
		"?format=png&progressive=false&width=10&height=10&fit=clip&scale=1.0&crop(x10,y10,w10,h10)&reverse=flip&quality=10exif=true",
		"?crop(h-40,w-30,x01,y-100)&q=9.000&fit=max&scale=aaaaaaa1.0&&reverse=flip&exif=true",
		//"?crop(w30,h50)",
		"?quality=09&&&&&&quality=8&q=9.000&format=png&format=11&quality=100aa&progressive=true&progressive=19&width=100&width=true&width=hgoe&crop=(w100,h100)&crop(w30,h50)",
		"?q=100%&aaa=pp",
		"?q=100",
		"?quality=1000",
		"?quality=100",
		"?quality=09&&&&&&quality=8&q=9.000&format=png&format=11&quality=100aa&progressive=true&progressive=19&width=100&width=true&width=hgoe&crop=(100,100)",
		""}

	p := &Peg{}
	log.Println("==========================================")
	log.Println("QueryParamater == ", testP[0])
	log.Println("==========================================")
	p.Buffer = testP[0]

	p.Init()
	p.Params = map[string]interface{}{}
	p.CropParams = map[string]interface{}{}
	err := p.Parse()
	if err != nil {
		fmt.Printf("Oops, Error! cause: %v\n", err)
	}

	p.Execute()

	log.Println(p.Params)
	crop, ok := p.Params["crop"].(map[string]interface{})

	if ok {
		log.Println(crop["height"])
	}


	fmt.Println("Success!!!")
}
