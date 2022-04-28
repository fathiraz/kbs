package main

import (
	"fmt"
	"fpb/config"
	"fpb/helper"
	"fpb/model"
)

// main program
func main() {

	// set data
	config.ParseEnv()
	environment := config.GetEnv()

	// set default parser
	parseFlag := helper.Parser{
		Owner:  environment.DefaultOwner,
		Apples: environment.DefaultApples,
		Cakes:  environment.DefaultCakes,
	}

	// parse command line
	parseFlag = helper.ParseFlag(&parseFlag)

	// set box
	box := model.Box{
		Cakes:  parseFlag.Cakes,
		Apples: parseFlag.Apples,
	}

	// show result
	fmt.Println(fmt.Sprintf("%s have %v cakes and %v apples.", parseFlag.Owner, box.GetCakes(), box.GetApples()))
	fmt.Println(fmt.Sprintf("%s can make %v box.", parseFlag.Owner, box.GetBoxEvenly()))
	fmt.Println(fmt.Sprintf("There are %v cakes and %v apples in each box.", box.GetBoxCakes(), box.GetBoxApples()))
}
