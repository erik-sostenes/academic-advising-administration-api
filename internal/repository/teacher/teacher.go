package teacher

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
type TeacherStorer interface {
	// Find method that seeks teachers with the requirements that are needed
	Find(ctx context.Context, subjectId, universityCourseId string) (model.TeacherCards, error)
}
//go:generate  mockery --case=snake --outpkg=repositorymocks --output=repositorymocks --name=TeacherStorer

// slqTeacherStorer implements TeacherStorer interface
type slqTeacherStorer struct {
	SQL *sql.DB
}

// NewSqlTeacherStorer returns a structure that implements the TeacherStorer interface
func NewSqlTeacherStorer(c repository.Configuration) TeacherStorer {
	switch c.Type {
	case repository.SQL:
		return &slqTeacherStorer{
			SQL: c.SQL,
		}	
	default:
		panic(fmt.Sprintf("%T type is not supported", repository.SQL))
	}
}

func (s *slqTeacherStorer) Find(ctx context.Context, subjectId, universityCourseId string) (teacherCards model.TeacherCards, err error) {
	queryCtx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()

	rows, err := s.SQL.QueryContext(queryCtx, selectTeachersByCareerAndSubject,
		subjectId,
		universityCourseId,
	)
	defer rows.Close()

	if err != nil {
		return 
	}

	for rows.Next() {
		var teacherCard model.TeacherCard

		if err = rows.Scan(
			&teacherCard.Tuition,
			&teacherCard.Name,
			&teacherCard.Surnames,
			&teacherCard.Email,
			&teacherCard.IsActive,
			&teacherCard.Cubicle,
			&teacherCard.SubjectId,
			&teacherCard.UniversityCourseId,
			&teacherCard.SubcoordinatorTuition,
			&teacherCard.CoordinatorTuition,
		); err != nil {
			return 
		}
		teacherCards = append(teacherCards, teacherCard)
	}

	if err = rows.Err(); err != nil {
		return 
	}

	return
}

// CacheTeacherStorer interface that find the teachers in the cache
type CacheTeacherStorer interface {
	TeacherStorer
	// Save encoding data and caches it with its unique key
	Save(ctx context.Context, subjectId, universityCourseId string, teacherCards model.TeacherCards) error
}
//go:generate  mockery --case=snake --outpkg=repositorymocks --output=repositorymocks --name=CacheStorer

// cacheTeacherStorer implements CacheTeacherStorer interface
type cacheTeacherStorer struct {
	RDB *redis.Client
}

// NewCacheTeacherStorer returns a structure that implements the CacheTeacherStorer interface
func NewCacheTeacherStorer(c repository.Configuration) CacheTeacherStorer {
	switch c.Type {
		case repository.NoSQL:
			return &cacheTeacherStorer {
				RDB: c.NoSQL,
			}
		default: 
			panic(fmt.Sprintf("%T type is not supported", repository.NoSQL))
	}
}

// Find data is cached (redis)
func (r *cacheTeacherStorer) Find(ctx context.Context, subjectId, universityCourseId string)(teacherCards model.TeacherCards, err error) {
	key := r.generateKey(subjectId, universityCourseId)

	value, err := r.RDB.Get(ctx, key).Result()
	if err != nil {
		return 
	}

	err = json.Unmarshal([]byte(value), &teacherCards)

	return
}

func (r *cacheTeacherStorer) Save(ctx context.Context, subjectId, universityCourseId string, teacherCards model.TeacherCards) (err error){
	key := r.generateKey(subjectId, universityCourseId)

	if r.RDB.Exists(ctx, key).Val() == 0 {
		return 
	}

	data, err := json.Marshal(teacherCards)
	if err != nil {
		return
	}

	return r.RDB.SetNX(ctx, key, string(data), time.Minute * 10).Err()
}

// generateKey generates the cache key, which will serve as a unique identifier
func (r *cacheTeacherStorer) generateKey(subjectId, universityCourseId string) string {
	// key = teachers-available:by-subjectId-and-universityCourseId
	return "teachers-available" + ":" + "by-" + subjectId + "-and-" + universityCourseId
}
