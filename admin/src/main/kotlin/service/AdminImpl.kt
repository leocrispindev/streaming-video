package service

import dao.VideoDAO
import model.VideoInfo

object AdminImpl : Admin {

    private var videoDAO = VideoDAO()
    override fun insert(videoInfo: VideoInfo): Int {
        return videoDAO.insert(videoInfo)
    }

    override fun update(videoInfo: VideoInfo): Int {
        TODO("Not yet implemented")
    }

    override fun upload() {
        TODO("Not yet implemented")
    }

    override fun index(videoInfo: VideoInfo): Int {
        TODO("Not yet implemented")
    }

}