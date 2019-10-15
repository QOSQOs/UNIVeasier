# Database Tables Documentation
This document contains a brief documentation of each table, the objective is to try to explain the content of each column. If the description of a column contains the _trivial_ text, it is because the name of the column is enough to interpret it.

Feel free to add more documentation if necessary, for example the summary of what each table is or enhance the description of each column.

## 1. comment
Column       | Type     | Description
:------------| :--------|:---------------
id           | int      | auto-increment count
content      | string   | trivial
like         | int      | number of users who like the comment
dislike      | int      | number of users who dislike the comment
story_id     | int      | story where the comment was added
date         | date      | when the comment was made
person_id    | int      | who created the comment

## 2. contribution
Column      | Type      | Description
:-----------| :-------- | :---------------
id          | int       | auto-increment count
date        | date      | when the contribution was made
score       | int       | points earned for the contribution
description | string    | why was the contribution?
person_id   | int       | user who did the contribution

## 3. enrollment
Column       | Type     | Description
:------------| :--------|:---------------
id           | int      | auto-increment count
person_id    | int      | who performed the enrollement
course_id    | int      | which course the enrollment was made
status       | int      | approved, disapproved or failed
grade        | float    | final grade obtained
is_verified  | bool     | trivial
doc_verifier | []byte   | doc received to verifier the enrollment (photo)

## 4. career
Column        | Type      | Description
:-------------| :-------- |:---------------
id            | int       | auto-increment count
acronym       | string    | trivial
name          | string    | trivial
description   | string    | trivial
logo          | string    | pic of the career
is_verified   | bool      | trivial
doc_verifier  | []byte    | doc received to verifier it (photo)
reference     | string    | link of any reference
faculty_id    | int       | faculty to which it belongs
album_id      | int       | trivial

## 5. award
Column       | Type     | Description
:------------| :--------|:---------------
id           | int      | auto-increment count
year         | int      | year in which the award was obtained
period       | int      | semester in which the award was obtained
description  | string   | trivial
event        | string   | event in which the award was obtained
place        | int      | place obtained
is_verified  | bool     | trivial
doc_verifier | []byte   | doc received to verifier its existence (photo)
career_id    | int      | career to which it belongs
album_id     | int      | trivial

## 6. interest
Column       | Type     | Description
:------------| :--------|:---------------
id           | int      | auto-increment count
tag          | string   | brief title that represent the interest
is_skill     | bool     | is it a skill or just an interest?
person_id    | int      | user who belong the interest

## 7. photo
Column       | Type     | Description
:------------| :--------|:---------------
id           | int      | auto-increment count
pic          | []byte   | trivial
descripion   | string   | brief description
album_id     | int      | album to which the photo belongs

## 8. schedule
Column        | Type      | Description
:-------------| :---------|:---------------
id            | int       | auto-increment count
course_id     | string    | course to which the schedule belongs
start         | string    | start time
end           | int       | end time
day           | int       | day (monday, tuesday, etc.)
is_verified   | int       | trivial
doc_verifier  | bool      | doc received to verifie it (photo)

## 9. story
Column            | Type       | Description
:-----------------| :----------|:---------------
id                | int        | auto-increment count
title             | string     | trivial
body              | string     | trivial
recommended       | int        | number of users who recommend the story
unrecommended     | int        | number of user who unrecommend the story
views             | int        | number of users who saw the story
is_verified       | bool       | trivial
doc_verifier      | string     | doc received to verifier the story (photo)
allow_comments    | bool       | trivial
layer             | int        | in which layer the story was created
date              | date       | when the story was created
person_id         | int        | who created the story

## 10. tag_container
Column        | Type      | Description
:-------------| :---------|:---------------
id            | int       | auto-increment count
tag           | string    | brief title
agree         | string    | number of users who agree the tag
disagree      | int       | number of users who disagree the tag
reference     | int       | link or any reference
type          | int       | two possible types: speciality or requirement
layer         | int       | in which layer the story was created
career_id     | int       | career to which it belong
person_id     | int       | who created the tag

## 11. course_connection
Column            | Type    | Description
:-----------------| :-------|:---------------
id                | int     | auto-increment count
year              | int     | year to which the connection belongs
period            | int     | semester
child             | int     | id of the course child
parent            | int     | id of the course parent
threshold_credits | int     | minimum number of credits to open the parent course
career_id         | int     | career which has that course connection

## 12. course
Column          | Type    | Description
:---------------| :-------|:---------------
id              | int     | auto-increment count
code            | string  | code of the course
name            | string  | trivial
description     | string  | trivial
credits         | int     | number of credits of the course
group           | int     | a course can have many groups
year            | int     | year in which the course opened
period          | int     | semeste in which the course opened
semester        | int     | semester to which the course belongs
total_hours     | int     | trivial
is_verified     | bool    | trivial
doc_verifier    | []byte  | doc received to verifier it (photo)
type_course_id  | int     | can be many types: mandatory, optional, etc.
career_id       | int     | career to which the course belongs
person_id       | int     | professor in charge of the course

## 13. type_course
Column          | Type    | Description
:---------------| :-------|:---------------
id              | int     | auto-increment count
acronym         | string  | course type code
name            | string  | trivial
description     | string  | trivial
is_verified     | bool    | trivial
doc_verifier    | []byte  | doc received to verifier it (photo)
university_id   | int     | university to which the type courses belongs

## 14. faculty
Column          | Type    | Description
:---------------| :-------|:---------------
id              | int     | auto-increment count
acronym         | string  | faculty code
name            | string  | trivial
description     | string  | trivial
logo            | []byte  | trivial
is_verified     | bool    | trivial
doc_verifier    | []byte  | doc received to verifier it (photo)
university_id   | int     | university to which the faculty belongs

## 15. career_person
Column          | Type    | Description
:---------------| :-------|:---------------
person_id       | int     | trivial
career_id       | int     | trivial
is_verified     | bool    | trivial
doc_verifier    | []byte  | doc received to verifier it (photo)

## 16. person
Column         | Type     | Description
:--------------| :--------|:---------------
id             | int      | auto-increment count
code           | string   | university id
gender         | int      | trivial
first_name     | string   | trivial
last_name      | string   | trivial
phone          | string   | trivial
email          | string   | trivial
avatar         | []byte   | trivial
type           | int      | trivial
home_city      | string   | trivial
current_city   | string   | trivial
ethnic         | int      | trivial
nacionality    | string   | trivial
birth_date     | date     | trivial
admission_year | int      | year in which the user started university
period         | int      | semester
is_verified    | bool     | trivial
doc_verifier   | []byte   | doc received to verifier it (photo)

## 17. type_university
Column         | Type     | Description
:--------------| :--------|:---------------
id             | int      | auto-increment count
name           | string   | trivial
description    | string   | trivial
is_verified    | bool     | trivial
doc_verifier   | []byte   | doc received to verifier it (photo)

## 18. university
Column             | Type     | Description
:------------------| :--------|:---------------
id                 | int      | auto-increment count
local_rank         | int      | trivial
global_rank        | int      | trivial
acronym            | string   | trivial
name               | string   | trivial
region             | string   | trivial
description        | string   | trivial
latitude           | float    | trivial
longitude          | float    | trivial
logo               | []byte   | pic of the university
is_verified        | bool     | trivial
doc_verifier       | []byte   | doc received to verifier it (photo)
reference          | string   | link for the main web page of the university
type_university_id | int      | can be publi, private, etc.
album_id           | int      | trivial

## 19. album
Column       | Type     | Description
:------------| :--------|:---------------
id           | int      | auto-increment count
name         | string   | trivial
description  | string   | trivial