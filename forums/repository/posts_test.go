package repository

import (
	"log"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rs401/letsgorip/db"
	"github.com/rs401/letsgorip/forums/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	testPost   *models.Post
	testPostId uint64
)

func init() {
	err := godotenv.Load("../.env.dev")
	if err != nil {
		log.Printf("Error loading .env file (production?): %v\n", err)
	}

	cfg := db.NewConfig()
	conn, err = db.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fr := &forumsRepository{
		db: conn.DB(),
	}
	err = fr.DeleteAllPosts()
	if err != nil {
		log.Fatalf("Error deleting all posts before tests: %v", err)
	}
}

func TestGetPosts(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name      string
		fields    fields
		wantPosts []*models.Post
		wantErr   bool
	}{
		{
			name:      "Test Get Empty Posts",
			fields:    fields{db: conn.DB()},
			wantPosts: []*models.Post{},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			gotPosts, err := fr.GetPosts(testForumId)
			if (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.GetPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPosts, tt.wantPosts) {
				t.Errorf("forumsRepository.GetPosts() = %v, want %v", gotPosts, tt.wantPosts)
			} else {
				t.Logf("forumsRepository.GetPosts() = %v, want %v", gotPosts, tt.wantPosts)
			}
		})
	}
}

func TestCreatePost(t *testing.T) {
	tf := &models.Forum{
		UserId:      1,
		Title:       "Test forum from posts test",
		Description: "Test forum desc.",
	}
	tfr := &forumsRepository{
		db: conn.DB(),
	}
	err := tfr.CreateForum(tf)
	assert.NoError(t, err)
	testForum = tf
	testForumId = tf.Id
	tThread := &models.Thread{
		ForumId: testForumId,
		UserId:  1,
		Title:   "Test thread from posts test",
		Msg:     "Test thread msg",
	}
	err = tfr.CreateThread(tThread)
	assert.NoError(t, err)
	testThread = tThread
	testThreadId = tThread.Id
	tPost := &models.Post{
		ThreadId: testThreadId,
		UserId:   1,
		Msg:      "Test post msg",
	}
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		post *models.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Test create post",
			fields: fields{db: conn.DB()},
			args: args{
				post: tPost,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			if err := fr.CreatePost(tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.CreatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	testPost = tPost
	testPostId = tPost.Id
}

func TestGetPost(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id uint64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantPost *models.Post
		wantErr  bool
	}{
		{
			name:     "Test get post by id",
			fields:   fields{db: conn.DB()},
			args:     args{id: testPostId},
			wantPost: testPost,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			gotPost, err := fr.GetPost(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.GetPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPost, tt.wantPost) {
				t.Errorf("forumsRepository.GetPost() = %v, want %v", gotPost, tt.wantPost)
			}
		})
	}
}

func TestUpdatePost(t *testing.T) {
	var tPost *models.Post = new(models.Post)
	tPost.Id = testPostId
	tPost.ThreadId = testThreadId
	tPost.UserId = testPost.UserId
	tPost.Msg = "Updated Msg"
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		post *models.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Test update post",
			fields:  fields{db: conn.DB()},
			args:    args{post: tPost},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			if err := fr.UpdatePost(tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeletePost(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Test delete post",
			fields:  fields{db: conn.DB()},
			args:    args{id: testPostId},
			wantErr: false,
		},
		{
			name:    "Test delete post not exist",
			fields:  fields{db: conn.DB()},
			args:    args{id: testPostId},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			if err := fr.DeletePost(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteAllPosts(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Test delete all posts",
			fields:  fields{db: conn.DB()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			if err := fr.DeleteAllPosts(); (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.DeleteAllPosts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
