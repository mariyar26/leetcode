func strStr(haystack string, needle string) int {
    if len(needle)==0 {
        return 0
    }
    j:=0
    for i:=0;i<len(haystack);i+=1 {
        if haystack[i]==needle[j] {
            k:=i
            for ; j<len(needle) && k<len(haystack) && haystack[k]==needle[j];j,k=j+1,k+1 {
            } 
            if j==len(needle) {
                return i
            }else {
                j=0
            }
        }
    }
    return -1
}


func strStr(haystack string, needle string) int {
    for i:=0;i<=len(haystack)-len(needle);i+=1 {
        if haystack[i:i+len(needle)]==needle {
            return i
        }
    }
    return -1
}

//Runtime: 4 ms, faster than 91.15% of Go online submissions for Implement strStr().
//Memory Usage: 2.4 MB, less than 100.00% of Go online submissions for Implement strStr().
//
//
//Implement strStr().
//
//Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
//Clarification:
//
//What should we return when needle is an empty string? This is a great question to ask during an interview.
//
//For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
