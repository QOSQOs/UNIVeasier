SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema dbqosqos
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `dbqosqos` DEFAULT CHARACTER SET utf8 ;
USE `dbqosqos` ;

-- -----------------------------------------------------
-- Table `dbqosqos`.`person`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`person` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`person` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(50) NULL,
  `gender` INT NULL,
  `first_name` VARCHAR(100) NULL,
  `last_name` VARCHAR(100) NULL,
  `phone` VARCHAR(20) NULL,
  `email` VARCHAR(100) NULL,
  `avatar` BLOB NULL,
  `type` INT NULL,
  `home_city` VARCHAR(100) NULL,
  `current_city` VARCHAR(100) NULL,
  `ethnic` INT NULL,
  `nationality` VARCHAR(100) NULL,
  `birth_date` DATE NULL,
  `admission_year` INT NULL,
  `period` INT NULL,
  `is_verified` INT NULL,
  `doc_verifier` BLOB NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`type_university`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`type_university` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`type_university` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(200) NULL,
  `description` VARCHAR(1000) NULL,
  `is_verified` INT NULL,
  `doc_verifier` BLOB NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `created_by` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_type_university_person1_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_type_university_person1`
    FOREIGN KEY (`created_by`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`album`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`album` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`album` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(200) NULL,
  `description` VARCHAR(1000) NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `created_by` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_album_person1_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_album_person1`
    FOREIGN KEY (`created_by`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`university`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`university` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`university` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `local_rank` INT NULL,
  `global_rank` INT NULL,
  `acronym` VARCHAR(50) NULL,
  `name` VARCHAR(200) NULL,
  `region` INT NULL,
  `description` VARCHAR(1000) NULL,
  `latitude` FLOAT NULL,
  `longitude` FLOAT NULL,
  `logo` BLOB NULL,
  `is_verified` INT NULL,
  `doc_verifier` BLOB NULL,
  `reference` VARCHAR(200) NULL,
  `type_university_id` INT NULL,
  `album_id` INT NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `created_by` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_university_type_university_idx` (`type_university_id` ASC) VISIBLE,
  INDEX `fk_university_album1_idx` (`album_id` ASC) VISIBLE,
  INDEX `fk_university_person1_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_university_type_university`
    FOREIGN KEY (`type_university_id`)
    REFERENCES `dbqosqos`.`type_university` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_university_album1`
    FOREIGN KEY (`album_id`)
    REFERENCES `dbqosqos`.`album` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_university_person1`
    FOREIGN KEY (`created_by`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`faculty`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`faculty` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`faculty` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `acronym` VARCHAR(50) NULL,
  `name` VARCHAR(200) NULL,
  `description` VARCHAR(1000) NULL,
  `logo` BLOB NULL,
  `is_verified` INT NULL,
  `doc_verifier` BLOB NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `university_id` INT NULL,
  `created_by` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_faculty_university1_idx` (`university_id` ASC) VISIBLE,
  INDEX `fk_faculty_person1_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_faculty_university1`
    FOREIGN KEY (`university_id`)
    REFERENCES `dbqosqos`.`university` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_faculty_person1`
    FOREIGN KEY (`created_by`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`career`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`career` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`career` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `acronym` VARCHAR(50) NULL,
  `name` VARCHAR(200) NULL,
  `description` VARCHAR(1000) NULL,
  `logo` BLOB NULL,
  `is_verified` INT NULL,
  `doc_verifier` BLOB NULL,
  `reference` VARCHAR(200) NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `faculty_id` INT NULL,
  `album_id` INT NULL,
  `created_by` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_career_faculty1_idx` (`faculty_id` ASC) VISIBLE,
  INDEX `fk_career_album1_idx` (`album_id` ASC) VISIBLE,
  INDEX `fk_career_person1_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_career_faculty1`
    FOREIGN KEY (`faculty_id`)
    REFERENCES `dbqosqos`.`faculty` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_career_album1`
    FOREIGN KEY (`album_id`)
    REFERENCES `dbqosqos`.`album` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_career_person1`
    FOREIGN KEY (`created_by`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`type_course`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`type_course` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`type_course` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `acronym` VARCHAR(50) NULL,
  `name` VARCHAR(200) NULL,
  `description` VARCHAR(1000) NULL,
  `is_verified` INT NULL,
  `doc_verifier` BLOB NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `required_credits` INT NULL,
  `career_id` INT NULL,
  `created_by` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_type_course_career1_idx` (`career_id` ASC) VISIBLE,
  INDEX `fk_type_course_person1_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_type_course_career1`
    FOREIGN KEY (`career_id`)
    REFERENCES `dbqosqos`.`career` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_type_course_person1`
    FOREIGN KEY (`created_by`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`course`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`course` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`course` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(50) NULL,
  `name` VARCHAR(200) NULL,
  `description` VARCHAR(1000) NULL,
  `credits` INT NULL,
  `batch` INT NULL,
  `year` INT NULL,
  `period` INT NULL,
  `semester` INT NULL,
  `total_hours` INT NULL,
  `is_verified` INT NULL,
  `doc_verifier` BLOB NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `type_course_id` INT NULL,
  `career_id` INT NULL,
  `professor_id` INT NULL,
  `created_by` INT NULL,
  INDEX `fk_course_type_course1_idx` (`type_course_id` ASC) VISIBLE,
  PRIMARY KEY (`id`),
  INDEX `fk_course_person1_idx` (`professor_id` ASC) VISIBLE,
  INDEX `fk_course_person2_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_requirement_career_career1`
    FOREIGN KEY (`career_id`)
    REFERENCES `dbqosqos`.`career` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_course_type_course1`
    FOREIGN KEY (`type_course_id`)
    REFERENCES `dbqosqos`.`type_course` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_course_person1`
    FOREIGN KEY (`professor_id`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_course_person2`
    FOREIGN KEY (`created_by`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`schedule`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`schedule` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`schedule` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `start` DATE NULL,
  `end` DATE NULL,
  `day` INT NULL,
  `room` VARCHAR(50) NULL,
  `type_lesson` INT NULL,
  `is_verified` INT NULL,
  `doc_verifier` BLOB NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `course_id` INT NULL,
  `created_by` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_schedule_course1_idx` (`course_id` ASC) VISIBLE,
  INDEX `fk_schedule_person1_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_schedule_course1`
    FOREIGN KEY (`course_id`)
    REFERENCES `dbqosqos`.`course` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_schedule_person1`
    FOREIGN KEY (`created_by`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`courses_connection`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`courses_connection` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`courses_connection` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `year` INT NULL,
  `period` INT NULL,
  `threshold_credits` INT NULL,
  `child_course` VARCHAR(45) NULL,
  `parent_course` VARCHAR(45) NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `career_id` INT NULL,
  `created_by` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_couse_conection_career1_idx` (`career_id` ASC) VISIBLE,
  INDEX `fk_courses_connection_person1_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_couse_conection_career1`
    FOREIGN KEY (`career_id`)
    REFERENCES `dbqosqos`.`career` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_courses_connection_person1`
    FOREIGN KEY (`created_by`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`career_person`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`career_person` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`career_person` (
  `career_id` INT NOT NULL,
  `person_id` INT NOT NULL,
  `is_verified` INT NULL,
  `doc_verifier` BLOB NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  PRIMARY KEY (`career_id`, `person_id`),
  INDEX `fk_career_has_person_person1_idx` (`person_id` ASC) VISIBLE,
  INDEX `fk_career_has_person_career1_idx` (`career_id` ASC) VISIBLE,
  CONSTRAINT `fk_career_has_person_career1`
    FOREIGN KEY (`career_id`)
    REFERENCES `dbqosqos`.`career` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_career_has_person_person1`
    FOREIGN KEY (`person_id`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`enrollment`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`enrollment` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`enrollment` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `status` INT NULL,
  `grade` FLOAT NULL,
  `is_verified` INT NULL,
  `doc_verifier` BLOB NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `person_id` INT NULL,
  `course_id` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_enrollment_person1_idx` (`person_id` ASC) VISIBLE,
  INDEX `fk_enrollment_course1_idx` (`course_id` ASC) VISIBLE,
  CONSTRAINT `fk_enrollment_person1`
    FOREIGN KEY (`person_id`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_enrollment_course1`
    FOREIGN KEY (`course_id`)
    REFERENCES `dbqosqos`.`course` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`contribution`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`contribution` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`contribution` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `score` INT NULL,
  `description` VARCHAR(1000) NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `person_id` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_contribution_person1_idx` (`person_id` ASC) VISIBLE,
  CONSTRAINT `fk_contribution_person1`
    FOREIGN KEY (`person_id`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`interest`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`interest` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`interest` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `tag` VARCHAR(100) NULL,
  `is_skill` TINYINT(1) NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `created_by` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_Interest_person1_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_Interest_person1`
    FOREIGN KEY (`created_by`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`story`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`story` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`story` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(100) NULL,
  `body` VARCHAR(500) NULL,
  `recommended` INT NULL,
  `unrecommended` INT NULL,
  `views` INT NULL,
  `is_verified` INT NULL,
  `doc_verifier` BLOB NULL,
  `allow_comments` TINYINT(1) NULL,
  `layer` INT NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `created_by` INT NULL,
  `career_id` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_story_person1_idx` (`created_by` ASC) VISIBLE,
  INDEX `fk_story_career1_idx` (`career_id` ASC) VISIBLE,
  CONSTRAINT `fk_story_person1`
    FOREIGN KEY (`created_by`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_story_career1`
    FOREIGN KEY (`career_id`)
    REFERENCES `dbqosqos`.`career` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`comment`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`comment` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`comment` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `content` VARCHAR(1000) NULL,
  `likes` INT NULL,
  `dislikes` INT NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `story_id` INT NULL,
  `created_by` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_comment_story1_idx` (`story_id` ASC) VISIBLE,
  INDEX `fk_comment_person1_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_comment_story1`
    FOREIGN KEY (`story_id`)
    REFERENCES `dbqosqos`.`story` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_comment_person1`
    FOREIGN KEY (`created_by`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`award`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`award` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`award` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `year` INT NULL,
  `period` INT NULL,
  `description` VARCHAR(500) NULL,
  `event` VARCHAR(100) NULL,
  `place` INT NULL,
  `is_verified` INT NULL,
  `doc_verifier` BLOB NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `career_id` INT NULL,
  `album_id` INT NULL,
  `created_by` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_Award_career1_idx` (`career_id` ASC) VISIBLE,
  INDEX `fk_Award_album1_idx` (`album_id` ASC) VISIBLE,
  INDEX `fk_award_person1_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_Award_career1`
    FOREIGN KEY (`career_id`)
    REFERENCES `dbqosqos`.`career` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Award_album1`
    FOREIGN KEY (`album_id`)
    REFERENCES `dbqosqos`.`album` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_award_person1`
    FOREIGN KEY (`created_by`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`photo`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`photo` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`photo` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `pic` BLOB NULL,
  `description` VARCHAR(200) NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `album_id` INT NULL,
  `created_by` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_photo_album1_idx` (`album_id` ASC) VISIBLE,
  INDEX `fk_photo_person1_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_photo_album1`
    FOREIGN KEY (`album_id`)
    REFERENCES `dbqosqos`.`album` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_photo_person1`
    FOREIGN KEY (`created_by`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `dbqosqos`.`tag_container`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `dbqosqos`.`tag_container` ;

CREATE TABLE IF NOT EXISTS `dbqosqos`.`tag_container` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `tag` VARCHAR(100) NULL,
  `agree` INT NULL,
  `disagree` INT NULL,
  `reference` VARCHAR(200) NULL,
  `type` INT NULL,
  `layer` INT NULL,
  `created_date` DATE NULL,
  `last_modified_date` DATE NULL,
  `is_verified` INT NULL,
  `doc_verifier` BLOB NULL,
  `career_id` INT NULL,
  `created_by` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_requirement_career1_idx` (`career_id` ASC) VISIBLE,
  INDEX `fk_tag_container_person1_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_requirement_career1`
    FOREIGN KEY (`career_id`)
    REFERENCES `dbqosqos`.`career` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_tag_container_person1`
    FOREIGN KEY (`created_by`)
    REFERENCES `dbqosqos`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
