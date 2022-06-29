package teacher

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/itsoeh/academic-advising-administration-api/internal/model"
)

// TeachersStorer interface containing the methods to interact with the storage
type TeacherStorer interface {
	// StorageFindTechers method that seeks teachers with the requirements that are needed
	StorageFindTechers(ctx context.Context, subjectId, universityCourseId string) (model.TeacherCards, error)
}
//go:generate  mockery --case=snake --outpkg=repositorymocks --output=internal/repository/repositorymocks

// slqTeacherStorer implements TeacherStorer interface
type slqTeacherStorer struct {
	DB *sql.DB
}

// NewSQLTeacherStorer returns a structure that implements the TeacherStorer interface
func NewSQLTeacherStorer(DB *sql.DB) TeacherStorer {
	return &slqTeacherStorer{
		DB: DB,
	}
}

func (s *slqTeacherStorer) StorageFindTechers(ctx context.Context, subjectId, universityCourseId string) (model.TeacherCards, error) {
	queryCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	rows, err := s.DB.QueryContext(queryCtx, selectTeachersByCareerAndSubject,
		subjectId,
		universityCourseId,
	)
	defer rows.Close()

	if err != nil {
		return model.TeacherCards{}, model.InternalServerError("An error has ocurred while obtainig the teachers.")
	}

	var teacherCards model.TeacherCards

	for rows.Next() {
		var teacherCard model.TeacherCard

		if err := rows.Scan(
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
			return teacherCards, model.InternalServerError("Error when searching for the teacher.")
		}
		teacherCards = append(teacherCards, teacherCard)
	}

	if err := rows.Err(); err != nil {
		return teacherCards, model.InternalServerError(err.Error())
	}

	return teacherCards, err
}

// redisTeacherStorer implements TeacherStorer interface
type redisTeacherStorer struct {
	RDB *redis.Client
}

// NewRedisTeacherStorer returns a structure that implements the TeacherStorer interface
func NewRedisTeacherStorer(RDB *redis.Client) TeacherStorer {
	return &redisTeacherStorer{
		RDB: RDB,
	}
}

// StorageFindTechers data is cached (redis)
func (r *redisTeacherStorer) StorageFindTechers(ctx context.Context, subjectId, universityCourseId string)(model.TeacherCards, error) {
	var teacherCards model.TeacherCards

	key := r.GenerateCacheKey(subjectId, universityCourseId)

	value, err := r.RDB.Get(ctx, key).Result()
	if err != nil {
		return teacherCards, err
	}

	err = json.Unmarshal([]byte(value), &teacherCards)

	return teacherCards, err
}

// StorageSaveCache encoding data and caches it with its unique key
func (r *redisTeacherStorer) StorageSaveCache(ctx context.Context, key string, teacherCards model.TeacherCards) error{
	b, err := json.Marshal(teacherCards)
	if err != nil {
		return err
	}

	return r.RDB.SetNX(ctx, key, string(b), time.Minute*10).Err()
}

// ExistsKey checks if key exists in redis
func (r *redisTeacherStorer) ExistsKey(ctx context.Context, key string) bool {
	return r.RDB.Exists(ctx, key).Val() != 0
}

// GenerateCacheKey generates the cache key, which will serve as a unique identifier
func (r *redisTeacherStorer) GenerateCacheKey(subjectId, universityCourseId string) string {
	// key = teachers-available:by-subjectId-and-universityCourseId
	return "teachers-available" + ":" + "by-" + subjectId + "-and-" + universityCourseId
}
