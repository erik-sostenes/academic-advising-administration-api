package schedule

var (
	// insertTeacherSchedules sql query to insert a new teahcer schedule
	insertTeacherSchedules = `
		INSERT INTO teachers_schedules(
			teacher_schedule_id,
			schedule_at,
			from_date,
			to_date,
			teacher_tuition
			)
		VALUES(?, ?, ?, ?, ?);`
	// insertTeacherSchedules sql query to select teacher schedules by Teacher Tuition 
	selectTeacherSchedulesByTeacherTuition = `
  	SELECT
			DAYNAME(ts.schedule_at),
			t.name,
			t.surnames,
			ts.teacher_schedule_id,
			ts.schedule_at,
			ts.from_date,
			ts.to_date,
			ts.student_accountant,
			t.tuition
		FROM teachers t
			INNER JOIN  teachers_schedules ts ON (t.tuition = ts.teacher_tuition)
		WHERE t.tuition = ? AND t.is_active = ?;`
)
