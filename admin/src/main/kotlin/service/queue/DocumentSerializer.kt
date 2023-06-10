package service.queue

import com.google.gson.Gson
import model.Document
import org.apache.kafka.common.serialization.Serializer

object DocumentSerializer : Serializer<Document> {
    override fun serialize(topic: String?, data: Document?): ByteArray? {

        try {
            return Gson().toJson(data).toByteArray()
        }catch (e : Exception) {
            println("Error in serializing object"+ data?.videoInfo?.id);
        }

        return null
    }
}