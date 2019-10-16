# Database Tables Documentation
This document contains a brief documentation of each table, the objective is to try to explain the content of each column. If the description of a column contains the _trivial_ text, it is because the name of the column is enough to interpret it.

Feel free to add more documentation if necessary, for example the summary of what each table is or enhance the description of each column.

## 1. comment
Column              | Type     | Description
:-------------------| :--------|:---------------
id                  | int      | auto-increment count
content             | string   | trivial
likes               | int      | number of users who like the comment
dislikes            | int      | number of users who dislike the comment
created_date        | date     | when the comment was made
last_modified_date  | date     | when the comment was modified by last time
story_id            | int      | story where the comment was added
created_by          | int      | who created the comment

## 2. contribution
Column              | Type     | Description
:-------------------| :------- | :---------------
id                  | int      | auto-increment count
score               | int      | points earned for the contribution
description         | string   | why was the contribution?
created_date        | date     | when the contribution was made
last_modified_date  | date     | when the contribution was modified by last time
person_id           | int      | user who did the contribution

## 3. enrollment
Column              | Type     | Description
:-------------------| :--------|:---------------
id                  | int      | auto-increment count
status              | int      | approved, disapproved or failed
grade               | float    | final grade obtained
is_verified         | int      | trivial
doc_verifier        | []byte   | doc received to verifier the enrollment (photo)
created_date        | date     | when the enrollment was made
last_modified_date  | date     | when the enrollment was modified by last time
person_id           | int      | user who did the enrollment
course_id           | int      | which course the enrollment was made

## 4. career
Column              | Type      | Description
:-------------------| :-------- |:---------------
id                  | int       | auto-increment count
acronym             | string    | trivial
name                | string    | trivial
description         | string    | trivial
logo                | string    | pic of the career
is_verified         | int      | trivial
doc_verifier        | []byte    | doc received to verifier it (photo)
reference           | string    | link of any reference
created_date        | date      | when the career was created
last_modified_date  | date      | when the career was modified by last time
faculty_id          | int       | faculty to which it belongs
album_id            | int       | trivial
created_by          | int       | user who created the career

## 5. award
Column              | Type     | Description
:-------------------| :--------|:---------------
id                  | int      | auto-increment count
year                | int      | year in which the award was obtained
period              | int      | semester in which the award was obtained
description         | string   | trivial
event               | string   | event in which the award was obtained
place               | int      | place obtained
date                | date     | when the award item was created
is_verified         | int      | trivial
doc_verifier        | []byte   | doc received to verifier its existence (photo)
created_date        | date     | when the award was created
last_modified_date  | date     | when the award was modified by last time
career_id           | int      | career to which it belongs
album_id            | int      | trivial
created_by          | int      | user who created the award

## 6. interest
Column              | Type     | Description
:-------------------| :--------|:---------------
id                  | int      | auto-increment count
tag                 | string   | brief title that represent the interest
is_skill            | bool     | is it a skill or just an interest?
created_date        | date     | when the interest was created
last_modified_date  | date     | when the interest was modified by last time
created_by          | int      | user who created the interest

## 7. photo
Column              | Type     | Description
:-------------------| :--------|:---------------
id                  | int      | auto-increment count
pic                 | []byte   | trivial
descripion          | string   | brief description
created_date        | date     | when the photo was created
last_modified_date  | date     | when the photo was modified by last time
album_id            | int      | album to which the photo belongs
created_by          | int      | user who created the photo

## 8. schedule
Column              | Type      | Description
:-------------------| :---------|:---------------
id                  | int       | auto-increment count
start               | string    | start time
end                 | int       | end time
day                 | int       | day (monday, tuesday, etc.)
room                | varchar   | where the course is teach
type_lesson         | int       | types lesson (lab, theoretical, etc)
is_verified         | int       | trivial
doc_verifier        | bool      | doc received to verifie it (photo)
created_date        | date      | when the schedule was created
last_modified_date  | date      | when the schedule was modified by last time
course_id           | string    | course to which the schedule belongs
created_by          | int       | user who created the photo

## 9. story
Column              | Type      | Description
:-------------------| :---------|:---------------
id                  | int       | auto-increment count
title               | string    | trivial
body                | string    | trivial
recommended         | int       | number of users who recommend the story
unrecommended       | int       | number of user who unrecommend the story
views               | int       | number of users who saw the story
is_verified         | int       | trivial
doc_verifier        | string    | doc received to verifier the story (photo)
allow_comments      | bool      | trivial
layer               | int       | in which layer the story was created
created_date        | date      | when the story was created
last_modified_date  | date      | when the story was modified by last time
created_by          | int       | who created the story
career_id           | int       | the story belong to this career

## 10. tag_container
Column              | Type     | Description
:-------------------| :--------|:---------------
id                  | int      | auto-increment count
tag                 | string   | brief title
agree               | string   | number of users who agree the tag
disagree            | int      | number of users who disagree the tag
reference           | int      | link or any reference
type                | int      | tree possible types: speciality, requirement and jobs
layer               | int      | in which layer the story was created
created_date        | date     | when the tag was created
last_modified_date  | date     | when the tag was modified by last time
is_verified         | int      | trivial
doc_verifier        | string   | doc received to verifier the tag (photo)
career_id           | int      | career to which it belong
created_by          | int      | who created the tag

