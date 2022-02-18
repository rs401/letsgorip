package repository

import (
	"log"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rs401/letsgorip/db"
	"github.com/rs401/letsgorip/forums/models"
	"gorm.io/gorm"
)

var (
	conn           db.Connection
	testForum      *models.Forum
	testForumId    uint64
	testForumTitle string
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
	err = fr.DeleteAllForums()
	if err != nil {
		log.Fatalf("Error deleting all forums before tests: %v", err)
	}

	testForumTitle = "Test forum"
}

func TestGetForums(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name       string
		fields     fields
		wantForums []*models.Forum
		wantErr    bool
	}{
		{
			name:       "Test Get Empty Forums",
			fields:     fields{db: conn.DB()},
			wantForums: []*models.Forum{},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			gotForums, err := fr.GetForums()
			if (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.GetForums() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotForums, tt.wantForums) {
				t.Errorf("forumsRepository.GetForums() = %v, want %v", gotForums, tt.wantForums)
			} else {
				t.Logf("forumsRepository.GetForums() = %v, want %v", gotForums, tt.wantForums)
			}
		})
	}
}

func TestCreateForum(t *testing.T) {
	tf := &models.Forum{
		UserId:      1,
		Title:       "Test forum",
		Description: "Test forum desc.",
	}
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		forum *models.Forum
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Test create forum",
			fields: fields{db: conn.DB()},
			args: args{
				forum: tf,
			},
			wantErr: false,
		},
		{
			name:   "Test create forum not unique",
			fields: fields{db: conn.DB()},
			args: args{
				forum: &models.Forum{
					UserId:      1,
					Title:       "Test forum",
					Description: "Test forum desc.",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			if err := fr.CreateForum(tt.args.forum); (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.CreateForum() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	testForum = tf
	testForumId = tf.Id
}

func TestGetForum(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id uint64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantForum *models.Forum
		wantErr   bool
	}{
		{
			name:      "Test get forum by id",
			fields:    fields{db: conn.DB()},
			args:      args{id: testForumId},
			wantForum: testForum,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			gotForum, err := fr.GetForum(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.GetForum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotForum, tt.wantForum) {
				t.Errorf("forumsRepository.GetForum() = %v, want %v", gotForum, tt.wantForum)
			}
		})
	}
}

func TestGetForumByTitle(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		title string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantForum *models.Forum
		wantErr   bool
	}{
		{
			name:      "Test get forum by id",
			fields:    fields{db: conn.DB()},
			args:      args{title: testForumTitle},
			wantForum: testForum,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			gotForum, err := fr.GetForumByTitle(tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.GetForumByTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotForum, tt.wantForum) {
				t.Errorf("forumsRepository.GetForumByTitle() = %v, want %v", gotForum, tt.wantForum)
			}
		})
	}
}

func TestUpdateForum(t *testing.T) {
	var tf *models.Forum = new(models.Forum)
	tf.Id = testForumId
	tf.UserId = testForum.UserId
	tf.Title = testForum.Title
	tf.Description = "Updated Description"
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		forum *models.Forum
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Test update forum",
			fields:  fields{db: conn.DB()},
			args:    args{forum: tf},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			if err := fr.UpdateForum(tt.args.forum); (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.UpdateForum() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteForum(t *testing.T) {
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
			name:    "Test delete forum",
			fields:  fields{db: conn.DB()},
			args:    args{id: testForumId},
			wantErr: false,
		},
		{
			name:    "Test delete forum not exist",
			fields:  fields{db: conn.DB()},
			args:    args{id: testForumId},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			if err := fr.DeleteForum(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.DeleteForum() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteAllForums(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Test delete all forums",
			fields:  fields{db: conn.DB()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			if err := fr.DeleteAllForums(); (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.DeleteAllForums() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
