package model

type Session struct {
	BcryptTocken      string `bson:"bcryptTocken"`
	TimeCreatedTocken string `bson:"timeCreatedTocken"`
	Guid              string `bson:"guid"`
}