## 11. courses_connection
Column              | Type    | Description
:-------------------| :-------|:---------------
id                  | int     | auto-increment count
year                | int     | year to which the connection belongs
period              | int     | semester
threshold_credits   | int     | minimum number of credits to open the parent course
child_course        | string  | code of the course child
parent_course       | string  | code of the course parent
created_date        | date    | when the connection was created
last_modified_date  | date    | when the connection was modified by last time
career_id           | int     | career which has that course connection
created_by          | int     | who created the connection

## 12. course
Column              | Type    | Description
:-------------------| :-------|:---------------
id                  | int     | auto-increment count
code                | string  | code of the course
name                | string  | trivial
description         | string  | trivial
credits             | int     | number of credits of the course
batch               | int     | a course can have many groups
year                | int     | year in which the course opened
period              | int     | semeste in which the course opened
semester            | int     | semester to which the course belongs
total_hours         | int     | trivial
is_verified         | int     | trivial
doc_verifier        | []byte  | doc received to verifier it (photo)
created_date        | date    | when the course was created
last_modified_date  | date    | when the course was modified by last time
type_course_id      | int     | can be many types: mandatory, optional, etc.
career_id           | int     | career to which the course belongs
professor_id        | int     | professor in charge of the course
created_by          | int     | who created the course

## 13. type_course
Column              | Type    | Description
:-------------------| :-------|:---------------
id                  | int     | auto-increment count
acronym             | string  | course type code
name                | string  | trivial
description         | string  | trivial
is_verified         | int     | trivial
doc_verifier        | []byte  | doc received to verifier it (photo)
created_date        | date    | when the type course was created
last_modified_date  | date    | when the type course was modified by last time
required_credits    | int     | total number of credits required for this type of course.
career_id           | int     | career to which the type courses belongs
created_by          | int     | who created the type course

## 14. faculty
Column              | Type    | Description
:-------------------| :-------|:---------------
id                  | int     | auto-increment count
acronym             | string  | faculty code
name                | string  | trivial
description         | string  | trivial
logo                | []byte  | trivial
is_verified         | int     | trivial
doc_verifier        | []byte  | doc received to verifier it (photo)
created_date        | date    | when the faculty was created
last_modified_date  | date    | when the faculty was modified by last time
university_id       | int     | university to which the faculty belongs
created_by          | int     | who created the faculty


## 15. career_person
Column              | Type    | Description
:-------------------| :-------|:---------------
person_id           | int     | trivial
career_id           | int     | trivial
is_verified         | int     | trivial
doc_verifier        | []byte  | doc received to verifier it (photo)
created_date        | date    | when the record was created
last_modified_date  | date    | when the record was modified by last time

## 16. person
Column              | Type     | Description
:-------------------| :--------|:---------------
id                  | int      | auto-increment count
code                | string   | university id
gender              | int      | trivial
first_name          | string   | trivial
last_name           | string   | trivial
phone               | string   | trivial
email               | string   | trivial
avatar              | []byte   | trivial
type                | int      | trivial
home_city           | string   | trivial
current_city        | string   | trivial
ethnic              | int      | trivial
nacionality         | string   | trivial
birth_date          | date     | trivial
admission_year      | int      | year in which the user started university
period              | int      | semester
is_verified         | int      | trivial
doc_verifier        | []byte   | doc received to verifier it (photo)
created_date        | date     | when the record was created
last_modified_date  | date     | when the record was modified by last time

## 17. type_university
Column              | Type    | Description
:-------------------| :-------|:---------------
id                  | int     | auto-increment count
name                | string  | trivial
description         | string  | trivial
is_verified         | int     | trivial
doc_verifier        | []byte  | doc received to verifier it (photo)
created_date        | date    | when the type university was created
last_modified_date  | date    | when the type university was modified by last time
created_by          | int     | who created the type university


## 18. university
Column              | Type    | Description
:-------------------| :-------|:---------------
id                  | int     | auto-increment count
local_rank          | int     | trivial
global_rank         | int     | trivial
acronym             | string  | trivial
name                | string  | trivial
region              | string  | trivial
description         | string  | trivial
latitude            | float   | trivial
longitude           | float   | trivial
logo                | []byte  | pic of the university
is_verified         | int    | trivial
doc_verifier        | []byte  | doc received to verifier it (photo)
reference           | string  | link for the main web page of the university
type_university_id  | int     | can be publi, private, etc.
album_id            | int     | trivial
created_date        | date    | when the university was created
last_modified_date  | date    | when the university was modified by last time
created_by          | int     | who created the university

## 19. album
Column              | Type     | Description
:-------------------| :--------|:---------------
id                  | int      | auto-increment count
name                | string   | trivial
description         | string   | trivial
created_date        | date     | when the album was created
last_modified_date  | date     | when the album was modified by last time
created_by          | int      | who created the album