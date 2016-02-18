package main
//import "fmt"

func CountIslands(grid[][] int)int{
counts:=0
	var i, j int
	for i=0;i<4;i++{
		for j=0;j<5;j++{
			if(grid[i][j]==1){
				if((i==0||grid[i-1][j]==0)&&(j==0||grid[i][j-1]==0)){
					counts++
				}
			}
		}
	}

return counts
}

/*func main() {

var grid=[][]int{{1,1,0,0,0},{1,1,0,0,0},{0,0,1,0,0},{0,0,0,1,1}}
fmt.Println(CountIslands(grid))


}*/