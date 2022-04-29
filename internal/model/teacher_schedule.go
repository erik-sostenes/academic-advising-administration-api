package model

type(
	// MockTeacherSchedule  represents the mock of a teacher schedule 
	MockTeacherSchedule struct {
		day            string
		teacherName    string
		surnameTeacher string
		mockSchedule   MockSchedule
	}
	// MockTeacherSchedules represents a collection of teacher schedules 
	MockTeacherSchedules []MockTeacherSchedule
)
// NewMockTeacherSchedule returns an instance of MockTeacherSchedule if everything is correct
func NewMockTeacherSchedule(day, teacherName, surnameTeacher string, mock MockSchedule) (MockTeacherSchedule, error) {
	return MockTeacherSchedule{
		day: day,
		teacherName: teacherName,
		surnameTeacher: surnameTeacher,
		mockSchedule: mock,
	}, nil
}
// Day represents the day on which the teacher applies an academic advisory.
func (m *MockTeacherSchedule) Day() string {
	return m.day
}
// TeacherName represents the name of the teacher applying an academic advisory.
func (m *MockTeacherSchedule) TeacherName() string {
	return m.teacherName
}
// SurnameTeacher represents the surname of the teacher applying an academic advisory.
func (m *MockTeacherSchedule) SurnameTeacher() string {
	return m.surnameTeacher
}
// MockSchedule represents the mock of a schedule 
func (m *MockTeacherSchedule) MockSchedule() *MockSchedule {
	return &m.mockSchedule
}

