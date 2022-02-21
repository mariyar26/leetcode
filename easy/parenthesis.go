//
//Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//An input string is valid if:
//
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
//
//
//Example 1:
//
//Input: s = "()"
//Output: trueRuntime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
//Memory Usage: 2 MB, less than 94.21% of Go online submissions for Valid Parentheses.
//



func isValid(s string) bool {
    roundb:=0
    squareb:=0
    curlyb:=0
    type stack []int
    st:=make(stack,0)
    for _,c:=range(s) {
        if c=='(' {
            roundb+=1
            st=append(st,1)
        }else if c=='[' {
            squareb+=1
            st=append(st,2)
        }else if c=='{' {
            curlyb+=1
            st=append(st,3)
        }else if c==')' {
            if len(st)==0  || st[len(st)-1]!=1{
                return false
            }
            st=st[:len(st)-1]
            roundb-=1
        }else if c==']' {
            if len(st)==0 || st[len(st)-1]!=2{
                return false
            }
            st=st[:len(st)-1]
            squareb-=1
        }else if c=='}' {
            if len(st)==0 || st[len(st)-1]!=3{
                return false

