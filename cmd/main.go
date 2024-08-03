package main

import "fmt"


func minSwaps(nums []int) int {
    cnt := 0
    cnt2 := 0
    res := len(nums)
    for i:=0;i<len(nums);i++{
        if nums[i]==0{
            cnt++
        }
    }
    for i:=0;i<len(nums);i++{
        for j:=i;j<i+cnt;j++{
			if j == len(nums){
				break
			}
            if nums[j] == 1{
                cnt2++
            }
        }
        if cnt2<res{
			// fmt.Println("aa", cnt2, res)
            res = cnt2
        }
    }
    return res
}

func main(){
	nums := []int{0,1,1,1,0,0,1,1,0}
    fmt.Println(minSwaps(nums))
}
