//The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//
//And then read line by line: "PAHNAPLSIIGYIR"
//
//Write the code that will take a string and make this conversion given a number of rows:
//
//string convert(string s, int numRows);
//
// 
//
//Example 1:
//
//Input: s = "PAYPALISHIRING", numRows = 3
//Output: "PAHNAPLSIIGYIR"
//
//Example 2:
//
//Input: s = "PAYPALISHIRING", numRows = 4
//Output: "PINALSIGYAHRPI"
//Explanation:
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
//
//Example 3:
//
//Input: s = "A", numRows = 1
//Output: "A"
//
// 
//
//Constraints:
//
//    1 <= s.length <= 1000
//    s consists of English letters (lower-case and upper-case), ',' and '.'.
//    1 <= numRows <= 1000
//
//
func convert(s string, numRows int) string {
    if len(s)<3 || numRows==1 || len(s)<numRows{
        return s
    }

    r:=[]rune(s)
    row:=0
    prev_index:=0
    first:=(2*(numRows-1))
    last:=0
    r[0]=rune(s[0])
    for i:=0;i<len(s);{
        if prev_index==-1{
            index:=row
            i+=1
            r[i]=rune(s[index])
            prev_index=index
            continue
        }
        if first !=0 {
            index:=prev_index + first
            if index>=len(s){
                prev_index=-1
                row+=1
                first-=2
                last+=2
                continue
            }
            i+=1
            r[i]=rune(s[index])
            prev_index=index
        }
        if i+1==len(s){
            break
        }
        if last !=0 {
            index:=prev_index +last
            if index>=len(s){
                prev_index=-1
                row+=1
                first-=2
                last+=2
                continue
            }
            i+=1
            r[i]=rune(s[index])
            prev_index=index
        }

    }
    return string(r)
}


//Runtime: 6 ms, faster than 84.38% of Go online submissions for Zigzag Conversion.
//Memory Usage: 4.3 MB, less than 83.93% of Go online submissions for Zigzag Conversion.



import "unicode"
func convert(s string, numRows int) string {
    if numRows==1 {
        return s
    }
  var matrix [1000][1000]rune
    st:=[]rune(s)
    k:=0
    var output []rune
    for i:=0; i<len(st); {
        for j:=0; j<numRows-1 && i<len(st); j++ {
            matrix[k][j]=st[i]
            i++
        }
        for j:=numRows-1; j>0 && i<len(st) ; j-- {
            k++
            matrix[k][j]=st[i]
            i++
        }
        k++
    }

        for j:=0;j<numRows;j++ {
                for i:=0;i<k;i++ {
                    if unicode.IsLetter(matrix[i][j]) || matrix[i][j] ==',' || matrix[i][j] =='.'  {
             output=append(output,matrix[i][j])
            }
        }
    }
    return string(output)
}

//Runtime: 207 ms
//Memory Usage: 9.8 MB
