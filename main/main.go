package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//命令行mac生成.exe文件: CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
func main()  {

	srcDir :="J:/超星汇雅/150-03"
	files,_:=ioutil.ReadDir(srcDir);

	for _,file:= range files{

		fmt.Println(file.Name())


		srcDir2 :=srcDir+"/"+file.Name()
		files2,_:=ioutil.ReadDir(srcDir2)
		for _,file1:=range files2{
			if(file1.IsDir()&&strings.Contains(file1.Name(),"_")){

					oldPath:=srcDir2+"/"+file1.Name();
					strNames:= strings.Split(file1.Name(),"_")
					dirStr:=srcDir2
					for i:=0;i<len(strNames)-1;i++{
						dirStr= dirStr+"/"+strNames[i]
					}
					os.MkdirAll(dirStr,0777)

					fmt.Println("dirStr/"+dirStr)
					newPath:= dirStr+"/"+strNames[len(strNames)-1]
					fmt.Println("oldPath:"+oldPath)
					fmt.Println("newPath:"+newPath)
					// go 里面os.Rename方法 newPath应该是操作前不存在的路径  不会替换会报错
					err:=os.Rename(oldPath,newPath)
					if err!=nil{
						fmt.Println(err.Error())
						return
					}
			}
		}


		//if(strings.Contains(file.Name(),"_")){

		//}


	}


}