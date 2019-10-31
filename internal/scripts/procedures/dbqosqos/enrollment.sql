DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertEnrollment(
  _id INT,
  _status INT,
  _grade FLOAT,
  _is_verified INT,
  _doc_verifier BLOB,
  _created_date DATE,
  _last_modified_date DATE,
  _person_id INT,
  _course_id INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.enrollment(id, status, grade, is_verified, doc_verifier, created_date, last_modified_date, person_id, course_id) VALUES 
    (_id, _status, _grade, _is_verified, _doc_verifier, _created_date, _last_modified_date, _person_id, _course_id);
  ELSE
    UPDATE dbqosqos.enrollment
    SET
      status = _status,
      grade = _grade,
      is_verified = _is_verified,
      doc_verifier = _doc_verifier,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      person_id = _person_id,
      course_id = _course_id
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetEnrollmentById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.enrollment
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListEnrollment()
BEGIN
    SELECT * FROM dbqosqos.enrollment;
END //

DELIMITER ;