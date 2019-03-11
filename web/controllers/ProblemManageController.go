package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/micro/go-micro"
	"github.com/tealeg/xlsx"
	"log"
	"reflect"
	"context"
	proto "service/protoc/problemManage" //proto文件放置路径
	"strconv"
	"strings"
	"web/models/event"
	"web/models/problem"
)
func (this *ProblemManageController) initProblemManage() proto.ProblemManageService{
	//调用服务
	service := micro.NewService(micro.Name("ProblemManage.client"))
	service.Init()

	//create new client
	return proto.NewProblemManageService("ProblemManage",service.Client())
}

type ProblemManageController struct {
	beego.Controller
}

func (this *ProblemManageController) ProblemManageInit() {
	this.TplName = "manage/problem_manage.html"
}

func (this *ProblemManageController) ProblemManage() {
	offset,_ := this.GetInt32("offset")
	limit,_ := this.GetInt32("limit")
	//获取用户信息
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	}
	userId := userSession.(int)

	//call the userManage method
	problemManage := this.initProblemManage()
	req := proto.GetProblemListReq{Offset:offset,Limit:limit,ManageId:int64(userId)}
	rsp, err := problemManage.GetProblemListByOffstAndLimit(context.TODO(),&req)
	if err!=nil{
		beego.Info("======ProblemManage=====", rsp.ProblemList,"-------err--------",err)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	result["problem_data"] = rsp.ProblemList
	result["page_num"] = offset
	this.Data["json"] = result
	this.ServeJSON()
	return

}

func (this *ProblemManageController) ChangeProblem() {
	//change_id,_ := this.GetInt("change_id")
	//problem_name := this.GetString("problem_name")
	//login_name := this.GetString("login_name")
	//problem_phone_number := this.GetString("problem_phone_number")
	//problem_job_number := this.GetString("problem_job_number")
	//problem_gender,_ := this.GetInt("problem_gender")
	//
	//r :=problem.UpdateProblemById (change_id,problem_name,login_name,problem_phone_number,problem_job_number,problem_gender)
	//var result map[string]interface{}
	//result = make(map[string]interface{})
	//result["result"] = r
	//this.Data["json"] = result
	//this.ServeJSON()
	//return
}

func (this *ProblemManageController) AddProblem() {
	//problem_name := this.GetString("problem_name")
	//login_name := this.GetString("login_name")
	//problem_phone_number := this.GetString("problem_phone_number")
	//problem_job_number := this.GetString("problem_job_number")
	//problem_gender,_ := this.GetInt("problem_gender")
	//
	//r :=problem.AddProblem (problem_name,login_name,problem_phone_number,problem_job_number,problem_gender)
	//var result map[string]interface{}
	//result = make(map[string]interface{})
	//result["result"] = r
	//this.Data["json"] = result
	//this.ServeJSON()
	//return

}

func (this *ProblemManageController) DeleteProblem() {
	//delete_id,_ := this.GetInt("delete_id")
	//r :=problem.DeleteProblemById(delete_id)
	//var result map[string]interface{}
	//result = make(map[string]interface{})
	//result["result"] = r
	//this.Data["json"] = result
	//this.ServeJSON()
	//return
}

func (this *ProblemManageController) ProblemUploadInit() {
	new_event_id := this.GetSession("new_event_id")
	if new_event_id == nil { //未设置
		this.Ctx.Redirect(302, "/manage/event_insert_init")
		return
	}
	this.TplName = "manage/problem_upload.html"
}

