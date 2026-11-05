package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AtlasRegap/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		n, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit %q: %w", cmd.Args[0], err)
		}
		limit = n
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}

	fmt.Printf("Found %d posts for %s:\n", len(posts), user.Name)
	for _, post := range posts {
		date := "unknown date"
		if post.PublishedAt.Valid {
			date = post.PublishedAt.Time.Format("Mon 2 Jan 2006")
		}
		fmt.Printf("%s from %s\n", date, post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		if post.Description.Valid {
			fmt.Printf("    %v\n", post.Description.String)
		}
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil
}
