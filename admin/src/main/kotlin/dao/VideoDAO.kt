package dao

import model.VideoInfo
import service.database.Database
import java.sql.SQLException

class VideoDAO {

    private var connection = Database.getConn()

    fun insert(video: VideoInfo) : Int{

        try {
            val query = "INSERT INTO video_info (titulo, descricao, category, duration, thumbName, height, width) " +
                    "VALUES (?, ?, ?, ?, ?, ?, ?)"

            val statement = connection.prepareStatement(query)
            statement.setString(1, video.titulo)
            statement.setString(2, video.descricao)
            statement.setInt(3,video.category)
            statement.setDouble(4, video.duration)
            statement.setString(5, video.thumbName)
            statement.setInt(6, video.height)
            statement.setInt(7, video.width)

            val rowEffected = statement.executeUpdate()

            if (rowEffected != 1)
                throw SQLException("insert error, rows effected not equal 1")

            val key = statement.generatedKeys

            //Return ID
            return key.getInt(1)

        }catch (e : SQLException){
            throw e
        }

        return -1
    }

}