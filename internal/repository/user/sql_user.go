package user

var (
// selectStudentByTuition sql query to select the password of a student by tuition
	selectStudentByTuition = `
		SELECT
			s.password,
			s.tuition
		FROM students s 
		WHERE s.tuition = ?;`

// selectTeacherByTuition sql query to select the password of a teacher by tuition
	selectTeacherByTuition = `
		SELECT
			t.password,
			t.tuition
		FROM teachers t 
		WHERE t.tuition = ?;`
)
