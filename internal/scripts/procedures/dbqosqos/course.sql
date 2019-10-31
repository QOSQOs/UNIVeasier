DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertCourse(
  _id INT,
  _code VARCHAR(50),
  _name VARCHAR(200),
  _description VARCHAR(1000),
  _credits INT,
  _batch INT,
  _year INT,
  _period INT,
  _semester INT,
  _total_hours INT,
  _is_verified INT,
  _doc_verifier BLOB,
  _created_date DATE,
  _last_modified_date DATE,
  _type_course_id INT,
  _career_id INT,
  _professor_id INT,
  _created_by INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.course(id, code, name, description, credits, batch, year, period, semester, total_hours, is_verified, doc_verifier, created_date, last_modified_date, type_course_id, career_id, professor_id, created_by) VALUES 
    (_id, _code, _name, _description, _credits, _batch, _year, _period, _semester, _total_hours, _is_verified, _doc_verifier, _created_date, _last_modified_date, _type_course_id, _career_id, _professor_id, _created_by);
  ELSE
    UPDATE dbqosqos.course
    SET
      code = _code,
      name = _name,
      description = _description,
      credits = _credits,
      batch = _batch,
      year = _year,
      period = _period,
      semester = _semester,
      total_hours = _total_hours,
      is_verified = _is_verified,
      doc_verifier = _doc_verifier,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      type_course_id = _type_course_id,
      career_id = _career_id,
      professor_id = _professor_id,
      created_by = _created_by
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetCourseById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.course
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListCourse()
BEGIN
    SELECT * FROM dbqosqos.course;
END //

DELIMITER ;