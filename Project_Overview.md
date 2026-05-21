# Educational Institution Management System Backend

## Overview

This project is a backend system for managing educational institutions such as schools and colleges.

The system supports management of:
- Students
- Teachers
- Parents
- Departments
- Classes
- Subjects
- Assignments
- Quizzes
- Attendance
- Results
- Study Materials
- Notification (emails)

The backend follows a Role-Based Access Control (RBAC) architecture where each role has specific permissions.

## Roles

### 1. Institute Admin

Full control over the institution.

Permissions
- Manage departments
- Manage classes & sections
- Manage subjects
- Create teacher accounts
- Create student accounts
- Create parent accounts
- Assign teachers to subjects/classes
- Configure academic settings
- Manage announcements
- View reports & analytics

### 2. Head Of Department

Manages a specific department.

Permissions
- Manage department teachers
- View department students
- Manage department timetable
- Monitor attendance
- View department reports

### 3. Teacher

Permissions
- Create assignments
- Create quizzes
- Upload study materials
- Mark attendance
- Grade submissions
- Publish results
- Send announcements

### 4. Student

Permissions
- View assignments
- Submit assignments
- Attempt quizzes
- View attendance
- View results
- Access study materials

### 5. Parent

Permissions
- View student attendance
- View results
- Receive notifications
- Track student progress

## Core Modules

### Academic Management

- Departments
- Classes
- Sections
- Subjects
- Academic Years
- Semesters (optional)

### User Management

- Students
- Teachers
- Parents

### Assignment Management

Teachers can:
- Create assignments
- Attach files
- Set due dates
- Evaluate submissions

Students can:
- Submit assignments
- View feedback

### Quiz Management

Supports:
- MCQ quizzes
- Subjective questions
- Timer-based quizzes
- Auto evaluation
- Manual grading

### Attendance Management

- Daily attendance
- Subject-wise attendance
- Attendance reports

### Study Material Management

Teachers can upload:
- PDFs
- Videos
- Notes
- External links

### Notification System

Supports:
- Email notifications
- Announcements

## System Flow

### Institute Setup

```text
Institute Admin
    ↓
Create Departments
    ↓
Create Classes & Sections
    ↓
Create Subjects
    ↓
Create Academic Year
```

### User Creation Flow

```text
Institute Admin
    ├── Create Teachers
    ├── Create Students
    └── Create Parents
```

### Teacher Assignment Flow

```text
Institute Admin
    ↓
Assign Teacher
    ├── Department
    ├── Subjects
    └── Classes
```

### Student Enrollment Flow

```text
Institute Admin
    ↓
Assign Department
    ↓
Assign Class & Section
    ↓
Enroll Subjects
    ↓
Link Parent
```

### Assignment Flow

```text
Teacher
    ↓
Create Assignment
    ↓
Assign To Class/Students
    ↓
Students Submit
    ↓
Teacher Evaluates
    ↓
Results Published
```

### Quiz Flow

```text
Teacher
    ↓
Create Quiz
    ↓
Assign To Students
    ↓
Students Attempt Quiz
    ↓
Evaluation
    ↓
Results Published
```

## Recommended Database Tables

```text
users
roles
permissions

departments
classes
sections
subjects

students
teachers
parents

student_parents
teacher_subjects
teacher_classes
student_subjects

academic_years
semesters

assignments
assignment_submissions

quizzes
quiz_questions
quiz_attempts

attendance

study_materials

announcements
notifications

results
grades

timetables
events
```

## Recommended Features

### Audit Logs

Track:
- Login history
- User actions
- Resource changes

### Soft Delete

Use:
- deleted_at TIMESTAMP NULL

Instead of permanent deletion.

## Future Enhancements

- Live classes
- Video conferencing
- Student analytics
- Mobile application
- QR attendance
- Online fee payments
- Discussion forums
- Certificate generation

## Suggested API Structure

```text
/api/v1/auth
/api/v1/users
/api/v1/students
/api/v1/teachers
/api/v1/parents
/api/v1/departments
/api/v1/classes
/api/v1/subjects
/api/v1/assignments
/api/v1/quizzes
/api/v1/attendance
/api/v1/results
/api/v1/notifications
```