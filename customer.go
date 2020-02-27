package main;
import (
	"fmt";
	"time";
)

type Customer struct {
 name string
 old string
 sdt string
 address string
}
type geometry interface {
 area() int
}
type reg struct{
 w int
 h int
}
func (r reg) area() int{
 return r.w * r.h 
}
func s(g geometry){
 fmt.Println(g.area());
}
func goroutin2(){
 fmt.Println("hello i'am goroutine2");
}
func goroutin(){
 fmt.Println("hello i'm goroutine");
}
func main (){
go goroutin();
go goroutin2();
time.Sleep(1 * time.Second);
 fmt.Println("xinchao");
 vt := Customer{
  name: "minh",
  old:"20",
  sdt:"0971725797",
  address:"okay", 
 };
fmt.Println(vt);
hv1 :=  reg{
 w : 5,
 h: 6,
};
 s(hv1);
}
