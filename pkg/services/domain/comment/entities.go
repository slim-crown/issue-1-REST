package comment

import "time"

// Comment represents standard comments users can attach
// to a post or another comment.
// replyTo is either and id of another comment or -1 if
// it's only a reply to original post.
type Comment struct {
	ID           int       `json:"id"`
	OriginPost   int       `json:"originPost,omitempty"`
	Commenter    string    `json:"commenter"`
	Content      string    `json:"content"`
	ReplyTo      int       `json:"replyTo,omitempty"`
	CreationTime time.Time `json:"creationTime,omitempty"`
}
