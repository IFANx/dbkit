package randomly

func ProbFiveOne() bool {
	return RandIntGap(0, 5) == 0
}

func ProbTenOne() bool {
	return RandIntGap(0, 10) == 0
}

func ProbTwentyOne() bool {
	return RandIntGap(0, 10) == 0
}

func ProbFiftyOne() bool {
	return RandIntGap(0, 50) == 0
}
