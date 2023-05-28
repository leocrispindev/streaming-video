package service

import dao.VideoDAO
import model.VideoDTO
import model.VideoInfo
import java.sql.SQLException

object AdminImpl : Admin {

    private var videoDAO = VideoDAO()
    override fun insert(video: VideoDTO): Int {

        val vd = VideoInfo(id = null, titulo = video.titulo, descricao = video.descricao, category = video.category, duration = 0.0, indexless = video.indexless)

        return videoDAO.insert(vd)
    }

    override fun update(video: VideoDTO) {
        try {
            val vd = VideoInfo(id = video.id, titulo = video.titulo, descricao = video.descricao, category = video.category)

            videoDAO.update(vd)
        }catch (e : SQLException) {
            throw SQLException("Error on update video", e)
        }
    }

    override fun index(videoInfo: VideoInfo): Int {
        TODO("Not yet implemented")
    }

    override fun getAll(): ArrayList<VideoInfo> {
        return videoDAO.get()
    }

    override fun delete(id: Int) {
       videoDAO.delete(id)
    }


}