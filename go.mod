module github.com/ohad83/mock

go 1.20

require (
	go.uber.org/mock v0.2.0
	golang.org/x/mod v0.11.0
	golang.org/x/tools v0.2.0
)

require golang.org/x/sys v0.1.0 // indirect

replace go.uber.org/mock => github.com/ohad83/mock v0.0.0-20230717153152-7beb269acad6
