package main

import "fmt"

func devide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("This is an error")
	}

	return a / b, nil
}

func main() {
	// ans, _ := devide(10, 0);

	ans,err := devide(10,0);
	if err != nil {
		fmt.Println("This is an error");
		return;
	}

	fmt.Println("this is an answer",ans);
}
