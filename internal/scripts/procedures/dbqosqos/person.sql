DELIMITER //

CREATE PROCEDURE dbqosqos.InsertPerson(
  _id INT,
  _code VARCHAR(50),
  _gender INT,
  _first_name VARCHAR(100),
  _last_name VARCHAR(100),
  _phone VARCHAR(20),
  _email VARCHAR(100),
  _avatar BLOB,
  _type INT,
  _home_city VARCHAR(100),
  _current_city VARCHAR(100),
  _ethnic INT,
  _nationality VARCHAR(100),
  _birth_date DATE,
  _admission_year INT,
  _period INT,
  _is_verified INT,
  _doc_verifier BLOB,
  _created_date DATE,
  _last_modified_date DATE)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.person(code, gender, first_name, last_name, phone, email, avatar, type, home_city, current_city, ethnic, nationality, birth_date, admission_year, period, is_verified, doc_verifier, created_date, last_modified_date) VALUES
    (_code, _gender, _first_name, _last_name, _phone, _email, _avatar, _type, _home_city, _current_city, _ethnic, _nationality, _birth_date, _admission_year, _period, _is_verified, _doc_verifier, _created_date, _last_modified_date);
  ELSE
    UPDATE dbqosqos.person
    SET
      code = _code,
      gender = _gender,
      first_name = _first_name,
      last_name = _last_name,
      phone = _phone,
      email = _email,
      avatar = _avatar,
      type = _type,
      home_city = _home_city,
      current_city = _current_city,
      ethnic = _ethnic,
      nationality = _nationality,
      birth_date = _birth_date,
      admission_year = _admission_year,
      period = _period,
      is_verified = _is_verified,
      doc_verifier = _doc_verifier,
      created_date = _created_date,
      last_modified_date = _last_modified_date
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetPersonById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.person 
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListPerson()
BEGIN
    SELECT * FROM dbqosqos.person;
END //

DELIMITER ;