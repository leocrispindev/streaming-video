package com.br.watcher.services

import com.br.watcher.broker.Consumer
import com.br.watcher.broker.Producer
import com.br.watcher.model.Message
import com.br.watcher.model.VideoBrokerMessage
import com.google.gson.Gson
import io.ktor.websocket.*
import org.apache.kafka.clients.consumer.ConsumerRecords

object Watch {

    var producer = Producer
    suspend fun Start(conn : DefaultWebSocketSession) {

        for (frame in conn.incoming) {
            val message = Gson().toJson(frame)
            val videoInfo = Gson().fromJson(message, Message::class.java)

            val normalizeVideoName = videoInfo.videoName.replace(" ", "-")

            val topicName = "video-topic-${normalizeVideoName}"

            try {
                producer.beginTransaction()

                producer.send("stream-content", VideoBrokerMessage(normalizeVideoName, conn.toString(), topicName))
            }catch (e : Exception) {
                producer.abortTransaction()
                // TODO send error notification to client
                return
            }

            // TODO init consumer for topicName
            val consumer = Consumer(topicName)

            while (true) run {
                val records: ConsumerRecords<String, String> = consumer.StartPoll(1);

                for (record in records) {
                    conn.send(record.value())
                }

            }


        }

    }

}

