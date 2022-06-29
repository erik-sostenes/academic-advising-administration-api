package teacher

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/itsoeh/academic-advising-administration-api/internal/model"
	"github.com/itsoeh/academic-advising-administration-api/internal/repository"
)

func Test_TeacherStorer_StorageFindTechers(t *testing.T) {
	t.Run("find successful teachers 'MySQL'", func(t *testing.T) {
		teacherStorer := NewSQLTeacherStorer(repository.NewMySQL())
			
		got, err := teacherStorer.StorageFindTechers(context.Background(), "B43B4F", "JDS7G7")
		if err != nil {
			t.Fatal(err)
		}

		expected := model.TeacherCards{}
		if fmt.Sprintf("%T", got) != fmt.Sprintf("%T", expected) {
			t.Fatalf("expected %T, got %T", expected, got)
		}
	})

	t.Run("find successful teachers 'redis'", func(t *testing.T) {
		teacherStorer := NewRedisTeacherStorer(repository.NewRedis())

		got, err := teacherStorer.StorageFindTechers(context.TODO(), "HJHDS", "KNDJJ")
		if err != redis.Nil {
			t.Fatal(err)
		}
		
		expected := model.TeacherCards{}
		if fmt.Sprintf("%T", got) != fmt.Sprintf("%T", expected) {
			t.Fatalf("expected %T, got %T", expected, got)
		}
	})
} 
