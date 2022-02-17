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
	testThread      *models.Thread
	testThreadId    uint64
	testThreadTitle string
)

func init() {
	err := godotenv.Load("../.env.dev")
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	cfg := db.NewConfig()
	conn, err = db.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fr := &forumsRepository{
		db: conn.DB(),
	}
	err = fr.DeleteAllThreads()
	if err != nil {
		log.Fatalf("Error deleting all threads before tests: %v", err)
	}

	testThreadTitle = "Test thread"
}

func TestGetThreads(t *testing.T) {
	fr := &forumsRepository{db: conn.DB()}
	err := fr.DeleteAllThreads()
	assert.NoError(t, err)
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name        string
		fields      fields
		wantThreads []*models.Thread
		wantErr     bool
	}{
		{
			name:        "Test Get Empty Threads",
			fields:      fields{db: conn.DB()},
			wantThreads: []*models.Thread{},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			gotThreads, err := fr.GetThreads(testForumId)
			if (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.GetThreads() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotThreads, tt.wantThreads) {
				t.Errorf("forumsRepository.GetThreads() = %v, want %v", gotThreads, tt.wantThreads)
			} else {
				t.Logf("forumsRepository.GetThreads() = %v, want %v", gotThreads, tt.wantThreads)
			}
		})
	}
}

func TestCreateThread(t *testing.T) {
	tf := &models.Forum{
		UserId:      1,
		Title:       "Test forum from threads test",
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
		Title:   "Test thread",
		Msg:     "Test thread msg",
	}
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		thread *models.Thread
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Test create thread",
			fields: fields{db: conn.DB()},
			args: args{
				thread: tThread,
			},
			wantErr: false,
		},
		{
			name:   "Test create thread not unique",
			fields: fields{db: conn.DB()},
			args: args{
				thread: &models.Thread{
					UserId: 1,
					Title:  "Test thread",
					Msg:    "Test thread desc.",
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
			if err := fr.CreateThread(tt.args.thread); (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.CreateThread() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	testThread = tThread
	testThreadId = tThread.Id
}

func TestGetThread(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id uint64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantThread *models.Thread
		wantErr    bool
	}{
		{
			name:       "Test get thread by id",
			fields:     fields{db: conn.DB()},
			args:       args{id: testThreadId},
			wantThread: testThread,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			gotThread, err := fr.GetThread(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.GetThread() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotThread, tt.wantThread) {
				t.Errorf("forumsRepository.GetThread() = %v, want %v", gotThread, tt.wantThread)
			}
		})
	}
}

func TestUpdateThread(t *testing.T) {
	var tThread *models.Thread = new(models.Thread)
	tThread.Id = testThreadId
	tThread.ForumId = testForumId
	tThread.UserId = testThread.UserId
	tThread.Title = testThread.Title
	tThread.Msg = "Updated Msg"
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		thread *models.Thread
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Test update thread",
			fields:  fields{db: conn.DB()},
			args:    args{thread: tThread},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			if err := fr.UpdateThread(tt.args.thread); (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.UpdateThread() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteThread(t *testing.T) {
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
			name:    "Test delete thread",
			fields:  fields{db: conn.DB()},
			args:    args{id: testThreadId},
			wantErr: false,
		},
		{
			name:    "Test delete thread not exist",
			fields:  fields{db: conn.DB()},
			args:    args{id: testThreadId},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			if err := fr.DeleteThread(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.DeleteThread() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteAllThreads(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Test delete all threads",
			fields:  fields{db: conn.DB()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := &forumsRepository{
				db: tt.fields.db,
			}
			if err := fr.DeleteAllThreads(); (err != nil) != tt.wantErr {
				t.Errorf("forumsRepository.DeleteAllThreads() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
