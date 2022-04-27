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
)
