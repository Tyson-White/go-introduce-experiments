package operations

var resultMap = map[string]operation{
	"result": operation{},
	"last":   operation{},
}

func SumOfAB(a float64, b float64) map[string]operation {

	result := a + b

	op, last := appendOperation("sum", a, b, result)

	resultMap["result"] = op
	resultMap["last"] = last

	return resultMap

}

func DifferenceOfAB(a float64, b float64) map[string]operation {

	result := a - b

	op, last := appendOperation("diff", a, b, result)

	resultMap["result"] = op
	resultMap["last"] = last

	return resultMap

}

func MultipleAB(a float64, b float64) map[string]operation {

	result := a * b

	op, last := appendOperation("mult", a, b, result)

	resultMap["result"] = op
	resultMap["last"] = last

	return resultMap

}

func DivideAB(a float64, b float64) map[string]operation {

	result := a / b

	op, last := appendOperation("div", a, b, result)

	resultMap["result"] = op
	resultMap["last"] = last

	return resultMap

}
