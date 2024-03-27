package main

type Step struct {
	figure       *Figure
	from         int
	to           int
	killedFigure *Figure
	newFigure    *Figure
	isCheckWhite bool
	isCheckBlack bool
}
