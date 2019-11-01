DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertCareerPerson(
  _id INT,
  _career_id INT,
  _person_id INT,
  _is_verified INT,
  _doc_verifier BLOB,
  _created_date DATE,
  _last_modified_date DATE)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.career_person(career_id, person_id, is_verified, doc_verifier, created_date, last_modified_date) VALUES 
    (_career_id, _person_id, _is_verified, _doc_verifier, _created_date, _last_modified_date);
  ELSE
    UPDATE dbqosqos.career_person
    SET
      career_id = _career_id,
      person_id = _person_id,
      is_verified = _is_verified,
      doc_verifier = _doc_verifier,
      created_date = _created_date,
      last_modified_date = _last_modified_date
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetCareerPersonById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.career_person
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListCareerPerson()
BEGIN
    SELECT * FROM dbqosqos.career_person;
END //

DELIMITER ;