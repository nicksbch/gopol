package main

type InteractionType int

const (
	InteractionTypeMeeting InteractionType = iota
	InteractionTypeMessage
	InteractionTypeCall
	InteractionTypeInPerson
)

var InteractionTypeName = map[InteractionType]string{
	InteractionTypeMeeting:  "meeting",
	InteractionTypeMessage:  "message",
	InteractionTypeCall:     "call",
	InteractionTypeInPerson: "in person",
}

func (i InteractionType) String() string {
	return InteractionTypeName[i]
}

type InteractionSentiment int

const (
	InteractionSentimentPositive InteractionSentiment = iota
	InteractionSentimentNegative
	InteractionSentimentNeutral
)

var InteractionSentimentName = map[InteractionSentiment]string{
	InteractionSentimentPositive: "positive",
	InteractionSentimentNegative: "negative",
	InteractionSentimentNeutral:  "neutral",
}

func (i InteractionSentiment) String() string {
	return InteractionSentimentName[i]
}

type InteractionImpact int

const (
	InteractionImpactPositive InteractionImpact = iota
	InteractionImpactNegative
	InteractionImpactNeutral
)

var InteractionImpactName = map[InteractionImpact]string{
	InteractionImpactPositive: "positive",
	InteractionImpactNegative: "negative",
	InteractionImpactNeutral:  "neutral",
}

func (i InteractionImpact) String() string {
	return InteractionImpactName[i]
}
