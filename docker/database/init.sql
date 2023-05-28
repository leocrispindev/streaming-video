CREATE DATABASE IF NOT EXISTS stream_content; 


CREATE TABLE IF NOT EXISTS video_info (
    id INT AUTO_INCREMENT PRIMARY KEY,
    titulo VARCHAR(100) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    category INT NOT NULL,
    duration DOUBLE,
    thumbName VARCHAR(255),
    height INT,
    width INT,
    indeless NUMBER(1)
);

