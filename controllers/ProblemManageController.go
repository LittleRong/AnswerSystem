package controllers

import (
    "github.com/astaxie/beego"
)

type ProblemManageController struct{
	beego.Controller
}

func (this *ProblemManageController) ProblemManageInit(){
	this.TplName = "manage/problem_manage.html"
}

func (this *ProblemManageController) ProblemManage(){
	//offset,_ := this.GetInt("offset")
	//limit,_ := this.GetInt("limit")
	//
	//problem_list := problem.GetProblemListByOffstAndLimit(offset,limit)
	//
	////problem_data,page_num
	//beego.Info("======problem_list=====",problem_list)
	//var result map[string]interface{}
	//result = make(map[string]interface{})
	//result["problem_data"] = problem_list
	//result["page_num"] = offset
	//this.Data["json"] = result
	//this.ServeJSON()
	//return

}

func (this *ProblemManageController) ChangeProblem(){
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

func (this *ProblemManageController) AddProblem(){
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

func (this *ProblemManageController) DeleteProblem(){
	//delete_id,_ := this.GetInt("delete_id")
	//r :=problem.DeleteProblemById(delete_id)
	//var result map[string]interface{}
	//result = make(map[string]interface{})
	//result["result"] = r
	//this.Data["json"] = result
	//this.ServeJSON()
	//return
}
