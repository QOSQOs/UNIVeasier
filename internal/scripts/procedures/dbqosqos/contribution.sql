DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertContribution(
  _id INT,
  _score INT,
  _description VARCHAR(1000),
  _created_date DATE,
  _last_modified_date DATE,
  _person_id INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.contribution(id, score, description, created_date, last_modified_date, person_id) VALUES 
    (_id, _score, _description, _created_date, _last_modified_date, _person_id);
  ELSE
    UPDATE dbqosqos.contribution
    SET
      score = _score,
      description = _description,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      person_id = _person_id
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetContributionById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.contribution
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListContribution()
BEGIN
    SELECT * FROM dbqosqos.contribution;
END //

DELIMITER ;