DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertInterest(
  _id INT,
  _tag VARCHAR(100),
  _is_skill TINYINT(1),
  _created_date DATE,
  _last_modified_date DATE,
  _created_by INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.interest(id, tag, is_skill, created_date, last_modified_date, created_by) VALUES 
    (_id, _tag, _is_skill, _created_date, _last_modified_date, _created_by);
  ELSE
    UPDATE dbqosqos.interest
    SET
      tag = _tag,
      is_skill = _is_skill,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      created_by = _created_by
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetInterestById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.interest
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListInterest()
BEGIN
    SELECT * FROM dbqosqos.interest;
END //

DELIMITER ;