func (this *ProblemManageController) ProblemFileInsert() {
	new_event_id := this.GetSession("new_event_id")
	if new_event_id == nil { //未设置
		this.Ctx.Redirect(302, "/manage/event_insert_init")
		return
	}
	event_id := new_event_id.(int)

	f, h, err := this.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()

	//改名字!!
	this.SaveToFile("uploadname", "static/upload/"+h.Filename) // 保存位置在 static/upload

	//解析excel内容
	//格式：题目分类 题目类型 题目内容 答案 选项A 选项B 选项C 选项D
	excelFileName := "static/upload/" + h.Filename
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatal("文件内容错误", err)
	}

	//获取最大problem_id,这里还有问题
	max_problem_id := problem.GetEndProblemId()
	first_problem_id := 1
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			if (i == 0) {
				continue
			}
			max_problem_id++
			problem_class := row.Cells[0].String()
			problem_type, _ := row.Cells[1].Int()
			problem_content := row.Cells[2].String()

			var problem_answer string
			var problem_option string
			if (problem_type == 1) {
				//单选题
				col_num := len(row.Cells)
				var problem_option_array []map[string]string
				var problem_answer_map map[string]string
				problem_answer_map = make(map[string]string)
				for i := 4; i < col_num; i++ {
					var option map[string]string
					option = make(map[string]string)
					option["q_id"] = strconv.Itoa(max_problem_id*20 + i)
					option["content"] = row.Cells[i].String()
					abcd := string(95 - 4 + i)
					answer := strings.ToLower(row.Cells[3].String())
					if (answer == abcd) {
						problem_answer_map["q_id"] = strconv.Itoa(max_problem_id*20 + i)
						problem_answer_map["content"] = row.Cells[i].String()
					}
					problem_option_array = append(problem_option_array, option)
				}
				problem_answer_byte, _ := json.Marshal(problem_answer_map)
				problem_option_byte, _ := json.Marshal(problem_option_array)
				problem_answer = string(problem_answer_byte)
				problem_option = string(problem_option_byte)
			} else if (problem_type == 2) {
				//单选题
				col_num := len(row.Cells)
				var problem_option_array []map[string]string
				var problem_answer_array []map[string]string
				for i := 4; i < col_num; i++ {
					var option map[string]string
					option = make(map[string]string)
					option["q_id"] = strconv.Itoa(max_problem_id*20 + i)
					option["content"] = row.Cells[i].String()
					abcd := string(95 - 4 + i)
					answer := strings.ToLower(row.Cells[3].String())
					if (strings.Contains(answer, abcd)) {
						var problem_answer_map map[string]string
						problem_answer_map = make(map[string]string)
						problem_answer_map["q_id"] = strconv.Itoa(max_problem_id*20 + i)
						problem_answer_map["content"] = row.Cells[i].String()
						problem_answer_array = append(problem_answer_array, problem_answer_map)
					}
					problem_option_array = append(problem_option_array, option)
				}
				problem_answer_byte, _ := json.Marshal(problem_answer_array)
				problem_option_byte, _ := json.Marshal(problem_option_array)
				problem_answer = string(problem_answer_byte)
				problem_option = string(problem_option_byte)

			} else if (problem_type == 3) {
				if (row.Cells[3].String() == "是" || row.Cells[3].String() == "正确" || strings.ToLower(row.Cells[3].String()) == "yes" || row.Cells[3].String() == "1") {
					problem_answer = "true"
				} else {
					problem_answer = "false"
				}

			} else {
				problem_answer = row.Cells[3].String()
			}
			p := problem.Problem{Problem_content: problem_content, Problem_type: problem_type, Problem_class: problem_class, Problem_answer: problem_answer, Problem_option: problem_option}
			//插入problem表
			max_problem_id = problem.AddProblem(p)
			//插入event_problem表
			ep := event.EventProblem{Refer_event_id: event_id, Problem_id: max_problem_id}
			event.AddEventProblem(ep)

			if (i == 1) {
				first_problem_id = max_problem_id
			}

		}

		//接着进行查询
		single_arr := problem.GetNewProblemByType(first_problem_id, 1)
		multi_arr := problem.GetNewProblemByType(first_problem_id, 2)
		judge_arr := problem.GetNewProblemByType(first_problem_id, 3)
		fill_arr := problem.GetNewProblemByType(first_problem_id, 0)

		single_json := Struct2Map(single_arr)
		mutil_json := Struct2Map(multi_arr)
		judge_json := Struct2Map(judge_arr)
		fill_json := Struct2Map(fill_arr)

		var result map[string]interface{}
		result = make(map[string]interface{})
		result["single"] = single_json
		result["multiple"] = mutil_json
		result["judge"] = judge_json
		result["fill"] = fill_json
		beego.Info("======AddProblem's id=====", result)
		this.Data["json"] = result
		this.ServeJSON()
		return

	}

}

func Struct2Map(in []problem.Problem) []map[string]interface{} {
	var result []map[string]interface{}
	for _, obj := range in {
		t := reflect.TypeOf(obj)
		v := reflect.ValueOf(obj)

		var data = make(map[string]interface{})
		for i := 0; i < t.NumField(); i++ {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}
		result = append(result, data)
	}

	return result
}
