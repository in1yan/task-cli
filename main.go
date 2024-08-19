package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)
type Tasks struct{
	Tasks []Task `json:"tasks"`
}
type Task struct{
	Id int `json:"id"`
	Desc string `json:"description"`
	Status string `json:"status"`
	At string `json:"created-at"`
	UpAt string `json:"updated-at"`
}

func (p *Tasks) List(status string){
	for _,n := range p.Tasks{
			if status == ""||status==n.Status{
				fmt.Printf("ID : %v\nDescription : %v\nStatus : %v\nCreated At : %v\nLast Updated : %v\n",n.Id,n.Desc,n.Status,n.At,n.UpAt)
				fmt.Println("-------------------------------------------------------------")
			}
		}
}
func (p *Tasks) Add(desc string){
	task := Task{
		Id: len(p.Tasks)+1,
		Desc: desc,
		Status: "todo",
		At: time.Now().Format(time.RFC3339),
		UpAt: time.Now().Format(time.RFC3339),
	}	
	fmt.Println("Added!")
	fmt.Printf("ID : %v\nDescription : %v\nStatus : %v\nCreated At : %v\n",task.Id,task.Desc,task.Status,task.At)
	fmt.Println("-------------------------------------------------------------")
	p.Tasks = append(p.Tasks,task)
	p.Save("tasks.json")
}
func (p *Tasks) Delete(id int){
	index := 0
	for i := range p.Tasks{
		if p.Tasks[i].Id == id{
			index = i
		}
	}
	p.Tasks = append(p.Tasks[:index], p.Tasks[index+1:]...)
	fmt.Printf("Deleted task (Id #%v)\n",id)
	p.Save("tasks.json")
}
func (p *Tasks) Update(Id int, desc string,arg string){
	for i := range p.Tasks{
		if p.Tasks[i].Id == Id{
			if arg == "update"{
			p.Tasks[i].Desc = desc
			p.Tasks[i].UpAt = time.Now().Format(time.RFC3339)
			}else if arg == "mark-in-progress"{
			p.Tasks[i].Status = "in-progress"
			p.Tasks[i].UpAt = time.Now().Format(time.RFC3339)
			}else if arg == "mark-done"{
			p.Tasks[i].Status = "done"
			p.Tasks[i].UpAt = time.Now().Format(time.RFC3339)
			}
		}
	}
	fmt.Printf("Updated task (Id #%v)\n",Id)
	p.Save("tasks.json")
}
func (p *Tasks) Save(filename string){
	jsonData,err := json.MarshalIndent(p,"","    ")
	if err != nil{
		fmt.Printf("Error marshelling data")
	}
	file,err := os.OpenFile(filename,os.O_CREATE|os.O_RDWR|os.O_TRUNC,0777)
	if err != nil{
		fmt.Printf("Error opening file\n")
	}
	defer file.Close()
	_,err = file.Write(jsonData)
	if err != nil{
		fmt.Printf("Error writing file\n")
	}
}

func main(){
	args := os.Args[1:]
	if len(args) <1{
		return
	}
	file,err := os.OpenFile("tasks.json",os.O_CREATE|os.O_RDWR,0777)
	if err !=nil{
		file,err = os.Create("tasks.json")
	}
	defer file.Close()
	var tasks Tasks
	byteVal,_ := ioutil.ReadAll(file)
	json.Unmarshal(byteVal,&tasks)
	switch args[0]{
	case "list":{
		status := ""
		if len(args)>1{
			status = args[1]
		}
		tasks.List(status)
		
	}
	case "add":{
		tasks.Add(args[1])

	}
	case "update":{
		id,_ := strconv.Atoi(args[1])
		tasks.Update(id,args[2],args[0])

	}
	case "delete":{
		id,_ := strconv.Atoi(args[1])
		tasks.Delete(id)
	}
	case "mark-in-progress":{
		id,_ := strconv.Atoi(args[1])
		tasks.Update(id,"",args[0])
	}
	case "mark-done":{
		id,_ := strconv.Atoi(args[1])
		tasks.Update(id,"",args[0])
	}
	}
}
