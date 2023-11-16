package main

type StatusClient interface{
	Alerts() ([]string, error)
}
type Service struct {
	statusClient StatusClient
	store Store
}

func (s *Service) BackgroundService(){
	// s,err:=s.statusClient.Alerts()
	// if err!=nil{
	// 	fmt.Er
	// }
	// for _,v:=range s{

	//}
}

func Alerts() ([]string, error){
	return []string{"iuh","12386","1397"},nil
}