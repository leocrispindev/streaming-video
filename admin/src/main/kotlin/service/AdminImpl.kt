package service

import dao.VideoDAO
import model.VideoDTO
import model.VideoInfo
import java.sql.SQLException

object AdminImpl : Admin {

    private var videoDAO = VideoDAO()

    private var indexer = IndexerImpl()
    override fun insert(video: VideoDTO): Int {
        return try {
            val vd = VideoInfo(id = null, titulo = video.titulo, descricao = video.descricao, category = video.category, duration = 0.0, indexless = video.indexless)
            vd.id = videoDAO.insert(vd)

            indexer.index(vd)

            vd.id!!

        }catch (e : SQLException) {
            println("Error on video update [id]=${video.id}")
            -1
        }
    }

    override fun update(video: VideoDTO) {
        try {
            val vd = VideoInfo(id = video.id, titulo = video.titulo, descricao = video.descricao, category = video.category)

            videoDAO.update(vd)

            indexer.index(videoInfo = vd)
        }catch (e : SQLException) {
            println("Error on video update [id]=${video.id}")
        }
    }

    override fun getAll(): ArrayList<VideoInfo> {
        return videoDAO.get()
    }

    override fun delete(id: Int) {
        videoDAO.delete(id)

        indexer.delete(id)
    }


}