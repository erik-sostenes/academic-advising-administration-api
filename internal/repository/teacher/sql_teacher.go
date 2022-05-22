package teacher

var (
	// selectTeachersByCareerAndSubject sql query to find all teachers by career and subject
	selectTeachersByCareerAndSubject = `
		SELECT
			t.tuition,
			t.name,
			t.surnames,
			t.email,
			t.is_active,
			c.cubicle,
			s.subject_id,
			u.university_course_id,
			su.tuition,
			co.tuition
		FROM teachers t
			INNER JOIN cubicles c ON(c.cubicle_id = t.cubicle_id)
				INNER JOIN teachers_has_subjects ts ON(ts.teacher_tuition = t.tuition)
					INNER JOIN subjects s ON(s.subject_id = ts.subject_id)
						INNER JOIN teachers_has_university_courses tuc ON(tuc.teacher_tuition = t.tuition)
							INNER JOIN university_courses u ON(tuc.university_course_id = u.university_course_id)
								INNER JOIN subcoordinators su ON(su.university_course_id = u.university_course_id)
									INNER JOIN coordinators co ON(co.tuition = su.coordinator_tuition)
		WHERE s.subject_id = ? AND u.university_course_id = ? AND t.is_active = true;
		`
) 
