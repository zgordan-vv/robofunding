package resolvers

import (
	"github.com/graph-gophers/graphql-go"
)

func strToId(id string) *graphql.ID {
	gqlId := graphql.ID(id)
	return &gqlId
}

func idToStr(gqlID graphql.ID) string {
	return string(gqlID)
}
