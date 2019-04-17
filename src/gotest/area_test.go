package gotest

import "testing"

//功能测试
func TestArea(t *testing.T) {
	m := Area(20, 30)
	if m != 600 {
		t.Error("测试失败")
	}
}

//并发
func BenchmarkArea(b *testing.B) {
	for i := 1; i < 1000; i++ {
        Area(i,i + 1)
	}
}
//调用主程序
func TestMain(m *testing.M)  {
    m.Run()
}
//分层调用
func TestWork(t *testing.T){
	t.Run("add",TestAdd)
	t.Run("del",TestDel)
	t.Run("get",TestGetNum)
}
func TestAdd(t *testing.T) {

}

func TestDel(t *testing.T) {

}

func TestGetNum(t *testing.T) {

}