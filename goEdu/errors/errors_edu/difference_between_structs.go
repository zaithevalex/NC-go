package main

import "fmt"

//type Interface interface {
//	method() int
//}

type Struct struct {
	field int
}

//func (s Struct) Error() int {
//	s.field = 2
//	return 0
//}

func (Struct) method() int {
	return 0
}

//func (s *Struct) Error() string {
//	return fmt.Sprintf("ERROR #%d\n", s.field)
//}

func (s Struct) Error() string {
	return fmt.Sprintf("ERROR #%d", s.field)
}

func main() {
	struct1 := Struct{field: 1}
	fmt.Println("struct1:", struct1)
	struct1.Error()
	structSave := &struct1
	structSave1 := &structSave
	fmt.Println("&struct1:", structSave) // structSave.Error() <=> (*structSave).Error() by pointer
	// structSave.Error() - by meaning
	fmt.Println("&&struct1:", structSave1) // (*(&structSave)).Error(
	//(*structSave).Error() // (GO) struct1.Error() <=> (C++) struct1->Error()
	//M(struct1)
}
