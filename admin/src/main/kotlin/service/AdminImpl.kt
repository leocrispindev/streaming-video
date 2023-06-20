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
            val vd = VideoInfo(id = null, title = video.title, synopsis = video.synopsis, category = video.category, duration = 0, indexless = video.indexless, extension = video.extension)
            vd.id = videoDAO.insert(vd)

            if(vd.indexless)
                indexer.index(vd)

            vd.id!!

        }catch (e : SQLException) {
            println("Error on video insert [title]=${video.title}, [message]=${e.message}")
            throw e
        }
    }

    override fun update(video: VideoDTO) {
        try {
            val vd = VideoInfo(video.id, video.title, video.synopsis, video.category, duration = 0, indexless = video.indexless,video.extension)

            videoDAO.update(vd)

            if (vd.indexless)
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