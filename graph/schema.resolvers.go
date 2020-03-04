// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"log"
	"time"

	"github.com/yogihardi/graphqlstream/graph/generated"
)

func (r *subscriptionResolver) Ticker(ctx context.Context) (<-chan string, error) {
	log.Println("entering ticker resolver")
	stream := make(chan string)
	ticker := time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("context is done")
				return

			case t := <-ticker.C:
				stream <- t.String()
			}
		}
	}()

	return stream, nil
}

func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
