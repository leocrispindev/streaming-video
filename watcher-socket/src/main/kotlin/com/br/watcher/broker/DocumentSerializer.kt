package com.br.watcher.broker

import com.br.watcher.model.VideoBrokerMessage
import com.google.gson.Gson
import org.apache.kafka.common.serialization.Serializer

class DocumentSerializer : Serializer<VideoBrokerMessage> {
    override fun serialize(topic: String?, data: VideoBrokerMessage?): ByteArray? {

        try {
            return Gson().toJson(data).toByteArray()
        }catch (e : Exception) {
            if (data != null) {
                println("Error in serializing object"+ data.topicName)
            };
        }

        return null
    }
}