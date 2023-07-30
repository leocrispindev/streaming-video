package com.br.watcher.services

import com.br.watcher.broker.BrokerAdmin
import com.br.watcher.broker.Consumer
import com.br.watcher.broker.Producer
import com.br.watcher.model.Message
import com.br.watcher.model.VideoBrokerMessage
import com.google.gson.Gson
import io.ktor.websocket.*
import org.apache.kafka.clients.consumer.ConsumerRecords
import org.slf4j.MDC
import java.util.UUID

object Watch {

    var producer = Producer
    var admin = BrokerAdmin
    suspend fun Start(conn : DefaultWebSocketSession) {

        for (frame in conn.incoming) {
            frame as? Frame.Text ?: continue

            //val message = Gson().toJson(frame.data)

            val videoInfo = Gson().fromJson(String(frame.data), Message::class.java)

            val normalizeVideoName = videoInfo.videoName.replace(" ", "-")

            val topicName = "video-topic-${normalizeVideoName}"

            if (!admin.topicExist(topicName)){
                admin.createTopic(topicName)
                messageToFileSentry(normalizeVideoName, topicName, videoInfo.videoExtension)
            }

            val consumer = Consumer(topicName, videoInfo.session)
            println("Session: " + videoInfo.session)

            println("Consumer criado")
            while (true) run {
                val records: ConsumerRecords<String, String> = consumer.StartPoll(1);

                for (record in records) {
                    conn.send(record.value())
                }

            }
        }

    }

    private fun messageToFileSentry(normalizeVideoName : String, topicName : String, extension : String) {
        try {
            producer.beginTransaction()
            producer.send("stream-content", VideoBrokerMessage(normalizeVideoName, UUID.randomUUID().toString(), topicName, extension))
        }catch (e : Exception) {
            producer.abortTransaction()
            // TODO send error notification to client
            return
        }
    }

}

