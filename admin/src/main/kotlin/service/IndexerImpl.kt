package service

import Properties
import model.Document
import model.VideoInfo
import org.apache.kafka.common.errors.ProducerFencedException
import service.queue.Broker

class IndexerImpl : Indexer {

    val producer = Broker
    val index = Properties.TOPIC_INDEX.getValue()
    override fun index(videoInfo: VideoInfo) {
        producer.initTransaction()

        val id = videoInfo.id ?: throw Exception("Video ID cannot be empty")

        val key = "video-index-id-$id"
        try {
            producer.beginTransaction()

            val document = Document(action = 1, key = key, id = id, videoInfo = videoInfo)

            producer.send(index, document)

            producer.commit()

        }catch (e : ProducerFencedException) {
            println("Error commit index document: [key]=$key")
            producer.abortTransaction()
        }
    }

    override fun delete(id: Int) {
        producer.initTransaction()

        val key = "video-index-id-" + id
        try {
            producer.beginTransaction()

            val document = Document(action = 2, key = key, id = id ,videoInfo = null)

            producer.send(index, document)

            producer.commit()

        }catch (e : ProducerFencedException) {
            println("Error commit index document: [key]=$key")
            producer.abortTransaction()
        }
    }
}