package main

import "fmt"

func main() {
	service := NewFacebookService()
	userDetails, apiError, err := service.GetUserDetails()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if apiError != nil {
		fmt.Println(apiError.Error.Message)
		return
	}

	fmt.Println("User ==>" + userDetails.Name)
}
