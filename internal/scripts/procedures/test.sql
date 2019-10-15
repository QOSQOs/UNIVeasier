DELIMITER //

CREATE PROCEDURE GetListTest()
BEGIN
    SELECT *  FROM dbtest.tabletest;
END //

CREATE PROCEDURE GetTest(_name VARCHAR(255))
BEGIN
    SELECT *  FROM dbtest.tabletest WHERE name = _name;
END //

DELIMITER ;