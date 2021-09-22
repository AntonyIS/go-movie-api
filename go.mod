module github.com/AntonyIS/GoProject

go 1.16

replace github.com/AntonyIS/GoProject/models => ./models

replace github.com/AntonyIS/GoProject/controllers => ./controllers

require (
	github.com/AntonyIS/GoProject/controller v0.0.0-20210922130006-fb9021a4755d
	github.com/AntonyIS/GoProject/models v0.0.0-00010101000000-000000000000 // indirect
	github.com/gin-gonic/gin v1.7.4
)
