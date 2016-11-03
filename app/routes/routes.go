// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tLogin struct {}
var Login tLogin


func (_ tLogin) SignIn(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Login.SignIn", args).Url
}

func (_ tLogin) SignInHandler(
		name string,
		passwd string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "passwd", passwd)
	return revel.MainRouter.Reverse("Login.SignInHandler", args).Url
}

func (_ tLogin) SignUp(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Login.SignUp", args).Url
}

func (_ tLogin) SignUpHandler(
		name string,
		email string,
		passwd string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "passwd", passwd)
	return revel.MainRouter.Reverse("Login.SignUpHandler", args).Url
}


type tMain struct {}
var Main tMain


func (_ tMain) Main(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Main.Main", args).Url
}


type tAdmin struct {}
var Admin tAdmin


func (_ tAdmin) Main(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.Main", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


