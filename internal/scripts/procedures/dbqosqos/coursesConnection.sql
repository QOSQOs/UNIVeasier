DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertCoursesConnection(
  _id INT,
  _year INT,
  _period INT,
  _threshold_credits INT,
  _child_course VARCHAR(45),
  _parent_course VARCHAR(45),
  _created_date DATE,
  _last_modified_date DATE,
  _career_id INT,
  _created_by INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.courses_connection(year, period, threshold_credits, child_course, parent_course, created_date, last_modified_date, career_id, created_by) VALUES 
    (_year, _period, _threshold_credits, _child_course, _parent_course, _created_date, _last_modified_date, _career_id, _created_by);
  ELSE
    UPDATE dbqosqos.courses_connection
    SET
      year = _year,
      period = _period,
      threshold_credits = _threshold_credits,
      child_course = _child_course,
      parent_course = _parent_course,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      career_id = _career_id,
      created_by = _created_by
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetCoursesConnectionById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.courses_connection
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListCoursesConnection()
BEGIN
    SELECT * FROM dbqosqos.courses_connection;
END //

DELIMITER ;