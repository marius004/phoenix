package models

type ExecuteRequest struct {
	SourceCode []byte
	Language   	*Language
	Problem     *Problem
}
