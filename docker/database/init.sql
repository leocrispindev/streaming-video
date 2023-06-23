CREATE DATABASE IF NOT EXISTS stream_content; 


CREATE TABLE IF NOT EXISTS video_info (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    synopsis VARCHAR(255) NOT NULL,
    category INT NOT NULL,
    duration INT,
    extension VARCHAR(50),
    indexless TINYINT(1)
);

