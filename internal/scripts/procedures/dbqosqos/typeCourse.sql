DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertTypeCourse(
  _id INT,
  _acronym VARCHAR(50),
  _name VARCHAR(200),
  _description VARCHAR(1000),
  _is_verified INT,
  _doc_verifier BLOB,
  _created_date DATE,
  _last_modified_date DATE,
  _required_credits INT,
  _career_id INT,
  _created_by INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.type_course(acronym, name, description, is_verified, doc_verifier, created_date, last_modified_date, required_credits, career_id, created_by) VALUES 
    (_acronym, _name, _description, _is_verified, _doc_verifier, _created_date, _last_modified_date, _required_credits, _career_id, _created_by);
  ELSE
    UPDATE dbqosqos.type_course
    SET
      acronym = _acronym,
      name = _name,
      description = _description,
      is_verified = _is_verified,
      doc_verifier = _doc_verifier,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      required_credits = _required_credits,
      career_id = _career_id,
      created_by = _created_by
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetTypeCourseById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.type_course
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListTypeCourse()
BEGIN
    SELECT * FROM dbqosqos.type_course;
END //

DELIMITER ;