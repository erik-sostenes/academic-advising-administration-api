package student

var (
	selectStudentRequest = `
			SELECT
				st.tuition,
				st.name,
				st.email,
  			c.cubicle,
				s.subject,
				u.building_number,
				a.advisory_id,
				ts.teacher_schedule_id
			FROM advisories a
				INNER JOIN students st ON(st.tuition = a.student_tuition)
					INNER JOIN subjects s ON(s.subject_id = a.subject_id)
						INNER JOIN teachers t ON(t.tuition = a.teachers_tuition)
							INNER JOIN university_courses u ON(u.university_course_id = a.university_course_id)
								INNER JOIN cubicles c ON(c.cubicle_id = t.cubicle_id)
									INNER JOIN academic_advisory_schedule_record aasr ON(aasr.advisory_id = a.advisory_id)
										INNER JOIN teachers_schedules ts ON(ts.teacher_schedule_id = aasr.teacher_schedule_id)
			WHERE t.is_active = true AND a.is_accepted = false AND t.tuition = ?;`

			selectStudentRequestAccepted = `
				SELECT
					st.tuition,
					st.name,
					st.email,
					c.cubicle,
					s.subject,
					u.university_course
				FROM advisories a
  				INNER JOIN students st ON(st.tuition = a.student_tuition)
  					INNER JOIN subjects s ON(s.subject_id = a.subject_id)
  						INNER JOIN teachers t ON(t.tuition = a.teachers_tuition)
  							INNER JOIN university_courses u ON(u.university_course_id = a.university_course_id)
  								INNER JOIN cubicles c ON(c.cubicle_id = t.cubicle_id)
				WHERE t.is_active = true AND a.is_accepted = true AND t.tuition = ?`

)
