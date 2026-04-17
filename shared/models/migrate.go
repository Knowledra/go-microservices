package models

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		// &Institute{},

		&Admin{}, &Teacher{}, &Student{}, &Parent{},

		&StudentParent{},

		&Class{}, &Section{}, &Department{}, &Subject{},

		//&ClassSubject{}, //---> This is not required as students will be selecting their subjects and teachers will be assigned to subjects

		// This table will be used to assign teachers to subjects and classes
		// If a teacher is teaching a subject to multiple classes, there will be multiple entries for that teacher and subject with different class IDs
		// We can filter by class ID to get the specific teacher for that class and subject
		// &TeacherSubject{},

	// 	&Attendance{},

	// 	&Exam{}, &ExamSubject{}, &Result{},

	// 	&Assignment{}, &AssignmentSubmission{},

	// 	&Quiz{},
	)
}
