package student

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/itsoeh/academic-advising-administration-api/internal/model"
	"github.com/itsoeh/academic-advising-administration-api/internal/repository"
)

// TeacherStorer interface that find the teachers
type StudentStorer interface {
	// FindRequest method that obtains the requests of all the students who have requested to take an advisory
	FindRequests(ctx context.Context, teacherTuition string) (model.StudentRequests, error)
	// FindAcceptedAccepted method that obtains from all students that their requests have been accepted
	FindAcceptedRequests(ctx context.Context, teacherTuition string) (model.StudentAcceptedRequests, error)
}
//go:generate  mockery --case=snake --outpkg=repositorymocks --output=repositorymocks --name=StudentStorer

type studentStorer struct {
	DB *sql.DB	
}

func NewStudetnStorer(c repository.Configuration) StudentStorer {
	switch c.Type {
	case repository.SQL:
		return &studentStorer{
			DB: c.SQL,
		}
	default:
		panic(fmt.Sprintf("%T type is not supported", repository.SQL))
	} 
}

func (s *studentStorer) FindRequests(ctx context.Context, teacherTuition string) (model.StudentRequests, error) {
	queryCtx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()

	rows, err := s.DB.QueryContext(queryCtx, selectStudentRequest, teacherTuition)
	defer rows.Close()
	
	if err != nil {
		return model.StudentRequests{}, err
	}

	var studentRequests model.StudentRequests

	for rows.Next() {
		var studentRequest model.StudentRequest

		if err := rows.Scan(
			&studentRequest.Tuition,
			&studentRequest.Name,
			&studentRequest.Email,
			&studentRequest.CubicleNumber,
			&studentRequest.Subject,
			&studentRequest.BuildingNumber,
			&studentRequest.AdvisoryId,
			&studentRequest.TeacherScheduleId,
		); err != nil {
			return model.StudentRequests{}, err
		}
		studentRequests = append(studentRequests, studentRequest)
	}

	if err := rows.Err(); err != nil {
		return model.StudentRequests{}, err
	}

	return studentRequests, err 

}

func (s *studentStorer) FindAcceptedRequests(ctx context.Context, teacherTuition string) (model.StudentAcceptedRequests, error) {
	queryCtx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()

	rows, err := s.DB.QueryContext(queryCtx, selectStudentRequestAccepted, teacherTuition)
	defer rows.Close()
	
	if err != nil {
		return model.StudentAcceptedRequests{}, err 
	}

	var studentAcceptedRequests model.StudentAcceptedRequests

	for rows.Next() {
		var studentAcceptedRequest model.StudentAcceptedRequest

		if err := rows.Scan(
			&studentAcceptedRequest.Tuition,
			&studentAcceptedRequest.Name,
			&studentAcceptedRequest.Email,
			&studentAcceptedRequest.CubicleNumber,
			&studentAcceptedRequest.Subject,
			&studentAcceptedRequest.UniversityCourse,
		); err != nil {
			return model.StudentAcceptedRequests{}, err 
		}
		studentAcceptedRequests = append(studentAcceptedRequests, studentAcceptedRequest)
	}

	if err := rows.Err(); err != nil {
		return model.StudentAcceptedRequests{}, err
	}

	return studentAcceptedRequests, err 
}

type CacheStudentStorer interface {
	StudentStorer
	// SaveRequests encoding student requests and caches it with its unique key
	SaveRequests(ctx context.Context, teacherTuition string, isAccepted bool, studentRequests model.StudentRequests) error
	// SaveAcceptedRequests encoding student accepted requests and caches it with its unique key
	SaveAcceptedRequests(ctx context.Context, teacherTuition string, isAccepted bool, studentRequests model.StudentAcceptedRequests) error
	//Delete method used to delete each mock created in the unit tests
	Delete(ctx context.Context, subjectId string, isAccepted bool) error
} 

type cacheStudentStorer struct {
	RDB *redis.Client
}

func NewCacheStudentStorer(c repository.Configuration) CacheStudentStorer {
	switch c.Type {
	case repository.NoSQL:
		return &cacheStudentStorer{
			RDB: c.NoSQL,	
		}
	default:
		panic(fmt.Sprintf("%T type is not supported", repository.NoSQL))
	}
}

func (r *cacheStudentStorer) FindRequests(ctx context.Context, teacherTuition string) (studentRequests model.StudentRequests, err error) {
	key := r.generateKey(teacherTuition, false)

	value, err := r.RDB.Get(ctx, key).Result()
	if err != nil {
		return 
	}

	err = json.Unmarshal([]byte(value), &studentRequests)

	return
}

func (r *cacheStudentStorer) FindAcceptedRequests(ctx context.Context, teacherTuition string) (studentRequestsAccepted model.StudentAcceptedRequests, err error) {
	key := r.generateKey(teacherTuition, true)

	value, err := r.RDB.Get(ctx, key).Result()
	if err != nil {
		return 
	}

	err = json.Unmarshal([]byte(value), &studentRequestsAccepted)

	return
}

func (r *cacheStudentStorer) SaveRequests(ctx context.Context, teacherTuition string, isAccepted bool, studentRequests model.StudentRequests) (err error) {
	key := r.generateKey(teacherTuition, isAccepted)

	data, err := json.Marshal(studentRequests)
	if err != nil {
		return
	}

	return r.RDB.SetNX(ctx, key, string(data), time.Minute * 10).Err()
}

func (r *cacheStudentStorer) SaveAcceptedRequests(ctx context.Context, teacherTuition string, isAccepted bool, studentRequestsAccepted model.StudentAcceptedRequests) (err error) {
	key := r.generateKey(teacherTuition, isAccepted)

	data, err := json.Marshal(studentRequestsAccepted)
	if err != nil {
		return
	}

	return r.RDB.SetNX(ctx, key, string(data), time.Minute * 10).Err()
}

func (r *cacheStudentStorer) Delete(ctx context.Context, subjectId string, isAccepted bool) error {
	return r.RDB.Del(ctx, r.generateKey(subjectId, isAccepted)).Err()
}

// generateKey generates the cache key, which will serve as a unique identifier
func (r *cacheStudentStorer) generateKey(teacherTuition string, isAccepted bool) string {
	return fmt.Sprintf("requests:for-%v:accepted-%v", teacherTuition, isAccepted)
}
