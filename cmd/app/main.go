package main

func main() {
	s:=server.New()

	if err:=s.Run();err!=nil{
		log.Fatal(err)
	}
}
