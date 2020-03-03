package controllers

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/spf13/viper"
	"github.com/tealeg/xlsx"

	eventProto "service/protoc/eventManage"
	proto "service/protoc/problemManage"
	"web/common"
)

type ProblemManageController struct {
	beego.Controller
}

// @Title 获得题目管理页面
// @Description 获得题目管理页面
// @Success 200 {}
// @router / [get]
func (this *ProblemManageController) ProblemManageInit() {
	this.TplName = "manage/problem_manage.html"
}

// @Title 获取题目列表
// @Description 获取题目列表
// @Success 200 {}
// @Param   offset   query   string  true       "页码"
// @Param   limit query   string  true       "一页展示数量"
// @router /all [get]
func (this *ProblemManageController) ProblemManage() {
	offset, _ := this.GetInt32("offset")
	limit, _ := this.GetInt32("limit")
	//获取用户信息
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	}
	userId := userSession.(int64)

	//call the userManage method
	problemManage, ctx := common.InitProblemManage(this.CruSession)
	req := proto.GetProblemListReq{Offset: offset, Limit: limit, ManageId: userId}
	rsp, err := problemManage.GetProblemListByOffstAndLimit(ctx, &req)
	if err != nil {
		beego.Info("======ProblemManage=====", rsp.ProblemList, "-------err--------", err)
	}
	beego.Info("======ProblemManage=====", rsp.ProblemList)
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

// @Title 获得题目上传页面
// @Description 获得题目上传页面
// @Success 200 {}
// @router /upload [get]
func (this *ProblemManageController) ProblemUploadInit() {
	new_event_id := this.GetSession("new_event_id")
	if new_event_id == nil { //未设置
		this.Ctx.Redirect(302, "/manage/event_insert_init")
		return
	}
	this.TplName = "manage/problem_upload.html"
}


// @Title 题目通过excel批量上传
// @Description 题目通过excel批量上传
// @Success 200 {}
// @Param   uploadname   formData   file  true       "上传文件"
// @router /upload [post]
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
	problemManage, ctx := common.InitProblemManage(this.CruSession)
	req := proto.GetEndProblemIdReq{}
	rsp, err := problemManage.GetEndProblemId(ctx, &req)
	if err != nil {
		beego.Info("-------err--------", err)
	}
	max_problem_id := rsp.EndId
	var first_problem_id int64
	first_problem_id = 1
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
			if (problem_type == viper.GetInt("enum.problemType.singleType")) {
				//单选题
				col_num := len(row.Cells)
				var problem_option_array []map[string]string
				var problem_answer_map map[string]string
				problem_answer_map = make(map[string]string)
				for i := 4; i < col_num; i++ {
					var option map[string]string
					option = make(map[string]string)
					tmp := max_problem_id*20 + int64(i)
					option["q_id"] = strconv.FormatInt(tmp, 10)
					option["content"] = row.Cells[i].String()
					abcd := string(95 - 4 + i)
					answer := strings.ToLower(row.Cells[3].String())
					if (answer == abcd) {
						tmp := max_problem_id*20 + int64(i)
						problem_answer_map["q_id"] = strconv.FormatInt(tmp, 10)
						problem_answer_map["content"] = row.Cells[i].String()
					}
					problem_option_array = append(problem_option_array, option)
				}
				problem_answer_byte, _ := json.Marshal(problem_answer_map)
				problem_option_byte, _ := json.Marshal(problem_option_array)
				problem_answer = string(problem_answer_byte)
				problem_option = string(problem_option_byte)
			} else if (problem_type == viper.GetInt("enum.problemType.multipleType")) {
				//单选题
				col_num := len(row.Cells)
				var problem_option_array []map[string]string
				var problem_answer_array []map[string]string
				for i := 4; i < col_num; i++ {
					var option map[string]string
					option = make(map[string]string)
					tmp := max_problem_id*20 + int64(i)
					option["q_id"] = strconv.FormatInt(tmp, 10)
					option["content"] = row.Cells[i].String()
					abcd := string(95 - 4 + i)
					answer := strings.ToLower(row.Cells[3].String())
					if (strings.Contains(answer, abcd)) {
						var problem_answer_map map[string]string
						problem_answer_map = make(map[string]string)
						tmp := max_problem_id*20 + int64(i)
						problem_answer_map["q_id"] = strconv.FormatInt(tmp, 10)
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

			//插入problem表
			req := proto.ProblemMesssage{ProblemContent: problem_content, ProblemType: int32(problem_type), ProblemClass: problem_class, ProblemAnswer: problem_answer, ProblemOption: problem_option}
			rsp, err := problemManage.AddProblem(ctx, &req)
			if err != nil {
				beego.Info("-------err--------", err)
			}

			max_problem_id = rsp.ProblemId
			//插入event_problem表
			eventManage, ctx := common.InitEventManage(this.CruSession)
			eventReq := eventProto.AddEventProblemReq{EventId: int64(event_id), ProblemId: max_problem_id}
			ep, err := eventManage.AddEventProblem(ctx, &eventReq)
			if err != nil {
				beego.Info("-------err--------", err, ep)
			}

			if (i == 1) {
				first_problem_id = max_problem_id
			}

		}

		//接着进行查询
		req := proto.GetNewProblemByTypeReq{FirstProblemId: first_problem_id, ProblemType: viper.GetInt32("enum.problemType.singleType")}
		single_arr, _ := problemManage.GetNewProblemByType(ctx, &req)
		req = proto.GetNewProblemByTypeReq{FirstProblemId: first_problem_id, ProblemType: viper.GetInt32("enum.problemType.multipleType")}
		multi_arr, _ := problemManage.GetNewProblemByType(ctx, &req)
		req = proto.GetNewProblemByTypeReq{FirstProblemId: first_problem_id, ProblemType: viper.GetInt32("enum.problemType.judgeType")}
		judge_arr, _ := problemManage.GetNewProblemByType(ctx, &req)
		req = proto.GetNewProblemByTypeReq{FirstProblemId: first_problem_id, ProblemType: viper.GetInt32("enum.problemType.fillType")}
		fill_arr, _ := problemManage.GetNewProblemByType(ctx, &req)
		beego.Info("************single_arr**************", single_arr)
		beego.Info("************multi_arr**************", multi_arr)
		beego.Info("************judge_arr**************", judge_arr)
		beego.Info("************fill_arr**************", fill_arr)
		//
		//single_json := Struct2Map(single_arr)
		//mutil_json := Struct2Map(multi_arr)
		//judge_json := Struct2Map(judge_arr)
		//fill_json := Struct2Map(fill_arr)

		var result map[string]interface{}
		result = make(map[string]interface{})
		result["single"] = single_arr.ProblemList
		result["multiple"] = multi_arr.ProblemList
		result["judge"] = judge_arr.ProblemList
		result["fill"] = fill_arr.ProblemList
		beego.Info("======AddProblem's id=====", result)
		this.Data["json"] = result
		this.ServeJSON()
		return

	}

}

//func Struct2Map(in []problem.Problem) []map[string]interface{} {
//	var result []map[string]interface{}
//	for _, obj := range in {
//		t := reflect.TypeOf(obj)
//		v := reflect.ValueOf(obj)
//
//		var data = make(map[string]interface{})
//		for i := 0; i < t.NumField(); i++ {
//			data[t.Field(i).Name] = v.Field(i).Interface()
//		}
//		result = append(result, data)
//	}
//
//	return result
//}
