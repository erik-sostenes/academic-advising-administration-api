package schedule

var (
	insertTeacherSchedules = `
		INSERT INTO teachers_schedules(
			teacher_schedule_id,
			schedule_at,
			from_date,
			to_date,
			teacher_tuition
			)
		VALUES(?, ?, ?, ?, ?);`
)
