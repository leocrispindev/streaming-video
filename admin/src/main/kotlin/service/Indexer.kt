package service

import model.Document
import model.VideoInfo

interface Indexer {

    fun index(videoInfo: VideoInfo)

    fun delete(id : Int)

}