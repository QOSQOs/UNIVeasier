DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertComment(
  _id INT,
  _content VARCHAR(1000),
  _likes INT,
  _dislikes INT,
  _created_date DATE,
  _last_modified_date DATE,
  _story_id INT,
  _created_by INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.comment(id, content, likes, dislikes, created_date, last_modified_date, story_id, created_by) VALUES 
    (_id, _content, _likes, _dislikes, _created_date, _last_modified_date, _story_id, _created_by);
  ELSE
    UPDATE dbqosqos.comment
    SET
      content = _content,
      likes = _likes,
      dislikes = _dislikes,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      story_id = _story_id,
      created_by = _created_by
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetCommentById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.comment
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListComment()
BEGIN
    SELECT * FROM dbqosqos.comment;
END //

DELIMITER ;