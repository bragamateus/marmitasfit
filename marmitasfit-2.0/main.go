package main

func main(){

	//assure server closing at end of execution
	defer StopServer()

	//call start server function
	StartServer()
}
