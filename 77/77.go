package combine

import "container/list"

func combine(n int, k int) [][]int {
	ans := &[][]int{}
	serch3(n, k, list.New(), &ans)
	return *ans
}
func serch3(n int, k int, currut *list.List, ans **[][]int) {
    if n < k {
        return 
    }
	if k == 1 {
        for ;n>0;n--{
		l := currut.Len()
		an := make([]int, l+1)
		e := currut.Front()
            i :=0
		for {
			if e != nil {
				
                an[i] =e.Value.(int)
                i++
			} else {
				break
			}
			e = e.Next()
		}
		 an[i] =n
		**ans = append(**ans, an)
        }
	} else {
		
		k--
        
        for ;n>0; {
			currut.PushBack(n)
			n--
			
            
			serch3(n, k, currut, ans)
             currut.Remove(currut.Back())
           
		}
       
		
	}
     
	return
}

