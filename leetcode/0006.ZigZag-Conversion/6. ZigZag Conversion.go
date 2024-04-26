package leetcode

import "fmt"

func convert(s string, numRows int) string {
	matrix, down, up := make([][]byte, numRows, numRows), 0, numRows-2
	for i := 0; i != len(s); {
		if down != numRows {
			matrix[down] = append(matrix[down], byte(s[i]))
			down++
			i++
		} else if up > 0 {
			matrix[up] = append(matrix[up], byte(s[i]))
			up--
			i++
		} else {
			up = numRows - 2
			down = 0
		}
	}
	solution := make([]byte, 0, len(s))
	for _, row := range matrix {
		for _, item := range row {
			solution = append(solution, item)
		}
	}
	return string(solution)
}
func convert0315_1(input string, n int) string{
	//tmpSlice := [][]byte{}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	tmpSlice := make([][]byte, n)
	//m := len(input)
	i,j := 1,0
	
	firstLap := true
	m:=0;
	for m<len(input){
		// firstLap means starting the zag-zag in the first vertical down
		if firstLap {
			i = 0
		}else {
			// switch i to 1 if not the firstLap
			i = 1
			// j shall minus 1 for it is vertical down
			j--
		}
		firstLap = false
		for i < n {
			//tmpSlice[i][j] = input[m]
			tmpSlice[i] = append(tmpSlice[i], []byte{input[m]}...)
			i++
			// m always ++
			m++
		}
		i=1
		for i < n {
			//tmpSlice[][j+1] = input[m]
			fmt.Println("i is ", i, "n is " , n)
			fmt.Println(tmpSlice[n-1-i])
			tmpSlice[n-1-i] = append(tmpSlice[n-1-i], []byte{input[m]}...)
			// m always ++
			m++
			i++
			fmt.Println("=======i is ", i, "n is " , n)
			// j ++ for it is zag-zag
			j++
		}
	}
	tmpByteArr := []byte{}
	for _ , item := range tmpSlice {
		tmpByteArr = append(tmpByteArr, item...)
	}
	res := string(tmpByteArr)
	fmt.Println(res)
	return res
}

func convert0315(input string, n int) string{
	//tmpSlice := [][]byte{}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	tmpSlice := make([][]byte, n)
	//m := len(input)
	i,m,len := 1,0,len(input)
	
	firstLap := true
	
	for m<len{
		// firstLap means starting the zag-zag in the first vertical down
		if firstLap {
			i = 0
		}else {
			// switch i to 1 if not the firstLap
			i = 1
			
		}
		firstLap = false
		for i < n && m<len{
			//tmpSlice[i][j] = input[m]
			tmpSlice[i] = append(tmpSlice[i], []byte{input[m]}...)
			i++
			// m always ++
			m++
		}
		i=1
		for i < n && m<len{
			//tmpSlice[][j+1] = input[m]
			fmt.Println("i is ", i, "n is " , n)
			fmt.Println(tmpSlice[n-1-i])
			tmpSlice[n-1-i] = append(tmpSlice[n-1-i], []byte{input[m]}...)
			// m always ++
			m++
			i++
			fmt.Println("=======i is ", i, "n is " , n)
			
		}
	}
	tmpByteArr := []byte{}
	for _ , item := range tmpSlice {
		tmpByteArr = append(tmpByteArr, item...)
	}
	res := string(tmpByteArr)
	fmt.Println(res)
	return res
}