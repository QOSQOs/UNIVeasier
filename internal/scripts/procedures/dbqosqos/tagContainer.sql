DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertTagContainer(
  _id INT,
  _tag VARCHAR(100),
  _agree INT,
  _disagree INT,
  _reference VARCHAR(200),
  _type INT,
  _layer INT,
  _created_date DATE,
  _last_modified_date DATE,
  _is_verified INT,
  _doc_verifier BLOB,
  _career_id INT,
  _created_by INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.tag_container(tag, agree, disagree, reference, type, layer, created_date, last_modified_date, is_verified, doc_verifier, career_id, created_by) VALUES 
    (_tag, _agree, _disagree, _reference, _type, _layer, _created_date, _last_modified_date, _is_verified, _doc_verifier, _career_id, _created_by);
  ELSE
    UPDATE dbqosqos.tag_container
    SET
      tag = _tag,
      agree = _agree,
      disagree = _disagree,
      reference = _reference,
      type = _type,
      layer = _layer,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      is_verified = _is_verified,
      doc_verifier = _doc_verifier,
      career_id = _career_id,
      created_by = _created_by
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetTagContainerById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.tag_container
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListTagContainer()
BEGIN
    SELECT * FROM dbqosqos.tag_container;
END //

DELIMITER ;