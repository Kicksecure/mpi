package main

var outputPath = "./"

var architectures = []string{"binary-amd64", "binary-arm64"}

var components = []string{
	"main",
	"contrib",
	"non-free",
}

var mainurl = "https://deb.kicksecure.com/dists/"
var extrasurl = "https://deb.kicksecure.com/dists/"

var suites = []string{"bullseye", "bullseye-stable-proposed-updates", "bullseye-testers", "bullseye-developers"}
