//You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
//
//Find two lines that together with the x-axis form a container, such that the container contains the most water.
//
//Return the maximum amount of water a container can store.
//
//Notice that you may not slant the container.
//Constraints:
//
//n == height.length
//2 <= n <= 105
//0 <= height[i] <= 104
//
//
//Runtime: 149 ms, faster than 34.95% of Go online submissions for Container With Most Water.
//Memory Usage: 9.2 MB, less than 43.30% of Go online submissions for Container With Most Water.
//
func maxArea(height []int) int {
    max:=0
    for i:=0;i<len(height);i+=1 {
        t_max:=0
        for j:=len(height)-1;j>i;j-=1 {
            if height[j]>t_max {
                tmp:= f_min(height[i],height[j])*(j-i)
                if tmp > max {
                    max= tmp
                }
                t_max=height[j]
            }
        }
    }
    return max;
}

func f_min (x int, y int)(min int) {
    if x>y {
        return y
    }
    return x
}



func maxArea(height []int) int {
    max:=0
    l:=0
    r:=len(height)-1
    for l<r {
        tmp:=f_min(height[l],height[r])*(r-l)
        if max<tmp{
            max=tmp
        }
        if height[l]<height[r] {
            l+=1
        }else{
            r-=1
        }
    }
    return max;
}

func f_min (x int, y int)(min int) {
    if x>y {
        return y
    }
    return x
}
